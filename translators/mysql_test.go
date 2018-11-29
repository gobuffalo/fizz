package translators_test

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Load MySQL Go driver
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/fizz"
	"github.com/gobuffalo/fizz/translators"
)

var _ fizz.Translator = (*translators.MySQL)(nil)
var myt = translators.NewMySQL("", "")

func init() {
	u := "%s:%s@(%s:%s)/%s?parseTime=true&multiStatements=true&readTimeout=1s&collation=%s"
	u = fmt.Sprintf(u, envy.Get("MYSQL_USER", "root"), envy.Get("MYSQL_PASSWORD", "root"), envy.Get("MYSQL_HOST", "127.0.0.1"), envy.Get("MYSQL_PORT", "3306"), "pop_test", "utf8_general_ci")
	myt = translators.NewMySQL(u, "pop_test")
}

func (p *MySQLSuite) Test_MySQL_SchemaMigration() {
	r := p.Require()
	ddl := `CREATE TABLE ` + "`schema_migrations`" + ` (
` + "`version`" + ` VARCHAR (255) NOT NULL
) ENGINE=InnoDB;
CREATE UNIQUE INDEX ` + "`version_idx`" + ` ON ` + "`schema_migrations`" + ` (` + "`version`" + `);`

	res, err := myt.CreateTable(fizz.Table{
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

func (p *MySQLSuite) Test_MySQL_CreateTable() {
	r := p.Require()
	ddl := `CREATE TABLE ` + "`users`" + ` (
` + "`id`" + ` INTEGER NOT NULL AUTO_INCREMENT,
PRIMARY KEY(` + "`id`" + `),
` + "`first_name`" + ` VARCHAR (255) NOT NULL,
` + "`last_name`" + ` VARCHAR (255) NOT NULL,
` + "`email`" + ` VARCHAR (20) NOT NULL,
` + "`permissions`" + ` text,
` + "`age`" + ` INTEGER DEFAULT 40,
` + "`raw`" + ` BLOB NOT NULL,
` + "`json`" + ` JSON NOT NULL,
` + "`float`" + ` FLOAT(5) NOT NULL,
` + "`integer`" + ` INTEGER NOT NULL,
` + "`bytes`" + ` BLOB NOT NULL,
` + "`created_at`" + ` DATETIME NOT NULL,
` + "`updated_at`" + ` DATETIME NOT NULL
) ENGINE=InnoDB;`

	res, err := fizz.AString(`
	create_table("users") {
		t.Column("id", "integer", {"primary": true})
		t.Column("first_name", "string", {})
		t.Column("last_name", "string", {})
		t.Column("email", "string", {"size":20})
		t.Column("permissions", "text", {"null": true})
		t.Column("age", "integer", {"null": true, "default": 40})
		t.Column("raw", "blob", {})
		t.Column("json", "json", {})
		t.Column("float", "float", {"precision": 5})
		t.Column("integer", "integer", {})
		t.Column("bytes", "[]byte", {})
	}
	`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_CreateTable_UUID() {
	r := p.Require()
	ddl := `CREATE TABLE ` + "`users`" + ` (
` + "`first_name`" + ` VARCHAR (255) NOT NULL,
` + "`last_name`" + ` VARCHAR (255) NOT NULL,
` + "`email`" + ` VARCHAR (20) NOT NULL,
` + "`permissions`" + ` text,
` + "`age`" + ` INTEGER DEFAULT 40,
` + "`company_id`" + ` char(36) NOT NULL DEFAULT 'test',
` + "`uuid`" + ` char(36) NOT NULL,
PRIMARY KEY(` + "`uuid`" + `),
` + "`created_at`" + ` DATETIME NOT NULL,
` + "`updated_at`" + ` DATETIME NOT NULL
) ENGINE=InnoDB;`

	res, err := fizz.AString(`
	create_table("users") {
		t.Column("first_name", "string", {})
		t.Column("last_name", "string", {})
		t.Column("email", "string", {"size":20})
		t.Column("permissions", "text", {"null": true})
		t.Column("age", "integer", {"null": true, "default": 40})
		t.Column("company_id", "uuid", {"default_raw": "'test'"})
		t.Column("uuid", "uuid", {"primary": true})
	}
	`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_CreateTables_WithForeignKeys() {
	r := p.Require()
	ddl := `CREATE TABLE ` + "`users`" + ` (
` + "`id`" + ` INTEGER NOT NULL AUTO_INCREMENT,
PRIMARY KEY(` + "`id`" + `),
` + "`email`" + ` VARCHAR (20) NOT NULL,
` + "`created_at`" + ` DATETIME NOT NULL,
` + "`updated_at`" + ` DATETIME NOT NULL
) ENGINE=InnoDB;
CREATE TABLE ` + "`profiles`" + ` (
` + "`id`" + ` INTEGER NOT NULL AUTO_INCREMENT,
PRIMARY KEY(` + "`id`" + `),
` + "`user_id`" + ` INTEGER NOT NULL,
` + "`first_name`" + ` VARCHAR (255) NOT NULL,
` + "`last_name`" + ` VARCHAR (255) NOT NULL,
` + "`created_at`" + ` DATETIME NOT NULL,
` + "`updated_at`" + ` DATETIME NOT NULL,
FOREIGN KEY (` + "`user_id`" + `) REFERENCES ` + "`users`" + ` (` + "`id`" + `)
) ENGINE=InnoDB;`

	res, err := fizz.AString(`
	create_table("users") {
		t.Column("id", "INT", {"primary": true})
		t.Column("email", "string", {"size":20})
	}
	create_table("profiles") {
		t.Column("id", "INT", {"primary": true})
		t.Column("user_id", "INT", {})
		t.Column("first_name", "string", {})
		t.Column("last_name", "string", {})
		t.ForeignKey("user_id", {"users": ["id"]}, {})
	}
	`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_DropTable() {
	r := p.Require()

	ddl := `DROP TABLE ` + "`users`" + `;`

	res, err := fizz.AString(`drop_table("users")`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_RenameTable() {
	r := p.Require()

	ddl := `ALTER TABLE ` + "`users`" + ` RENAME TO ` + "`people`" + `;`

	res, err := fizz.AString(`rename_table("users", "people")`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_RenameTable_NotEnoughValues() {
	r := p.Require()

	_, err := myt.RenameTable([]fizz.Table{})
	r.Error(err)
}

func (p *MySQLSuite) Test_MySQL_ChangeColumn() {
	r := p.Require()
	ddl := `ALTER TABLE ` + "`users`" + ` MODIFY ` + "`mycolumn`" + ` VARCHAR (50) NOT NULL DEFAULT 'foo';`

	res, err := fizz.AString(`change_column("users", "mycolumn", "string", {"default": "foo", "size": 50})`, myt)
	r.NoError(err)

	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_AddColumn() {
	r := p.Require()
	ddl := `ALTER TABLE ` + "`users`" + ` ADD COLUMN ` + "`mycolumn`" + ` VARCHAR (50) NOT NULL DEFAULT 'foo';`

	res, err := fizz.AString(`add_column("users", "mycolumn", "string", {"default": "foo", "size": 50})`, myt)
	r.NoError(err)

	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_AddColumnAfter() {
	r := p.Require()
	ddl := `ALTER TABLE ` + "`users`" + ` ADD COLUMN ` + "`mycolumn`" + ` VARCHAR (50) NOT NULL DEFAULT 'foo' AFTER ` + "`created_at`" + `;`

	res, err := fizz.AString(`add_column("users", "mycolumn", "string", {"default": "foo", "size": 50, "after":"created_at"})`, myt)
	r.NoError(err)

	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_AddColumnFirst() {
	r := p.Require()
	ddl := `ALTER TABLE ` + "`users`" + ` ADD COLUMN ` + "`mycolumn`" + ` VARCHAR (50) NOT NULL DEFAULT 'foo' FIRST;`

	res, err := fizz.AString(`add_column("users", "mycolumn", "string", {"default": "foo", "size": 50, "first":true})`, myt)
	r.NoError(err)

	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_DropColumn() {
	r := p.Require()
	ddl := `ALTER TABLE ` + "`users`" + ` DROP COLUMN ` + "`mycolumn`" + `;`

	res, err := fizz.AString(`drop_column("users", "mycolumn")`, myt)
	r.NoError(err)

	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_RenameColumn() {
	r := p.Require()
	ddl := `ALTER TABLE ` + "`users`" + ` CHANGE ` + "`email`" + ` ` + "`email_address`" + ` varchar(50) NOT NULL DEFAULT 'foo@example.com';`

	res, err := fizz.AString(`rename_column("users", "email", "email_address")`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_AddIndex() {
	r := p.Require()
	ddl := `CREATE INDEX ` + "`users_email_idx`" + ` ON ` + "`users`" + ` (` + "`email`" + `);`

	res, err := fizz.AString(`add_index("users", "email", {})`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_AddIndex_Unique() {
	r := p.Require()
	ddl := `CREATE UNIQUE INDEX ` + "`users_email_idx`" + ` ON ` + "`users`" + ` (` + "`email`" + `);`

	res, err := fizz.AString(`add_index("users", "email", {"unique": true})`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_AddIndex_MultiColumn() {
	r := p.Require()
	ddl := `CREATE INDEX ` + "`users_id_email_idx`" + ` ON ` + "`users`" + ` (` + "`id`" + `, ` + "`email`" + `);`

	res, err := fizz.AString(`add_index("users", ["id", "email"], {})`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_AddIndex_CustomName() {
	r := p.Require()
	ddl := `CREATE INDEX ` + "`email_index`" + ` ON ` + "`users`" + ` (` + "`email`" + `);`

	res, err := fizz.AString(`add_index("users", "email", {"name": "email_index"})`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_DropIndex() {
	r := p.Require()
	ddl := `DROP INDEX ` + "`email_idx`" + ` ON ` + "`users`" + `;`

	res, err := fizz.AString(`drop_index("users", "email_idx")`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_RenameIndex() {
	r := p.Require()

	ddl := `ALTER TABLE ` + "`users`" + ` RENAME INDEX ` + "`email_idx`" + ` TO ` + "`email_address_idx`" + `;`

	res, err := fizz.AString(`rename_index("users", "email_idx", "email_address_idx")`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_AddForeignKey() {
	r := p.Require()
	ddl := `ALTER TABLE ` + "`profiles`" + ` ADD CONSTRAINT ` + "`profiles_users_id_fk`" + ` FOREIGN KEY (` + "`user_id`" + `) REFERENCES ` + "`users`" + ` (` + "`id`" + `);`

	res, err := fizz.AString(`add_foreign_key("profiles", "user_id", {"users": ["id"]}, {})`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}

func (p *MySQLSuite) Test_MySQL_DropForeignKey() {
	r := p.Require()
	ddl := `ALTER TABLE ` + "`profiles`" + ` DROP FOREIGN KEY  ` + "`profiles_users_id_fk`" + `;`

	res, err := fizz.AString(`drop_foreign_key("profiles", "profiles_users_id_fk", {})`, myt)
	r.NoError(err)
	r.Equal(ddl, res)
}
