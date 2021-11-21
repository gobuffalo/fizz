package e2e_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func run(s *suite.Suite, c *pop.Connection, checkSchema func()) {
	r := s.Require()

	dest, err := ioutil.TempDir(os.TempDir(), "fizz-up-integration")
	r.NoError(err)

	m, err := pop.NewFileMigrator("./migrations", c)
	r.NoError(err)
	m.SchemaPath = dest

	refreshFixtures, _ := strconv.ParseBool(os.Getenv("REFRESH_FIXTURES"))

	actualFilePath := filepath.Join(dest, "schema.sql")

	// k is be the number of migrations that should run
	s.Run("direction=up", func() {
		for k := 0; k < len(m.Migrations["up"]); k++ {
			r := s.Require()

			_, err := m.UpTo(1)
			r.NoError(err)
			r.NoError(m.DumpMigrationSchema())
			expectedFilePath := filepath.Join("fixtures", c.Dialect.Name(), "up", fmt.Sprintf("%d.sql", k))

			if refreshFixtures {
				content, err := ioutil.ReadFile(actualFilePath)
				r.NoError(err)
				r.NoError(ioutil.WriteFile(expectedFilePath, content, 0666))
			} else {
				_ = expectEqualFiles(s, expectedFilePath, actualFilePath)
			}
		}
	})

	s.Run("check=schema", checkSchema)

	s.Run("direction=down", func() {
		for k := len(m.Migrations["down"]) - 1; k >= 0; k-- {
			r := s.Require()

			r.NoError(m.Down(1))
			r.NoError(m.DumpMigrationSchema())
			expectedFilePath := filepath.Join("fixtures", c.Dialect.Name(), "down", fmt.Sprintf("%d.sql", k))

			if refreshFixtures {
				content, err := ioutil.ReadFile(actualFilePath)
				r.NoError(err)
				r.NoError(ioutil.WriteFile(expectedFilePath, content, 0666))
			} else {
				_ = expectEqualFiles(s, expectedFilePath, actualFilePath)
			}
		}
	})
}

func expectEqualFiles(s *suite.Suite, expected, actual string) bool {
	A := s.Assert()
	ac, err := ioutil.ReadFile(actual)
	A.NoError(err)

	ec, err := ioutil.ReadFile(expected)
	A.NoError(err)

	A.EqualValues(
		normalizeDump(string(ec)),
		normalizeDump(string(ac)),
		`
expected file:	%s
actual file: 	%s
actual SQL dump:

%s
`, expected, actual, ac)

	return normalizeDump(string(ec)) ==
		normalizeDump(string(ac))
}

var spaces = regexp.MustCompile(`\s+`)
var comments = regexp.MustCompile("(?m)^-(.*)$")

func normalizeDump(in string) string {
	in = comments.ReplaceAllString(in, "")
	spaces.ReplaceAllString(in, " ")

	return in
}

func retryOpen(c *pop.Connection) (err error) {
	for i := 0; i <= 60; i++ {
		time.Sleep(time.Second)

		err = c.Open()
		if err != nil {
			continue
		}

		err = c.RawQuery("SELECT 1").Exec()
		if err == nil {
			return nil
		}
	}
	return err
}

func runTestData(s *suite.Suite, c *pop.Connection, supportsUUID bool) func() {
	return func() {
		r := s.Require()

		if supportsUUID {
			r.Error(c.RawQuery("INSERT INTO e2e_authors (id, created_at, updated_at) VALUES (?, ?, ?)", "78dba9f7-dd20a39f64cb", time.Now(), time.Now()).Exec(), "should fail because uuid format is not correct")
		}

		runInsertUUID(c, r)
		runForeignKeyChecks(c, r)
		runUniqueKeyChecks(c, r)
		runNotNullChecks(c, r)
	}
}

func runInsertUUID(c *pop.Connection, r *require.Assertions) {
	r.NoError(c.RawQuery("INSERT INTO e2e_authors (id, created_at, updated_at) VALUES (?, ?, ?)", "78dba9f7-81af-415e-aa2b-dd20a39f64cb", time.Now(), time.Now()).Exec())
	r.Error(c.RawQuery("INSERT INTO e2e_authors (id, created_at, updated_at) VALUES (?, ?, ?)", "78dba9f7-81af-415e-aa2b-dd20a39f64cb", time.Now(), time.Now()).Exec(), "should fail because it is a duplicate primary key")
}

func runForeignKeyChecks(c *pop.Connection, r *require.Assertions) {
	err := c.RawQuery("INSERT INTO e2e_user_posts (id, author_id, slug, published) VALUES (?,?,?, false)", "acd6abe3-38fa-4c3c-a676-933e1b06fa42", "cc1debdc-5d5a-41f3-a36b-48b1f3a03089", "slug-1").Exec()
	r.Error(err, "should fail because foreign key constraint fails")
	r.Contains(strings.ToLower(err.Error()), "foreign")

	r.NoError(c.RawQuery("INSERT INTO e2e_user_posts (id, author_id, slug, published) VALUES (?,?,?, false)", "03c6c800-54c6-40cd-89f1-7dff731f9b54", "78dba9f7-81af-415e-aa2b-dd20a39f64cb", "slug-1").Exec())

	r.NoError(c.RawQuery("INSERT INTO e2e_address (id) VALUES (?)", "d8b79b1d-e510-4763-92b3-828244b54893").Exec())
	r.NoError(c.RawQuery("INSERT INTO e2e_flow (id) VALUES (?)", "96a8b5f8-6b1b-4936-9b88-397cf0886235").Exec())
	r.NoError(c.RawQuery("INSERT INTO e2e_token (id, token, e2e_flow_id, e2e_address_id) VALUES (?, ?, ?, ?)", "5ba7f9c5-c469-439a-a07c-859f7b7b3448", "1539b471f990", "96a8b5f8-6b1b-4936-9b88-397cf0886235", "d8b79b1d-e510-4763-92b3-828244b54893").Exec())
}

func runUniqueKeyChecks(c *pop.Connection, r *require.Assertions) {
	r.NoError(c.RawQuery("INSERT INTO e2e_user_posts (id, author_id, slug) VALUES (?,?,?)", "ccbf7278-092d-4c6f-a627-84cf45233c6a", "78dba9f7-81af-415e-aa2b-dd20a39f64cb", "dupe-slug").Exec())

	err := c.RawQuery("INSERT INTO e2e_user_posts (id, author_id, slug) VALUES (?,?,?)", "ff7eb268-1640-48d3-b295-57d9a20faf3f", "78dba9f7-81af-415e-aa2b-dd20a39f64cb", "dupe-slug").Exec()
	r.Error(err, "should fail because UNIQUE constraint fails")

	message := strings.ToLower(err.Error())
	r.True(
		strings.Contains(message, "duplicate") ||
			strings.Contains(message, "unique"))
}

func runNotNullChecks(c *pop.Connection, r *require.Assertions) {
	err := c.RawQuery("INSERT INTO e2e_user_posts (id, author_id) VALUES (?,?)", "a23e6e72-08f9-412f-afb8-01f6af234eb9", "78dba9f7-81af-415e-aa2b-dd20a39f64cb").Exec()
	r.Error(err, "should fail because NOT NULL fails")
	message := strings.ToLower(err.Error())
	r.True(
		strings.Contains(message, "null") ||
			strings.Contains(message, "default"))
}
