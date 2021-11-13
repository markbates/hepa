package filters

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PWD(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pwd, err := os.Getwd()
	r.NoError(err)
	r.NotEmpty(pwd)
	r.NotEqual("/", pwd)

	in := fmt.Sprintf("Getwd: %s", pwd)
	r.NotContains(in, "$PWD")

	out, err := PWD().Filter([]byte(in))
	r.NoError(err)

	act := string(out)
	r.Contains(act, "$PWD")

}
