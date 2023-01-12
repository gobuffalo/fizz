package fizz

import (
	"os"
	"strings"

	"github.com/gobuffalo/plush/v4"
)

type BubbleType int

type Bubbler struct {
	Translator
	data        []string
	DisableExec bool
}

func NewBubbler(t Translator) *Bubbler {
	return &Bubbler{
		Translator:  t,
		data:        []string{},
		DisableExec: false,
	}
}

func NewBubblerWithDisabledExec(t Translator) *Bubbler {
	return &Bubbler{
		Translator:  t,
		data:        []string{},
		DisableExec: true,
	}
}

func (b *Bubbler) String() string {
	return strings.Join(b.data, "\n")
}

func (b *Bubbler) Bubble(s string) (string, error) {
	f := fizzer{
		Bubbler:     b,
		DisableExec: b.DisableExec,
	}
	ctx := plush.NewContextWith(map[string]interface{}{
		"exec":             f.Exec(os.Stdout),
		"create_table":     f.CreateTable,
		"change_column":    f.ChangeColumn,
		"add_column":       f.AddColumn,
		"drop_column":      f.DropColumn,
		"rename_column":    f.RenameColumn,
		"raw":              f.RawSQL,
		"sql":              f.RawSQL,
		"add_index":        f.AddIndex,
		"drop_index":       f.DropIndex,
		"rename_index":     f.RenameIndex,
		"add_foreign_key":  f.AddForeignKey,
		"drop_foreign_key": f.DropForeignKey,
		"drop_table":       f.DropTable,
		"rename_table":     f.RenameTable,
	})

	err := plush.RunScript(s, ctx)

	return b.String(), err
}
