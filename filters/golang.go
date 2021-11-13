package filters

import (
	"bytes"
	"fmt"
	"os"
)

func Golang() FilterFn {
	return func(b []byte) ([]byte, error) {

		for _, env := range goEnvs() {
			fmt.Printf("TODO >> golang.go:13 env %[1]T %[1]v\n", env)
			r := fmt.Sprintf("$%s", env)

			fmt.Printf("TODO >> golang.go:16 r %[1]T %[1]v\n", r)
			b = bytes.ReplaceAll(b, []byte(os.Getenv(env)), []byte(r))

		}

		fmt.Printf("TODO >> golang.go:21 string(b) %[1]T %[1]v\n", string(b))
		return b, nil
	}
}

func goEnvs() []string {
	return []string{
		"GOCACHE",
		"GOENV",
		"GOMODCACHE",
		"GONOPROXY",
		"GONOSUMDB",
		"GOPATH",
		"GOPRIVATE",
		"GOPROXY",
		"GOROOT",
		"GOSUMDB",
		"GOTOOLDIR",
		"GOMOD",
	}
}
