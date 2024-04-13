package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //import implicito. Quem utiliza o pacote é o pacote sql
)

const (
	DB_NAME = "devbook"
	DB_HOST = "localhost"
	DB_USER = "golang"
	DB_PASS = "golang"
	DB_PORT = "3306"
)

func Conectar() (*sql.DB, error) {
	stringConexao := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro := db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
