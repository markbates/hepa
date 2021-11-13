package hepa

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/markbates/hepa/filters"
)

type Purifier struct {
	filter Filter
	parent *Purifier
	sync.Mutex
}

func (p *Purifier) Filter(b []byte) ([]byte, error) {
	if p == nil || p.filter == nil {
		return b, nil
	}

	p.Lock()
	b, err := p.filter.Filter(b)
	p.Unlock()

	if err != nil {
		return b, err
	}

	return b, nil
}

func (p *Purifier) Clean(r io.Reader) ([]byte, error) {
	bb := &bytes.Buffer{}

	if p == nil {
		p = &Purifier{}
	}

	reader := bufio.NewReader(r)

	for {
		line, _, err := reader.ReadLine()

		if err != nil && err == io.EOF {
			break
		}

		// filter the line
		line, err = p.Filter(line)
		if err != nil {
			return nil, err
		}

		fmt.Fprintln(bb, string(line))
	}

	if p.parent != nil {
		return p.parent.Clean(bb)
	}

	home := filters.Home()
	b, err := home.Filter(bb.Bytes())
	if err != nil {
		return nil, err
	}

	return b, nil
}

// New returns a new Purifier
// with all the filters added
func Deep() *Purifier {
	p := &Purifier{}
	p = With(p, filters.PWD())
	p = With(p, filters.Secrets())
	p = With(p, filters.Golang())
	return p
}
