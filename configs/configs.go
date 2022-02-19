package configs

import (
	"fmt"
	"os"
)

const (
	HASHKEY = "c5dda7a77f7dc8e29cd2d949ccc201c02e1afdd5d4a44993d2a81509d53c6954"

	MYSQL_USER     = "user123"
	MYSQL_PASSWORD = "2022Mysql"
	MYSQL_ADDRESS  = "localhost:3307"

	DB_NAME = "news_app"

	TZ = "Asia/Jakarta"
)

func GetMySqlDSN() (result string) {

	user := os.Getenv("MYSQL_USER")
	if user == "" {
		user = MYSQL_USER
	}
	password := os.Getenv("MYSQL_PASS")
	if password == "" {
		password = MYSQL_PASSWORD
	}
	address := os.Getenv("MYSQL_ADDRESS")
	if address == "" {
		address = MYSQL_ADDRESS
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = DB_NAME
	}

	result = fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, address, dbName)
	return
}
