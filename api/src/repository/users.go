package repository

import (
	"api/src/models"
	"database/sql"
)

// Users representa um repositório de usuários
type Users struct {
	db *sql.DB
}

// NewUsersRepository cria um repositório de usuários
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// CreateUser insere um usuário no banco de dados
func (repository Users) CreateUser(user models.User) (uint64, error) {

	// query
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, email, email_backup, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//executando a query
	result, err := statement.Exec(user.Name, user.Email, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	// pegando o último id inserido
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

// EditUser altera as informações de um usuário no banco de dados
func (repository Users) EditUser(ID uint64, user models.User) error {

	// query
	statement, err := repository.db.Prepare(
		"UPDATE users SET name = ?, email = ?, email_backup = ? WHERE id = ? ",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	// executando a query
	if _, err := statement.Exec(user.Name, user.Email, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// DisableUser marca o usuário como inativo, fazendo assim um soft delete
func (repository Users) DisableUser(ID uint64) error {
	// query
	statement, err := repository.db.Prepare(
		"UPDATE users SET active = false WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	// executando a query
	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// SearchByEmail busca um usuário por e-mail busca seu id e senha
func (repository Users) SearchByEmail(email string) (models.User, error) {
	// query
	query, err := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer query.Close()

	var user models.User

	// executando a query
	if query.Next() {
		if err = query.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
