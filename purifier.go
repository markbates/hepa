package hepa

import (
	"bufio"
	"bytes"
	"io"
)

type Purifier struct {
	parent *Purifier
	filter Filter
}

func (p Purifier) Filter(b []byte) ([]byte, error) {
	if p.filter == nil {
		if p.parent != nil {
			return p.parent.Filter(b)
		}
		return homeFilter(b)
	}
	return p.filter.Filter(b)
}

func (p Purifier) Clean(r io.Reader) ([]byte, error) {
	bb := &bytes.Buffer{}

	if p.filter == nil {
		if p.parent != nil {
			return p.parent.Clean(r)
		}
		_, err := io.Copy(bb, r)
		return bb.Bytes(), err
	}

	reader := bufio.NewReader(r)
	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		input, err = p.Filter(input)
		if err != nil {
			return nil, err
		}
		input, err = homeFilter(input)
		if err != nil {
			return nil, err
		}
		bb.Write(input)
		bb.Write([]byte("\n"))
	}

	return bb.Bytes(), nil
}

func New() Purifier {
	return Purifier{}
}
