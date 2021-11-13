package filters

import (
	"fmt"
	"os"
)

func Home() FilterFn {
	return func(b []byte) ([]byte, error) {
		u, err := os.UserHomeDir()

		if err != nil {
			return nil, err
		}

		return replace(b, u, "HOME"), nil
	}
}

func PWD() FilterFn {
	return func(b []byte) ([]byte, error) {
		u, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		fmt.Printf("TODO >> home.go:26 u %[1]T %[1]v\n", u)
		if len(u) == 0 {
			return b, nil
		}

		return replace(b, u, "PWD"), nil
	}
}
