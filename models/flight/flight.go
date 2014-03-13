package flight

import (
	"github.com/gocql/gocql"
	"log"
)

type Schema struct {
  Id gocql.UUID `json:"id"`
  Code string `json:"code"`
}
