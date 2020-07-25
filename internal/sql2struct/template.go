package sql2struct

import (
	"fmt"
	"os"
	"text/template"

	"github.com/xiezeyu-99/go-programming-tour-book/tour/internal/word"
)

const structTpl = `type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}}{{$length:=len .Comment}}{{if gt $length 0}}// {{.Comment}}{{else}}//{{.Name}}{{end}}
	{{$typeLen:=len .Type}}{{if gt $typeLen 0}}{{.Name | ToCamelCase}} {{.Type}} {{.Tag}}{{else}}{{.Name | ToCamelCase}}{{end}}
	{{end}}
}
					
func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

//数据库字段转结构体
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, columns := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    columns.ColumnName,
			Type:    DBTypeToStructType[columns.DataType],
			Tag:     fmt.Sprintf("`json:"+`"%s"`+"`", columns.ColumnName),
			Comment: columns.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	return err
}
