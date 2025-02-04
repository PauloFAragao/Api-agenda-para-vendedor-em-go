package repository

import (
	"api/src/models"
	"database/sql"
)

// Users representa um repositório de clientes
type Interactions struct {
	db *sql.DB
}

// NewInteractionsRepository cria um repositório de interações
func NewInteractionsRepository(db *sql.DB) *Interactions {
	return &Interactions{db}
}

// CreateInteractions insere no banco de dados uma interação
func (repository Interactions) CreateInteractions(interaction models.Interactions) (uint64, error) {

	// query
	statement, err := repository.db.Prepare(
		"INSERT INTO interactions (seller_id, client_id, status, date, interaction, content) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//executando a query
	result, err := statement.Exec(
		interaction.SellerID,
		interaction.ClientID,
		interaction.Status,
		interaction.Date,
		interaction.Interaction,
		interaction.Content)
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

// SearchAllInteractions busca no banco todas as interações do usuário com seus clientes
func (repository Interactions) SearchAllInteractions(userID uint64) ( /*[]models.Interactions*/ []models.Interaction, error) {
	// query de pesquisa
	query, err := repository.db.Query(
		//"SELECT * FROM interactions WHERE seller_id = ? AND active = true", userID)
		`SELECT c.name, i.status, i.date, i.interaction, i.content 
		FROM interactions i 
		INNER JOIN clients c ON c.id = i.client_id 
		WHERE i.seller_id = ? AND i.active = true`,
		userID)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	//var interactions []models.Interactions
	var interactions []models.Interaction

	// executando a query
	for query.Next() {
		//var interaction models.Interactions
		var interaction models.Interaction

		if err = query.Scan(
			//&interaction.ID,
			//&interaction.SellerID,
			//&interaction.ClientID,
			&interaction.ClientName,
			&interaction.Status,
			&interaction.Date,
			&interaction.Interaction,
			&interaction.Content,
			//&interaction.Active,
		); err != nil {
			return nil, err
		}

		interactions = append(interactions, interaction)
	}

	return interactions, nil
}

// SearchInteraction busca no banco de dados uma interação pelo seu id
func (repository Interactions) SearchByID(userID, interactionID uint64) ( /*models.Interactions*/ models.Interaction, error) {
	//query
	query, err := repository.db.Query(
		//"SELECT * FROM interactions WHERE id = ? AND seller_id = ? AND active = true", interactionID, userID)
		`SELECT c.name, i.status, i.date, i.interaction, i.content 
		FROM interactions i 
		INNER JOIN clients c ON c.id = i.client_id 
		WHERE i.id = ? AND i.seller_id = ? AND i.active = true`,
		interactionID, userID)
	if err != nil {
		// return models.Interactions{}, err
		return models.Interaction{}, err
	}
	defer query.Close()

	//var interaction models.Interactions
	var interaction models.Interaction

	// executando a query
	if query.Next() {
		if err = query.Scan(
			//&interaction.ID,
			//&interaction.SellerID,
			//&interaction.ClientID,
			&interaction.ClientName,
			&interaction.Status,
			&interaction.Date,
			&interaction.Interaction,
			&interaction.Content,
			//&interaction.Active,
		); err != nil {
			// return models.Interactions{}, err
			return models.Interaction{}, err
		}
	}

	return interaction, nil
}

// SearchByClient busca todas as interações do usuário com um cliente
func (repository Interactions) SearchByClient(userID, clientID uint64) ( /*[]models.Interactions*/ []models.Interaction, error) {
	//query
	query, err := repository.db.Query(
		//"SELECT * FROM interactions WHERE client_id = ? AND seller_id = ? AND active = true", clientID, userID)
		`SELECT c.name, i.status, i.date, i.interaction, i.content 
		FROM interactions i 
		INNER JOIN clients c ON c.id = i.client_id 
		WHERE i.client_id = ? AND i.seller_id = ? AND i.active = true`,
		clientID, userID)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	// var interactions []models.Interactions
	var interactions []models.Interaction

	// executando a query
	for query.Next() {
		// var interaction models.Interactions
		var interaction models.Interaction

		if err = query.Scan(
			//&interaction.ID,
			//&interaction.SellerID,
			//&interaction.ClientID,
			&interaction.ClientName,
			&interaction.Status,
			&interaction.Date,
			&interaction.Interaction,
			&interaction.Content,
			//&interaction.Active,
		); err != nil {
			return nil, err
		}

		interactions = append(interactions, interaction)
	}

	return interactions, nil
}

// GetLinkedUserId pesquisa o id do usuário vinculado a interação
func (repository Interactions) GetLinkedUserId(interactionID uint64) (uint64, error) {
	// query de pesquisa
	query, err := repository.db.Query(
		"SELECT seller_id FROM interactions WHERE id = ? AND active = true", interactionID)
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

// EditInteraction edita uma interação no banco de dados
func (repository Interactions) EditInteraction(interaction models.Interactions) error {
	// query
	statement, err := repository.db.Prepare(
		"UPDATE interactions SET client_id = ?, status = ?, date = ?, interaction = ?, content = ? Where id = ? ")
	if err != nil {
		return err
	}
	defer statement.Close()

	//executando a query
	if _, err := statement.Exec(interaction.ID, interaction.Status, interaction.Date, interaction.Interaction, interaction.Content, interaction.ID); err != nil {
		return err
	}

	return nil
}

// DisableInteraction marca a interação como inativa fazendo um soft delete
func (repository Interactions) DisableInteraction(ID uint64) error {
	// query
	statement, err := repository.db.Prepare(
		"UPDATE interactions SET active = false WHERE id = ?",
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

func (repository Interactions) SearchTaggedInteractions(userID uint64) ( /*[]models.Interactions*/ []models.Interaction, error) {
	// query de pesquisa
	query, err := repository.db.Query(
		//"SELECT * FROM interactions WHERE seller_id = ? AND active = true AND date >= NOW()", userID)
		`SELECT c.name, i.status, i.date, i.interaction, i.content 
		FROM interactions i 
		INNER JOIN clients c ON c.id = i.client_id 
		WHERE i.seller_id = ? AND i.active = true AND i.date >= NOW()`,
		userID)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	//var interactions []models.Interactions
	var interactions []models.Interaction

	// executando a query
	for query.Next() {
		//var interaction models.Interactions
		var interaction models.Interaction

		if err = query.Scan(
			//&interaction.ID,
			//&interaction.SellerID,
			//&interaction.ClientID,
			&interaction.ClientName,
			&interaction.Status,
			&interaction.Date,
			&interaction.Interaction,
			&interaction.Content,
			//&interaction.Active,
		); err != nil {
			return nil, err
		}

		interactions = append(interactions, interaction)
	}

	return interactions, nil
}

func (repository Interactions) SearchInteractionsMarkedOnDate(stringDate string, userID uint64) ([]models.Interaction, error) {
	// query de pesquisa
	query, err := repository.db.Query(
		`SELECT c.name, i.status, i.date, i.interaction, i.content 
		FROM interactions i 
		INNER JOIN clients c ON c.id = i.client_id 
		WHERE DATE(i.date) = ? AND i.seller_id = ? AND i.active = true`,
		stringDate, userID)

	if err != nil {
		return nil, err
	}
	defer query.Close()

	var interactions []models.Interaction

	// executando a query
	for query.Next() {

		var interaction models.Interaction

		if err = query.Scan(
			&interaction.ClientName,
			&interaction.Status,
			&interaction.Date,
			&interaction.Interaction,
			&interaction.Content,
		); err != nil {
			return nil, err
		}

		interactions = append(interactions, interaction)
	}

	return interactions, nil
}
