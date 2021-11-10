package repository

import (
	"context"
	"database/sql"
	"github.com/lhonda/clean-architecture-go-version/entity"
	"strings"
	"time"
)

//OrderMySQL mysql repo
type OrderMySQL struct {
	db *sql.DB
}

//NewOrderMySQL create new repository
func NewOrderMySQL(db *sql.DB) *OrderMySQL {
	return &OrderMySQL{
		db: db,
	}
}

//Create an order
func (r *OrderMySQL) Create(e *entity.Order) (*entity.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	pizzas := make([]string, 3)
	for _, p := range e.Pizzas {
		pizzas = append(pizzas, p.Name)
	}

	_, _ = tx.Exec("SET sql_mode ='' ;")
	_, err = tx.ExecContext(ctx, `insert into orders (id, owner,pizzas, created_at) values(?,?,?,?)`,
		e.ID,
		e.Owner,
		strings.Join(pizzas, ","),
		e.CreatedAt,
	)

	// insert Pizza IDs into relation table
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return e, nil
}

//Get an order
func (r *OrderMySQL) Get(id entity.ID) (*entity.Order, error) {
	stmt, err := r.db.Prepare(`select id,owner, pizzas, created_at from orders where id = ?`)
	if err != nil {
		return nil, err
	}
	var b entity.Order
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&b.ID, &b.Owner, &b.Pizzas, &b.CreatedAt)
	}

	return &b, nil
}

//List orders
func (r *OrderMySQL) List() ([]*entity.Order, error) {
	stmt, err := r.db.Prepare(`select id, owner,pizzas, created_at from orders`)
	if err != nil {
		return nil, err
	}
	var orders []*entity.Order
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b entity.Order
		err = rows.Scan(&b.ID, &b.Owner, &b.Pizzas, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &b)
	}
	return orders, nil
}

//Delete an order
func (r *OrderMySQL) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from orders where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
