package filters

import (
	"os"
	"path/filepath"
)

func Home() FilterFn {
	return func(b []byte) ([]byte, error) {
		u, err := os.UserHomeDir()

		if err != nil {
			return nil, err
		}

		return replace(b, u, "$HOME"), nil
	}
}

func PWD() FilterFn {
	return func(b []byte) ([]byte, error) {
		u, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		u, err = filepath.Abs(u)
		if err != nil {
			return nil, err
		}

		if len(u) == 0 || u == string(filepath.Separator) {
			return b, nil
		}

		return replace(b, u, "$PWD"), nil
	}
}

func Replace(s string, r string) FilterFn {
	return func(b []byte) ([]byte, error) {
		return replace(b, s, r), nil
	}
}
