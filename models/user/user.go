package user

import (
	"github.com/gocql/gocql"
	"github.com/thresholderio/go-processing/config/cassandra"
	"log"
	"reflect"
)

type Schema struct {
	Id   gocql.UUID `json:"id"`
	Name string     `json:"name"`
}

func FindAll() (interface{}, error) {
	iter := cassandra.Session.Query("SELECT * FROM users").Iter()
	columns := iter.Columns()

	rows := make([][]interface{}, 0)

	for {
		row := make([]interface{}, len(columns))
		for i := 0; i < len(columns); i++ {
			switch columns[i].TypeInfo.Type {
			case gocql.TypeVarchar:
				row[i] = new(string)
			case gocql.TypeBoolean:
				row[i] = new(bool)
			case gocql.TypeInt:
				row[i] = new(int)
			case gocql.TypeUUID:
				row[i] = new(gocql.UUID)
			default:
				log.Fatal("unhandled type: ", columns[i].TypeInfo)
			}
		}

		if !iter.Scan(row...) {
			break
		}

		for i := 0; i < len(columns); i++ {
			row[i] = reflect.ValueOf(row[i]).Elem().Interface()
		}

		rows = append(rows, row)
	}

	return rows, nil
}
