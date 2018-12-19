package fizz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

// Table is the table definition for fizz.
type Table struct {
	Name         string `db:"name"`
	Columns      []Column
	Indexes      []Index
	ForeignKeys  []ForeignKey
	Options      map[string]interface{}
	columnsCache map[string]struct{}
}

func (t Table) String() string {
	return t.Fizz()
}

// Fizz returns the fizz DDL to create the table.
func (t Table) Fizz() string {
	var buff bytes.Buffer
	timestampsOpt := t.Options["timestamps"].(bool)
	// Write table options
	o := make([]string, 0, len(t.Options))
	for k, v := range t.Options {
		// Special handling for timestamps option
		if k == "timestamps" {
			continue
		}
		vv, _ := json.Marshal(v)
		o = append(o, fmt.Sprintf("%s: %s", k, string(vv)))
	}
	if len(o) > 0 {
		sort.SliceStable(o, func(i, j int) bool { return o[i] < o[j] })
		buff.WriteString(fmt.Sprintf("create_table(\"%s\", {%s}) {\n", t.Name, strings.Join(o, ", ")))
	} else {
		buff.WriteString(fmt.Sprintf("create_table(\"%s\") {\n", t.Name))
	}
	// Write columns
	if timestampsOpt {
		for _, c := range t.Columns {
			if c.Name == "created_at" || c.Name == "updated_at" {
				continue
			}
			buff.WriteString(fmt.Sprintf("\t%s\n", c.String()))
		}
	} else {
		for _, c := range t.Columns {
			buff.WriteString(fmt.Sprintf("\t%s\n", c.String()))
		}
	}
	if timestampsOpt {
		buff.WriteString("\tt.Timestamps()\n")
	}
	// Write indexes
	for _, i := range t.Indexes {
		buff.WriteString(fmt.Sprintf("\t%s\n", i.String()))
	}
	buff.WriteString("}")
	return buff.String()
}

// UnFizz returns the fizz DDL to remove the table.
func (t Table) UnFizz() string {
	return fmt.Sprintf("drop_table(\"%s\")", t.Name)
}

func (t *Table) DisableTimestamps() {
	t.Options["timestamps"] = false
}

// Column adds a column to the table definition.
func (t *Table) Column(name string, colType string, options Options) error {
	if _, found := t.columnsCache[name]; found {
		return fmt.Errorf("duplicated column %s", name)
	}
	var primary bool
	if _, ok := options["primary"]; ok {
		primary = true
	}
	c := Column{
		Name:    name,
		ColType: colType,
		Options: options,
		Primary: primary,
	}
	t.columnsCache[name] = struct{}{}
	// Ensure id is first
	if name == "id" {
		t.Columns = append([]Column{c}, t.Columns...)
	} else {
		t.Columns = append(t.Columns, c)
	}
	return nil
}

// ForeignKey adds a new foreign key to the table definition.
func (t *Table) ForeignKey(column string, refs interface{}, options Options) error {
	fkr, err := parseForeignKeyRef(refs)
	if err != nil {
		return errors.WithStack(err)
	}
	fk := ForeignKey{
		Column:     column,
		References: fkr,
		Options:    options,
	}

	if options["name"] != nil {
		fk.Name = options["name"].(string)
	} else {
		fk.Name = fmt.Sprintf("%s_%s_%s_fk", t.Name, fk.References.Table, strings.Join(fk.References.Columns, "_"))
	}

	t.ForeignKeys = append(t.ForeignKeys, fk)
	return nil
}

// Index adds a new index to the table definition.
func (t *Table) Index(columns interface{}, options Options) error {
	i := Index{}
	switch tp := columns.(type) {
	default:
		return errors.Errorf("unexpected type %T for %s index columns", tp, t.Name) // %T prints whatever type t has
	case string:
		i.Columns = []string{tp}
	case []string:
		if len(tp) == 0 {
			return errors.Errorf("expected at least one column to apply %s index", t.Name)
		}
		i.Columns = tp
	case []interface{}:
		if len(tp) == 0 {
			return errors.Errorf("expected at least one column to apply %s index", t.Name)
		}
		cl := make([]string, len(tp))
		for i, c := range tp {
			cl[i] = c.(string)
		}
		i.Columns = cl
	}
	if options["name"] != nil {
		i.Name = options["name"].(string)
	} else {
		i.Name = fmt.Sprintf("%s_%s_idx", t.Name, strings.Join(i.Columns, "_"))
	}
	i.Unique = options["unique"] != nil && options["unique"].(bool)
	t.Indexes = append(t.Indexes, i)
	return nil
}

// Timestamp is a shortcut to add a timestamp column with default options.
func (t *Table) Timestamp(name string) error {
	return t.Column(name, "timestamp", Options{})
}

// Timestamps adds created_at and updated_at columns to the Table definition.
func (t *Table) Timestamps() error {
	if err := t.Timestamp("created_at"); err != nil {
		return err
	}
	return t.Timestamp("updated_at")
}

// ColumnNames returns the names of the Table's columns.
func (t *Table) ColumnNames() []string {
	cols := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		cols[i] = c.Name
	}
	return cols
}

// HasColumns checks if the Table has all the given columns.
func (t *Table) HasColumns(args ...string) bool {
	for _, a := range args {
		if _, ok := t.columnsCache[a]; !ok {
			return false
		}
	}
	return true
}

// NewTable creates a new Table.
func NewTable(name string, opts map[string]interface{}) Table {
	if opts == nil {
		opts = make(map[string]interface{})
	}
	// auto-timestamp as default
	if enabled, exists := opts["timestamps"]; !exists || enabled == true {
		opts["timestamps"] = true
	}
	return Table{
		Name:         name,
		Columns:      []Column{},
		Indexes:      []Index{},
		Options:      opts,
		columnsCache: map[string]struct{}{},
	}
}

func (f fizzer) CreateTable(name string, opts map[string]interface{}, help plush.HelperContext) error {
	t := NewTable(name, opts)
	if help.HasBlock() {
		ctx := help.Context.New()
		ctx.Set("t", &t)
		if _, err := help.BlockWith(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	if t.Options["timestamps"].(bool) {
		if !t.HasColumns("created_at", "updated_at") {
			t.Timestamps()
		}
	}

	f.add(f.Bubbler.CreateTable(t))
	return nil
}

func (f fizzer) DropTable(name string) {
	f.add(f.Bubbler.DropTable(Table{Name: name}))
}

func (f fizzer) RenameTable(old, new string) {
	f.add(f.Bubbler.RenameTable([]Table{
		{Name: old},
		{Name: new},
	}))
}
