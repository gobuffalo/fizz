package translators_test

import (
	"fmt"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/fizz"
	"github.com/gobuffalo/fizz/translators"
)

var _ fizz.Translator = (*translators.MariaDB)(nil)
var mat = translators.NewMariaDB("", "")

func init() {
	u := "%s:%s@(%s:%s)/%s?parseTime=true&multiStatements=true&readTimeout=1s&collation=%s"
	u = fmt.Sprintf(u, envy.Get("MYSQL_USER", "root"), envy.Get("MYSQL_PASSWORD", ""), envy.Get("MYSQL_HOST", "127.0.0.1"), envy.Get("MYSQL_PORT", "3306"), "pop_test", "utf8mb4_general_ci")
	mat = translators.NewMariaDB(u, "pop_test")
}

func (p *MariaDBSuite) Test_MySQL_SchemaMigration() {
	r := p.Require()
	ddl := `CREATE TABLE ` + "`schema_migrations`" + ` (
` + "`version`" + ` VARCHAR (191) NOT NULL
) ENGINE=InnoDB;
CREATE UNIQUE INDEX ` + "`version_idx`" + ` ON ` + "`schema_migrations`" + ` (` + "`version`" + `);`
	res, err := mat.CreateTable(fizz.Table{
		Name: "schema_migrations",
		Columns: []fizz.Column{
			{Name: "version", ColType: "string"},
		},
		Indexes: []fizz.Index{
			{Name: "version_idx", Columns: []string{"version"}, Unique: true},
		},
	})
	r.NoError(err)
	r.Equal(ddl, res)
}
