package seeds

import (
	"github.com/gocql/gocql"
	"github.com/jeffchao/cep/config/cassandra"
	"log"
)

func SeedUsers() {
	cassandra.Session.Query("DROP TABLE users").Exec()

	if err := cassandra.Session.Query("CREATE table IF NOT EXISTS users (id uuid, name varchar, primary key (id, name))").Exec(); err != nil {
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
