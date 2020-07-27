package gopg

import (
	"github.com/erichenry/dgo/infrastructure/orm/go-pg/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func NewDB(connectionString string) (*pg.DB, error) {
	opt, err := pg.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}
	db := pg.Connect(opt)

	return db, nil
}

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.Item)(nil),
		(*models.User)(nil),
	}

	for _, model := range models {
		m := db.Model(model)
		err := m.CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
