package seeds

import (
	"github.com/thresholderio/go-processing/config/cassandra"
	"log"
)

func SeedFlights() {
	if err := cassandra.Session.Query("CREATE table IF NOT EXISTS flights (code varchar primary key, airline varchar)").Exec(); err != nil {
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
	if err := cassandra.Session.Query("CREATE table IF NOT EXISTS users_by_flight (code varchar primary key, \"planning|1\" varchar, \"planning|2\" varchar, \"planning|3\" varchar, \"boarded|4\" varchar, \"planning|5\" varchar, \"planning|6\" varchar, \"boarded|7\" varchar)").Exec(); err != nil {
		log.Println(err)
	}

	if err := cassandra.Session.Query("INSERT into users_by_flight (code, \"planning|1\", \"planning|2\", \"planning|3\", \"boarded|4\") values (?, ?, ?, ?, ?)", "vx247", "crand", "jchao", "yliang", "jlai").Exec(); err != nil {
		log.Println(err)
	}

	if err := cassandra.Session.Query("INSERT into users_by_flight (code, \"planning|5\", \"planning|6\", \"boarded|7\") values (?, ?, ?, ?)", "vx111", "elee", "msilva", "clauver").Exec(); err != nil {
		log.Println(err)
	}
}
