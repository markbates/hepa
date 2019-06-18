package hepa

import (
	"bytes"
	"log"
	"os"
)

func WithFunc(p Purifier, fn FilterFunc) Purifier {
	c := New()
	c.parent = &p
	c.filter = fn
	return c
}

func With(p Purifier, f Filter) Purifier {
	c := New()
	c.parent = &p
	c.filter = f
	return c
}

var home = func() []byte {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		home = pwd
	}

	return []byte(home)
}()

func homeFilter(b []byte) ([]byte, error) {
	return bytes.ReplaceAll(b, home, []byte("$HOME")), nil
}
