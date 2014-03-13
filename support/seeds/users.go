package seeds

import (
	"github.com/thresholderio/go-processing/config/cassandra"
  "github.com/gocql/gocql"
	"log"
)

func SeedUsers() {
	if err := cassandra.Session.Query("CREATE table IF NOT EXISTS users (id uuid primary key, name varchar)").Exec(); err != nil {
		log.Println(err)
	}

  users := []string{"crand", "jchao", "elee", "msilva", "jlai", "yliang", "clauver"}

  for _, user := range users {
    if uuid, err := gocql.RandomUUID(); err != nil {
      log.Fatal(err)
    } else {
      cassandra.Session.Query("INSERT INTO users (id, name) values (?, ?)", uuid, user).Exec()
    }
  }
}
