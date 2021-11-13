package filters

import (
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

		return replace(b, u, "PWD"), nil
	}
}
