package orm

import (
	"encoding/json"
	"fmt"
	"testing"
)

var schemaCategory = `{
	"table_name": "TableTestName", 
	"columns": [{ "column_name":"test", "type": "integer"  },{ "column_name":"test2", "type": "varchar"  }]
}`

func TestInsertQueryPrototypeBySchema(t *testing.T) {
	c := []byte(schemaCategory)
	data := make(map[string]interface{})
	json.Unmarshal(c, &data)
	fmt.Print(InsertQueryPrototypeBySchema(data))
}

func TestCreateSchema(t *testing.T) {
	c := []byte(schemaCategory)
	data := make(map[string]interface{})
	json.Unmarshal(c, &data)
	fmt.Print(CreateSchema(data))
}
