package fizz

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Exec(t *testing.T) {
	r := require.New(t)

	b := NewBubbler(nil)
	f := fizzer{b}
	bb := &bytes.Buffer{}
	c := f.Exec(bb)
	r.NoError(c("echo hello"))
	r.Equal("hello", strings.TrimSpace(bb.String()))
}

func Test_ExecQuoted(t *testing.T) {
	r := require.New(t)

	b := NewBubbler(nil)
	f := fizzer{b}
	bb := &bytes.Buffer{}
	c := f.Exec(bb)
	// without proper splitting we would get "'a b c'"
	r.NoError(c("echo 'a b c'"))
	r.Equal("a b c", strings.TrimSpace(bb.String()))
}
