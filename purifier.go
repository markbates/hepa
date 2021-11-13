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

func (p *Purifier) clone() *Purifier {
	p.Lock()
	defer p.Unlock()

	return &Purifier{
		filter: p.filter,
		parent: p,
	}
}

func (p *Purifier) Filter(b []byte) ([]byte, error) {
	if p == nil {
		p = &Purifier{}
	}

	p.Lock()

	if p.filter == nil {
		p.filter = Noop()
	}

	f := p.filter

	p.Unlock()

	b, err := f.Filter(b)
	if err != nil {
		return b, err
	}

	home := filters.Home()

	b, err = home.Filter(b)
	if err != nil {
		return nil, err
	}

	if p.parent != nil {
		return p.parent.Filter(b)
	}

	return b, nil
}

func (p *Purifier) Clean(r io.Reader) ([]byte, error) {
	bb := &bytes.Buffer{}

	if p == nil {
		p = &Purifier{}
	}

	p.Lock()

	if p.filter == nil && p.parent != nil {
		p.Unlock()
		return p.parent.Clean(r)
	}

	p.Unlock()

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

	return bb.Bytes(), nil
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
