package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DSNはdataSourceNameを返します、もし必須の環境変数が設定されてなかった場合はerrorを返します
func DSN() (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	env := os.Getenv("ENVIRONMENT")

	if env == "dev" || env == "prd" {
		dbUser := os.Getenv("MYSQL_USERNAME")
		dbPassword := os.Getenv("MYSQL_PASSWORD")
		dbHost := os.Getenv("MYSQL_HOST")
		dbPort := os.Getenv("MYSQL_PORT")
		dbDatabase := os.Getenv("MYSQL_DATABASE")

		if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbDatabase == "" {
			return "", fmt.Errorf("ERROR : required environment variable not found")
		}
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?",
			dbUser,
			dbPassword,
			dbHost,
			dbPort,
			dbDatabase,
		) + "parseTime=true&collation=utf8mb4_bin", nil
	}

	return "", fmt.Errorf("Error : not match enviroment prd or dev , currently used %s", os.Getenv("ENVIRONMENT"))
}
