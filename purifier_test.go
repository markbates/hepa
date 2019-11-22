package hepa

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Purifier_Filter(t *testing.T) {
	r := require.New(t)

	in := []byte("this is us")

	p := New()
	p = WithFunc(p, func(b []byte) ([]byte, error) {
		return bytes.ReplaceAll(b, []byte(" is "), []byte(" was ")), nil
	})

	out, err := p.Filter(in)
	r.NoError(err)

	r.Equal([]byte("this was us"), out)
}
