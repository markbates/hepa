package filters

import (
	"os"
)

func Golang() FilterFn {
	return func(b []byte) ([]byte, error) {

		for _, env := range goEnvs() {
			b = replace(b, os.Getenv(env), env)
		}

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
