package models

type User struct {
	tableName struct{} `pg:"users"`
	ID        string   `pg:",pk"`
	Name      string
	Online    bool
}
