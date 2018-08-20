package translators

import (
	"fmt"
	"strings"

	"github.com/blang/semver"

	"github.com/gobuffalo/fizz"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type mysqlTableInfo struct {
	Field   string      `db:"Field"`
	Type    string      `db:"Type"`
	Null    string      `db:"Null"`
	Key     string      `db:"Key"`
	Default interface{} `db:"Default"`
	Extra   string      `db:"Extra"`
}

func (ti mysqlTableInfo) ToColumn() fizz.Column {
	c := fizz.Column{
		Name:    ti.Field,
		ColType: ti.Type,
		Primary: ti.Key == "PRI",
		Options: map[string]interface{}{},
	}
	if strings.ToLower(ti.Null) == "yes" {
		c.Options["null"] = true
	}
	if ti.Default != nil {
		d := fmt.Sprintf("%s", ti.Default)
		c.Options["default"] = d
	}
	return c
}

var mysql57Version = semver.MustParse("5.7")
var mysql80Version = semver.MustParse("8.0")

type mysqlSchema struct {
	Schema
	version *semver.Version
}

func (p *mysqlSchema) Version() (*semver.Version, error) {
	// Use the cached version, if available.
	var err error
	if p.version != nil {
		return p.version, err
	}
	var version *semver.Version

	p.db, err = sqlx.Open("mysql", p.URL)
	if err != nil {
		return version, err
	}
	defer p.db.Close()

	res, err := p.db.Queryx("SELECT VERSION()")
	if err != nil {
		return version, err
	}

	for res.Next() {
		err = res.Scan(&version)
		p.version = version
		return version, err
	}
	return nil, errors.New("could not locate MySQL version")
}

func (p *mysqlSchema) Build() error {
	v, err := p.Version()

	if v.GTE(mysql80Version) {
		// Skip the rest, we don't need it since MySQL 8.0
		return nil
	}

	p.db, err = sqlx.Open("mysql", p.URL)
	if err != nil {
		return err
	}
	defer p.db.Close()

	res, err := p.db.Queryx(fmt.Sprintf("select TABLE_NAME as name from information_schema.TABLES where TABLE_SCHEMA = '%s'", p.Name))
	if err != nil {
		return err
	}
	for res.Next() {
		table := &fizz.Table{
			Columns: []fizz.Column{},
			Indexes: []fizz.Index{},
		}
		err = res.StructScan(table)
		if err != nil {
			return err
		}
		err = p.buildTableData(table)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *mysqlSchema) buildTableData(table *fizz.Table) error {
	prag := fmt.Sprintf("describe %s", table.Name)

	res, err := p.db.Queryx(prag)
	if err != nil {
		return nil
	}

	for res.Next() {
		ti := mysqlTableInfo{}
		err = res.StructScan(&ti)
		if err != nil {
			return err
		}
		table.Columns = append(table.Columns, ti.ToColumn())
	}

	p.schema[table.Name] = table
	return nil
}
