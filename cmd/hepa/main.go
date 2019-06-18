package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/markbates/hepa"
	"github.com/markbates/hepa/filters"
)

var env bool

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args[1:]
	var r io.Reader = os.Stdin
	if fi.Mode()&os.ModeNamedPipe == 0 {
		if len(args) == 0 {
			fmt.Println("no pipe :(")
			return
		}
		r = strings.NewReader(args[0])
	}

	hep := hepa.New()

	hep = hepa.With(hep, filters.Golang())
	hep = hepa.With(hep, filters.Secrets())
	b, err := hep.Clean(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(b))
}
