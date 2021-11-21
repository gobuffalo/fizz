//go:build e2e

package e2e_test

import (
	"os"
	"testing"

	"github.com/gobuffalo/pop/v6"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestSpecificSuites(t *testing.T) {
	require.NoError(t, pop.LoadConfigFile())

	switch d := os.Getenv("SODA_DIALECT"); d {
	case "postgres":
		suite.Run(t, &PostgreSQLSuite{})
	case "cockroach":
		suite.Run(t, &CockroachSuite{})
	case "mysql":
		suite.Run(t, &MySQLSuite{})
	case "sqlite":
		suite.Run(t, &SQLiteSuite{})
	default:
		t.Fatalf("Got unsupported dialect: %s", d)
	}
}
