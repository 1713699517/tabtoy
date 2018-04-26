package json

// 报错行号+3
const templateText = `{
	"@Tool": "github.com/davyxu/tabtoy",
	"@Version": "{{.Version}}",	{{range $di, $datatab := .Datas}}
	"{{$datatab.Name}}":[ {{range $row,$rowData := $datatab.Rows}}
		{ {{range $col, $headType := $datatab.Header}}"{{$headType.FieldName}}": {{WrapJsonValue $ $datatab $row $col}}{{GenJsonTailComma $col $datatab.Header}} {{end}}}{{GenJsonTailComma $row $datatab.Rows}}{{end}} 
	]{{end}}
}`
