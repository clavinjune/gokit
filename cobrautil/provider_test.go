package cobrautil_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/clavinjune/gokit/slogutil"

	"github.com/clavinjune/gokit/cobrautil"
	"github.com/clavinjune/gokit/testutil"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	r := require.New(t)
	buf := new(bytes.Buffer)
	opt := slogutil.DefaultOption
	opt.Writer = buf

	root := cobrautil.New("testing", "test", &cobrautil.Option{
		SlogOption:   opt,
		SetOutToSlog: true,
		SetErrToSlog: true,
	})
	root.RunE = func(cmd *cobra.Command, args []string) error {
		cmd.Println("test json")
		return nil
	}

	_, err := testutil.CobraExecute(t, root, "--json")
	r.NoError(err)

	m := map[string]any{}
	r.NoError(json.NewDecoder(buf).Decode(&m))

	r.Equal("INFO", m["level"])
	r.Equal("test json", m["msg"])
}
