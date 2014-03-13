package seeds

import (
	"github.com/thresholderio/go-processing/config/cassandra"
	"log"
)

func SeedUsers() {
	if err := cassandra.Session.Query("CREATE table IF NOT EXISTS users (id int primary key, name varchar)").Exec(); err != nil {
		log.Println(err)
	}

	cassandra.Session.Query("INSERT into users (id, name) values (?, ?)", 1, "crand").Exec()
	cassandra.Session.Query("INSERT into users (id, name) values (?, ?)", 2, "jchao").Exec()
	cassandra.Session.Query("INSERT into users (id, name) values (?, ?)", 3, "yliang").Exec()
	cassandra.Session.Query("INSERT into users (id, name) values (?, ?)", 4, "jlai").Exec()
	cassandra.Session.Query("INSERT into users (id, name) values (?, ?)", 5, "elee").Exec()
	cassandra.Session.Query("INSERT into users (id, name) values (?, ?)", 6, "msilva").Exec()
	cassandra.Session.Query("INSERT into users (id, name) values (?, ?)", 7, "clauver").Exec()
}
