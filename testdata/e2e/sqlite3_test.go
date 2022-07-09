package e2e_test

import (
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

	c, err := pop.NewConnection(&pop.ConnectionDetails{URL: "sqlite3://" + td + "db.sql?mode=rwc&_fk=true"})
	r.NoError(err)
	r.NoError(c.Open(), "%s", c.URL())
	r.NoError(c.RawQuery("SELECT 1").Exec(), "%s", c.URL())

	run(&s.Suite, c, runTestData(&s.Suite, c, false))
}
