package fizz

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Exec(t *testing.T) {
	r := require.New(t)

	b := NewBubbler(nil)
	f := fizzer{b}
	bb := &bytes.Buffer{}
	c := f.Exec(bb)
	c("echo hello")
	r.Equal("hello\n", bb.String())
}

func Test_ExecQuoted(t *testing.T) {
	r := require.New(t)

	b := NewBubbler(nil)
	f := fizzer{b}
	bb := &bytes.Buffer{}
	c := f.Exec(bb)
	// without proper splitting we would get "'a b c'"
	c("echo 'a b c'")
	r.Equal("a b c\n", bb.String())
}
