package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Connect abre a conexão com o banco de dados e a retorna
func Connect() (*sql.DB, error) {
	//abrindo a conexão com o banco de dados
	db, err := sql.Open("mysql", config.DataBaseConnectionString)
	if err != nil {
		return nil, err
	}

	// testando a conexão como banco
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
