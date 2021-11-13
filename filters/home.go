package filters

import (
	"bytes"
	"os"
)

func Home() FilterFn {
	return func(b []byte) ([]byte, error) {
		u, err := os.UserHomeDir()

		if err != nil {
			return nil, err
		}

		return bytes.ReplaceAll(b, []byte(u), []byte("$HOME")), nil
	}
}

func PWD() FilterFn {
	return func(b []byte) ([]byte, error) {
		u, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		return bytes.ReplaceAll(b, []byte(u), []byte("$PWD")), nil
	}
}
