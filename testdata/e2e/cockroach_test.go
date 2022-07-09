package e2e_test

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/suite"
)

type CockroachSuite struct {
	suite.Suite
}

func (s *CockroachSuite) Test_Cockroach_MigrationSteps() {
	r := s.Require()

	c, err := pop.Connect("cockroach")
	r.NoError(err)
	r.NoError(retryOpen(c))

	run(&s.Suite, c, runTestData(&s.Suite, c, true))
}
