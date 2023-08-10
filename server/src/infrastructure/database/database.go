package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/Uwasaru/Sayka/config"
)

type Conn struct {
	DB *sqlx.DB
}

func NewConn() (*Conn, error) {
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
