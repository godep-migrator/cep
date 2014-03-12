package seeds

import (
  "github.com/thresholderio/go-processing/config/cassandra"
  "log"
)

func SeedUsersByFlight() {
  if err := cassandra.Session.Query("CREATE table users_by_flight (code varchar primary key, \"planning|1\" varchar, \"planning|2\" varchar, \"planning|3\" varchar, \"planning|4\" varchar, \"planning|5\" varchar, \"planning|6\" varchar, \"planning|7\" varchar)").Exec(); err != nil {
    log.Println(err)
  }

  if err := cassandra.Session.Query("INSERT into users_by_flight (code, \"planning|1\", \"planning|2\", \"planning|3\", \"planning|4\", \"planning|5\", \"planning|6\", \"planning|7\") values (?, ?, ?, ?, ?, ?, ?, ?)", "vx247", "crand", "jchao", "yliang", "jlai", "elee", "msilva", "clauver").Exec(); err != nil {
    log.Println(err)
  }
}
