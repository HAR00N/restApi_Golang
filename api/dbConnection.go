package api

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func DbConnection() *sql.DB {
	godotenv.Load(".env")

	config := mysql.Config{
		User:      os.Getenv("db_username"),
		Passwd:    os.Getenv("db_password"),
		Net:       "tcp",
		Addr:      os.Getenv("db_addr"),
		DBName:    os.Getenv("db_name"),
		Collation: "utf8mb4_general_ci",
	}
	connecter, err := mysql.NewConnector(&config)
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(connecter)
	return db
}
