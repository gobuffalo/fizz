package fizz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Table_HasColumnNoCache(t *testing.T) {
	r := require.New(t)
	table := NewTable("users", nil)
	r.NoError(table.Column("firstname", "string", nil))
	r.NoError(table.Column("lastname", "string", nil))
	table.columnsCache = map[string]struct{}{}
	r.True(table.HasColumns("firstname", "lastname"))
	r.False(table.HasColumns("age"))
}
