package e2e_test

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/suite"
)

type PostgreSQLSuite struct {
	suite.Suite
}

func (s *PostgreSQLSuite) Test_PostgreSQL_MigrationSteps() {
	r := s.Require()

	c, err := pop.Connect("postgres")
	r.NoError(err)
	r.NoError(retryOpen(c))

	run(&s.Suite, c, runTestData(&s.Suite, c, true))
}
