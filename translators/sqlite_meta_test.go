package translators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanonicalizeSQLiteTable(t *testing.T) {
	assert.Equal(t, "e2e_users", canonicalizeSQLiteTable("_e2e_users_tmp"))
	assert.Equal(t, "e2e_users_tmp", canonicalizeSQLiteTable("_e2e_users_tmp_tmp"))
	assert.Equal(t, "e2e_users_tmp", canonicalizeSQLiteTable("e2e_users_tmp"))
	assert.Equal(t, "_e2e_users_tm", canonicalizeSQLiteTable("_e2e_users_tm"))
	assert.Equal(t, "_e2e_users_tmp_", canonicalizeSQLiteTable("_e2e_users_tmp_"))
}
