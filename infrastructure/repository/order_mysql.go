package repository

import (
	"context"
	"database/sql"
	"github.com/lhonda/clean-architecture-go-version/entity"
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

	_ , err = tx.ExecContext(ctx, `insert into orders (id, owner, created_at) values(?,?,?)`,
		e.ID,
		e.Owner,
		e.CreatedAt,
	)

	// insert Pizza IDs into relation table
	if err != nil {
		return nil, err
	}

	for _, p := range e.Pizzas {
		_, err = tx.ExecContext(ctx, `insert into pizza (id, name, ingredients, order_id, created_at) values(?,?,?,?,?)`,
			p.ID,
			p.Name,
			p.Ingredients,
			e.ID,
			p.CreatedAt,
		)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return e, nil
}

//Get an order
func (r *OrderMySQL) Get(id entity.ID) (*entity.Order, error) {
	stmt, err := r.db.Prepare(`select id,owner , created_at from orders where id = ?`)
	if err != nil {
		return nil, err
	}
	var b entity.Order
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&b.ID, &b.Owner, &b.CreatedAt)
	}

	stmt, err = r.db.Prepare(`select id,name, ingredients, order_id, created_at from pizza where order_id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err = stmt.Query(id)
	if err != nil {
		return nil, err
	}

	var pizzas []entity.Pizza
	for rows.Next() {
		p := entity.Pizza{}
		err = rows.Scan(&p.ID, &p.Name, &p.Ingredients, &p.CreatedAt)
		pizzas = append(pizzas, p)
	}

	b.Pizzas = pizzas
	return &b, nil
}

//List orders
func (r *OrderMySQL) List() ([]*entity.Order, error) {
	stmt, err := r.db.Prepare(`select id, owner, created_at from orders`)
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
		err = rows.Scan(&b.ID, &b.Owner, &b.CreatedAt)
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
	// TODO: setup delete cascade on init.sql
	_, err = r.db.Exec("delete from pizza where order_id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
