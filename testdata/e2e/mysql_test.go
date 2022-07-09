package e2e_test

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/suite"
)

type MySQLSuite struct {
	suite.Suite
}

func (s *MySQLSuite) Test_MySQL_MigrationSteps() {
	r := s.Require()

	c, err := pop.Connect("mysql")
	r.NoError(err)
	r.NoError(retryOpen(c))

	run(&s.Suite, c, runTestData(&s.Suite, c, false))
}
