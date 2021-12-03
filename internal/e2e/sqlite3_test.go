package e2e_test

import (
	"github.com/gobuffalo/pop/v6/logging"
	"io/ioutil"

	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/suite"
)

type SQLiteSuite struct {
	suite.Suite
}

func (s *SQLiteSuite) Test_SQLite_MigrationSteps() {
	r := s.Require()

	td, err := ioutil.TempDir("", "pop-e2e-sqlite")
	r.NoError(err)
	td = td + ".sqlite?mode=rwc&_fk=true"
	s.Suite.T().Logf("Storing SQLite DB in %s", td)

	c, err := pop.NewConnection(&pop.ConnectionDetails{URL: "sqlite3://" + td})
	r.NoError(err)
	r.NoError(c.Open(), "%s", c.URL())
	r.NoError(c.RawQuery("SELECT 1").Exec(), "%s", c.URL())

	pop.Debug = false
	pop.SetLogger(func(lvl logging.Level, s string, args ...interface{}) {
	})

	run(&s.Suite, c, runTestData(&s.Suite, c, false))
}
