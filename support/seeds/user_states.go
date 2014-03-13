package seeds

import (
	"github.com/gocql/gocql"
	"github.com/thresholderio/go-processing/config/cassandra"
	"log"
)

func SeedUserStates() {
	cassandra.Session.Query("DROP TABLE user_states").Exec()

	if err := cassandra.Session.Query("CREATE table IF NOT EXISTS user_states (id uuid, name varchar, primary key (id, name))").Exec(); err != nil {
		log.Println(err)
	}

	user_states := []string{"planning", "delayed", "boarded"}

	for _, state := range user_states {
		if uuid, err := gocql.RandomUUID(); err != nil {
			log.Fatal(err)
		} else {
			cassandra.Session.Query("INSERT INTO user_states (id, name) values (?, ?)", uuid, state).Exec()
		}
	}
}
