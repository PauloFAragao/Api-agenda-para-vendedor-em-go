package repository

import (
	"api/src/models"
	"database/sql"
)

// Users representa um repositório de clientes
type Sales struct {
	db *sql.DB
}

// NewUsersRepository cria um repositório de usuários
func NewSalesRepository(db *sql.DB) *Sales {
	return &Sales{db}
}

// CreateSale insere os dados de uma venda no banco de dados
func (repository Sales) CreateSale(sale models.Sales) (uint64, error) {
	// query
	statement, err := repository.db.Prepare(
		"INSERT INTO sales (seller_id, client_id, date, sale) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	//executando a query
	result, err := statement.Exec(
		sale.SellerID,
		sale.ClientID,
		sale.Date,
		sale.Sale,
	)
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

// SearchAllSales busca todas as vedas do usuário
func (repository Sales) SearchAllSales(userID uint64) ( /*[]models.Sales*/ []models.Sale, error) {
	// query de pesquisa
	query, err := repository.db.Query(
		"SELECT date, sale FROM sales WHERE seller_id = ? AND active = true", userID)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	// var sales []models.Sales
	var sales []models.Sale

	// executando a query
	for query.Next() {
		// var sale models.Sales
		var sale models.Sale

		if err = query.Scan(
			//&sale.ID,
			//&sale.SellerID,
			//&sale.ClientID,
			&sale.Date,
			&sale.Sale,
			//&sale.Active,
		); err != nil {
			return nil, err
		}

		sales = append(sales, sale)
	}

	return sales, nil
}

// SearchByID busca uma venda especifica
func (repository Sales) SearchByID(userID, saleID uint64) ( /*models.Sales*/ models.Sale, error) {
	//query
	query, err := repository.db.Query(
		"SELECT date, sale FROM sales WHERE id = ? AND seller_id = ? AND active = true", saleID, userID)
	if err != nil {
		// return models.Sales{}, err
		return models.Sale{}, err
	}
	defer query.Close()

	// var sale models.Sales
	var sale models.Sale

	// executando a query
	if query.Next() {
		if err = query.Scan(
			//&sale.ID,
			//&sale.SellerID,
			//&sale.ClientID,
			&sale.Date,
			&sale.Sale,
			//&sale.Active,
		); err != nil {
			// return models.Sales{}, err
			return models.Sale{}, err
		}
	}

	return sale, nil
}

// SearchByClient busca todas as vendas a um cliente
func (repository Sales) SearchByClient(userID, clientID uint64) ( /*[]models.Sales*/ []models.Sale, error) {
	// query
	query, err := repository.db.Query(
		"SELECT date, sale FROM sales WHERE client_id = ? AND seller_id = ? AND active = true", clientID, userID)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	// var sales []models.Sales
	var sales []models.Sale

	// executando a query
	for query.Next() {
		// var sale models.Sales
		var sale models.Sale

		if err = query.Scan(
			//&sale.ID,
			//&sale.SellerID,
			//&sale.ClientID,
			&sale.Date,
			&sale.Sale,
			//&sale.Active,
		); err != nil {
			return nil, err
		}

		sales = append(sales, sale)
	}

	return sales, nil
}

// GetLinkedUserId é chamado para pesquisar o id do usuário que realizou uma vendas
func (repository Sales) GetLinkedUserId(saleID uint64) (uint64, error) {
	// query de pesquisa
	query, err := repository.db.Query(
		"SELECT seller_id FROM sales WHERE id = ? AND active = true", saleID)
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

// EditSale é chamado para editar uma venda
func (repository Sales) EditSale(sale models.Sales) error {
	// query
	statement, err := repository.db.Prepare(
		"UPDATE sales SET client_id = ?, date = ?, sale = ? Where id = ? ")
	if err != nil {
		return err
	}
	defer statement.Close()

	// executando a query
	if _, err := statement.Exec(sale.ClientID, sale.Date, sale.Sale, sale.ID); err != nil {
		return err
	}

	return nil
}

// DisableSale marca a interação como inativa fazendo um soft delete
func (repository Sales) DisableSale(ID uint64) error {
	// query
	statement, err := repository.db.Prepare(
		"UPDATE sales SET active = false WHERE id = ?",
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
