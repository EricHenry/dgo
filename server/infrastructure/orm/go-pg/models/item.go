package models

type Item struct {
	tableName   struct{} `pg:"items"`
	ID          string   `pg:",pk"`
	Name        string
	Description string
	Price       float32
}
