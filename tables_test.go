package fizz_test

import (
	"testing"

	"github.com/gobuffalo/fizz"
	"github.com/stretchr/testify/require"
)

func Test_Table_Stringer(t *testing.T) {
	r := require.New(t)

	expected :=
		`create_table("users") {
	t.Column("name", "string")
	t.Column("alive", "boolean", {null: true})
	t.Column("birth_date", "timestamp", {null: true})
	t.Column("bio", "text", {null: true})
	t.Column("price", "numeric", {default: "1.00", null: true})
	t.Column("email", "string", {default: "foo@example.com", size: 50})
}`

	table := fizz.Table{
		Name: "users",
		Columns: []fizz.Column{
			fizz.Column{
				Name:    "name",
				ColType: "string",
			},
			fizz.Column{
				Name:    "alive",
				ColType: "boolean",
				Options: map[string]interface{}{
					"null": true,
				},
			},
			fizz.Column{
				Name:    "birth_date",
				ColType: "timestamp",
				Options: map[string]interface{}{
					"null": true,
				},
			},
			fizz.Column{
				Name:    "bio",
				ColType: "text",
				Options: map[string]interface{}{
					"null": true,
				},
			},
			fizz.Column{
				Name:    "price",
				ColType: "numeric",
				Options: map[string]interface{}{
					"null":    true,
					"default": "1.00",
				},
			},
			fizz.Column{
				Name:    "email",
				ColType: "string",
				Options: map[string]interface{}{
					"size":    50,
					"default": "foo@example.com",
				},
			},
		},
	}

	r.Equal(expected, table.String())
}
