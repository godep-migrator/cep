package seeds

import (
	"github.com/thresholderio/go-processing/config/cassandra"
	"log"
)

func SeedFlights() {
  cassandra.Session.Query("DROP TABLE flights").Exec()
	if err := cassandra.Session.Query("CREATE table IF NOT EXISTS flights (code varchar, airline varchar, primary key (code, airline))").Exec(); err != nil {
		log.Println(err)
	}

  flights := map[string]string{
    "vx-1": "Virgin America",
    "vx-2": "Virgin America",
    "vx-3": "Virgin America",
    "b6-1": "Jet Blue",
    "ua-1": "United",
    "aa-1": "American Airlines",
    "dl-1": "Delta",
    "as-1": "Alaska Airlines",
    "ha-1": "Hawaii Airlines",
  }

  for code, airline := range flights {
    cassandra.Session.Query("INSERT INTO flights (code, airline) values (?, ?)", code, airline).Exec()
  }
}

func SeedUsersByFlight() {
  cassandra.Session.Query("DROP TABLE users_by_flight").Exec()

  if err := cassandra.Session.Query("CREATE TABLE IF NOT EXISTS users_by_flight (flight_code varchar, user_name varchar, primary key (flight_code, user_name))").Exec(); err != nil {
    log.Fatal(err)
  }

  users1 := []string{"crand", "jchao", "elee", "msilva"}
  users2 := []string{"yliang", "jlai", "clauver"}

  for _, user := range users1 {
    cassandra.Session.Query("INSERT INTO users_by_flight (flight_code, user_name) values (?, ?)", "vx-1", user).Exec()
  }

  for _, user := range users2 {
    cassandra.Session.Query("INSERT INTO users_by_flight (flight_code, user_name) values (?, ?)", "vx-2", user).Exec()
  }
}
