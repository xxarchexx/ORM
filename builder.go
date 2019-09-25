package orm

import (
	"bytes"
	"fmt"
	"html/template"
)

// type TableSchema struct {
// 	Schema map[string]interface{}
// }

// type TableSchemaArray struct {
// 	TablesSchemas []TableSchema
// }

// // func buildQuery(schema) string {
// // 	var t *template.Template
// // 	template.New("query")
// // }

//MapPostgresType do map to database type from go type for sql create table query
func MapPostgresType(columntype string) (restype string) {
	switch columntype {
	case "int":
		restype = "interger"
	case "string":
		restype = "text"
	default:
		restype = columntype
	}
	return
}

//CreateInsertParams create params array for query to db
func CreateInsertParams(data map[string]interface{}) (args []interface{}) {
	args = []interface{}{}
	cols, ok := data["columns"].([]interface{})
	if !ok {
		return nil
	}

	for _, val := range cols {
		args = append(args, val)
	}

	return
}

//InsertQueryPrototypeBySchema build  string from schema
func InsertQueryPrototypeBySchema(mapSchema map[string]interface{}) string {

	var insertQuery = `
		Insert Into {{.table_name}} {{ .empty}}{{range .columns}} {{.column_name }}{{end }}	values {{ .empty -}}
		({{- range .columns}}
				{{- comma_func}}${{ii}}
			{{- end}})`

	var tmp *template.Template

	var inc = 1

	tmp, _ = template.New("query").Funcs(template.FuncMap{
		"ii": func() (res string) {
			if inc == 1 {
				res = fmt.Sprintf("%v", inc)
			} else {
				res = fmt.Sprint(inc)
			}

			inc = inc + 1
			return
		},
		"comma_func": func() (res string) {
			if inc == 1 {
				res = ""
			} else {
				res = ","
			}

			return
		},
	}).Parse(insertQuery)

	var buf bytes.Buffer
	tmp.Execute(&buf, mapSchema)
	return string(buf.Bytes())
}

//CreateSchema create sql query for postgres
func CreateSchema(mapSchema map[string]interface{}) (result int) {

	result = 0

	var createTemplate = `
		Create Table {{.table_name}} {{.empty -}}
		(
			{{- range .columns }}
				 {{- .column_name}} {{ get_column_type .type}},{{.empty -}} 
			{{end -}}	
		)	
	`
	var tmp *template.Template
	tmp, _ = template.New("New").Funcs(template.FuncMap{
		"get_column_type": MapPostgresType,
	}).Parse(createTemplate)
	var buf bytes.Buffer
	tmp.Execute(&buf, mapSchema)
	fmt.Print(string(buf.Bytes()))

	// var tm2 = `{{printf "%v" "\""}}`

	// // Create a new template and parse the letter into it.

	// tpl, _ = template.New("New").Parse(createTemplate)

	// m := make(map[string]string)
	// m["table_name"] = "test"

	// const (
	// 	master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
	// 	overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	// )
	// var (
	// 	funcs     = template.FuncMap{"join": strings.Join}
	// 	guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	// )
	// masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
	// 	log.Fatal(err)
	// }
	//	t.Execute(os.Stdout, "OK")

	// for key, value := range data {
	// 	if _, ok := value.(map[string]interface{}); ok {
	// 		log.Printf("{ \"%v\":", key)
	// 		// dumpMap(str+"\t", value)
	// 		log.Printf("}\n")
	// 	} else {

	// 		log.Printf("%v : %v", key, value)
	// 	}
	// }
	return 0
}
