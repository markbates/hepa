package filters

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Golang(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	bb := &bytes.Buffer{}
	c := exec.CommandContext(ctx, "go", "env")
	c.Stdout = bb

	err := c.Run()
	r.NoError(err)

	b, err := Golang().Filter(bb.Bytes())
	r.NoError(err)

	b, err = Home().Filter(b)
	r.NoError(err)

	u, err := os.UserHomeDir()
	r.NoError(err)

	// fmt.Println(string(b))
	r.NotContains(string(b), u)

}
