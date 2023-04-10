package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/muratovdias/diplom/internal/app/config"
)

func IniitDB(config *config.DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.User, config.Pwd, config.Host, config.Port, config.Name)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db, err
}

// func CreateTables(db *sqlx.DB) error {
// 	tables := []string{users, trainerSchedule, clientSchedule}
// 	for _, table := range tables {
// 		_, err := db.Exec(table)
// 		if err != nil {
// 			return fmt.Errorf("failed creating tables: %w", err)
// 		}
// 	}
// 	return nil
// }
