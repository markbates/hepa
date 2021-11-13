package filters

import (
	"bytes"
	"os/exec"
	"strings"
	"sync"
)

var defaultGolang = &golang{}

func Golang() FilterFn {
	return func(b []byte) ([]byte, error) {
		b, err := defaultGolang.Filter(b)
		if err != nil {
			return nil, err
		}

		return replace(b, "$GOMOD", "go.mod"), nil
	}
}

type golang struct {
	sync.Mutex
	env map[string]string
	sync.Once
}

func (g *golang) Filter(b []byte) ([]byte, error) {

	env, err := g.environ()
	if err != nil {
		return nil, err
	}

	for _, k := range g.keys() {
		v := env[k]
		if len(v) == 0 {
			continue
		}

		k = "$" + k
		b = replace(b, v, k)
	}

	b = replace(b, "$GOMOD", "\"go.mod\"")
	b = replace(b, "/usr/local", "$USR")
	return b, nil
}

func (g *golang) keys() []string {
	return []string{
		"GOMOD",
		"GOGCCFLAGS",
		"GOROOT",
		"GOTOOLDIR",
		"GOPATH",
		"GOPRIVATE",
	}
}

func (g *golang) environ() (map[string]string, error) {
	g.Lock()
	defer g.Unlock()

	var err error
	g.Do(func() {
		m := map[string]string{}

		bb := &bytes.Buffer{}
		c := exec.Command("go", "env")
		c.Stdout = bb

		err = c.Run()
		if err != nil {
			return
		}

		body := bb.String()

		for _, line := range strings.Split(body, "\n") {
			spl := strings.Split(line, "=")
			if len(spl) < 2 {
				continue
			}

			key := spl[0]

			for _, k := range g.keys() {
				if k == key {
					val := strings.Join(spl[1:], "=")
					m[key] = val
					break
				}
			}

		}
		g.env = m
	})

	return g.env, err
}
