package filters

import (
	"os"
)

type FilterFunc func([]byte) ([]byte, error)

func (f FilterFunc) Filter(b []byte) ([]byte, error) {
	return f(b)
}

type dir struct {
	Dir string
	Err error
}

var home = func() dir {
	var d dir
	home, ok := os.LookupEnv("HOME")
	if !ok {
		pwd, err := os.Getwd()
		if err != nil {
			d.Err = err
			return d
		}
		home = pwd
	}
	d.Dir = home

	return d
}()
