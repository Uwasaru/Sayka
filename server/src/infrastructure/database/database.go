package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Uwasaru/Sayka/config"
)

func NewDB() (*sql.DB, error) {
	dbDSN, err := config.DSN()
	if err != nil {
		return nil, err
	}
	fmt.Println(dbDSN)

	db, err := sql.Open("mysql", dbDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL : %w", err)
	}

	err = db.Ping()

	if err != nil {
		log.Println("db connect error ")
		panic(err)
	}
	return db, err
}
