package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users representa um repositório de clientes
type Clients struct {
	db *sql.DB
}

// NewUsersRepository cria um repositório de usuários
func NewClientsRepository(db *sql.DB) *Clients {
	return &Clients{db}
}

// CreateUser insere um cliente no banco de dados
func (repository Clients) CreateClient(client models.Clients) (uint64, error) {

	// query
	statement, err := repository.db.Prepare(
		"INSERT INTO clients (seller_id, name, contacts, address) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//executando a query
	result, err := statement.Exec(client.SellerID, client.Name, client.Contacts, client.Address)
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

// SearchAllClients Busca todos os clientes que o usuário tem cadastrados
func (repository Clients) SearchAllClients(userID uint64) ( /*[]models.Clients*/ []models.Client, error) {
	// query de pesquisa
	query, err := repository.db.Query(
		"SELECT name, contacts, address FROM clients WHERE seller_id = ? AND active = true", userID)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	// var clients []models.Clients
	var clients []models.Client

	// executando
	for query.Next() {
		// var client models.Clients
		var client models.Client

		if err = query.Scan(
			//&client.ID,
			//&client.SellerID,
			&client.Name,
			&client.Contacts,
			&client.Address,
			//&client.Active,
		); err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

// SearchByID pesquisa um cliente do banco de dados, usando o id do cliente, o id do usuário e filtrando se o cliente está ativo
func (repository Clients) SearchByID(clientID, userID uint64) ( /*models.Clients*/ models.Client, error) {

	// query de pesquisa
	query, err := repository.db.Query(
		"SELECT name, contacts, address FROM clients WHERE id = ? AND seller_id = ? AND active = true", clientID, userID)
	if err != nil {
		// return models.Clients{}, err
		return models.Client{}, err
	}
	defer query.Close()

	// var client models.Clients
	var client models.Client

	// executando a query
	if query.Next() {
		if err = query.Scan(
			//&client.ID,
			//&client.SellerID,
			&client.Name,
			&client.Contacts,
			&client.Address,
			//&client.Active,
		); err != nil {
			// return models.Clients{}, err
			return models.Client{}, err
		}
	}

	return client, nil
}

// GetLinkedUserId pesquisa o id do usuário vinculado ao cliente
func (repository Clients) GetLinkedUserId(clientID uint64) (uint64, error) {
	// query de pesquisa
	query, err := repository.db.Query(
		"SELECT seller_id FROM clients WHERE id = ? AND active = true", clientID)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	var sellerID uint64

	// executando a query
	if query.Next() {
		if err = query.Scan(&sellerID); err != nil {
			return 0, err
		}
	}

	return sellerID, nil
}

// EditClient edita um cliente no banco de dados
func (repository Clients) EditClient(client models.Clients) error {

	// query
	statement, err := repository.db.Prepare(
		"UPDATE clients SET name = ?, contacts = ?, address = ? Where id = ? ")
	if err != nil {
		return err
	}
	defer statement.Close()

	//executando a query
	if _, err := statement.Exec(client.Name, client.Contacts, client.Address, client.ID); err != nil {
		return err
	}

	return nil
}

// DisableClient marca o cliente como inativo, fazendo assim um soft delete
func (repository Clients) DisableClient(ID uint64) error {
	// query
	statement, err := repository.db.Prepare(
		"UPDATE clients SET active = false WHERE id = ?",
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

// SearchByName busca por nome entre os clientes que o usuário tem cadastrado
func (repository Clients) SearchByName(clientName string, userID uint64) ( /*[]models.Clients*/ []models.Client, error) {

	// adicionando % no começo e no final para fazer a busca no banco de dados usando a clausula LIKE
	clientName = fmt.Sprintf("%%%s%%", clientName)

	// query de pesquisa
	query, err := repository.db.Query(
		"SELECT name, contacts, address FROM clients WHERE name LIKE ? AND seller_id = ? AND active = true", clientName, userID)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	// var clients []models.Clients
	var clients []models.Client

	// executando
	for query.Next() {
		// var client models.Clients
		var client models.Client

		if err = query.Scan(
			//&client.ID,
			//&client.SellerID,
			&client.Name,
			&client.Contacts,
			&client.Address,
			//&client.Active,
		); err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}
