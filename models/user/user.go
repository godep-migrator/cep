package user

const columnFamily = "users"

type Schema struct {
	Id gocql.UUID `json:"id"`
}
