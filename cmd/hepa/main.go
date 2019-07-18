package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/markbates/hepa"
	"github.com/markbates/hepa/filters"
)

var env bool

var opts = struct {
	*flag.FlagSet
	file  string
	dir   string
	write bool
}{
	FlagSet: flag.NewFlagSet("hepa", flag.ExitOnError),
}

func main() {
	opts.StringVar(&opts.file, "f", "", "read and filter file (read-only)")
	opts.StringVar(&opts.dir, "d", "", "read and filter dir (read-only/recursive)")
	opts.BoolVar(&opts.write, "w", false, "write to disk")

	if err := opts.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	hep := hepa.New()
	hep = hepa.With(hep, filters.Golang())
	hep = hepa.With(hep, filters.Secrets())

	if len(opts.file) > 0 {
		cleanFile(opts.file, hep)
		return
	}

	if len(opts.dir) > 0 {
		cleanDir(opts.dir, hep)
		return
	}

	cleanPipe(hep)
}

func cleanFile(p string, hep hepa.Purifier) {
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}

	b, err := hep.Clean(f)
	if err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	if !opts.write {
		fmt.Print(string(b))
		return
	}

	f, err = os.Create(p)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.Write(b); err != nil {
		log.Fatal(err)
	}
}

func cleanDir(root string, hep hepa.Purifier) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		cleanFile(path, hep)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func cleanPipe(hep hepa.Purifier) {
	args := opts.Args()
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var r io.Reader = os.Stdin
	if fi.Mode()&os.ModeNamedPipe == 0 {
		if len(args) == 0 {
			fmt.Println("no pipe :(")
			return
		}
		r = strings.NewReader(args[0])
		args = args[1:]
	}

	for _, a := range args {
		hep = hepa.Clean(hep, []byte(a))
	}

	b, err := hep.Clean(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(b))
}
