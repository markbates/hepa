package hepa

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Purifier_Filter(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	u, err := os.UserHomeDir()
	r.NoError(err)

	in := fmt.Sprintf("HOME: %s\n", u)

	var me FilterFn = func(b []byte) ([]byte, error) {
		return bytes.ReplaceAll(b, []byte("ME"), []byte("me")), nil
	}
	var ho FilterFn = func(b []byte) ([]byte, error) {
		return bytes.ReplaceAll(b, []byte("HO"), []byte("ho")), nil
	}

	table := []struct {
		name string
		p    *Purifier
		exp  string
		err  bool
	}{
		{name: "clean $HOME by default", exp: `HOME: $HOME`},
		{name: "clean me", p: WithFunc(&Purifier{}, me), exp: `HOme: $HOme`},
		{
			name: "clean ho/me",
			p:    WithFunc(WithFunc(&Purifier{}, me), ho),
			exp:  `home: $home`,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			r := require.New(t)

			p := tt.p

			out, err := p.Clean(strings.NewReader(in))
			if tt.err {
				r.Error(err)
				return
			}

			r.NoError(err)

			act := string(out)

			// always make sure home dir is not in the output
			r.NotContains(act, u)

			r.Contains(act, tt.exp)

		})
	}
}
