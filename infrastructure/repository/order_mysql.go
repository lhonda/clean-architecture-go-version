package repository

import (
	"database/sql"
	"github.com/lhonda/clean-architecture-go-version/entity"
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
	stmt, err := r.db.Prepare(`insert into order (id, owner, created_at) 
		values(?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Owner,
		e.CreatedAt,
	)

	// insert Pizza IDs into relation table
	stmt, err = r.db.Prepare(`insert into pizza (id, name, ingredients, order_id, created_at) values(?,?,?,?,?)`)
	if err != nil {
		return e, err
	}

	for _, p := range e.Pizzas {
		_, err = stmt.Exec(
			p.ID,
			p.Name,
			p.Ingredients,
			p.Order.ID,
			p.CreatedAt,
		)
	}

	err = stmt.Close()
	if err != nil {
		return e, err
	}
	return e, nil
}

//Get an order
func (r *OrderMySQL) Get(id entity.ID) (*entity.Order, error) {
	stmt, err := r.db.Prepare(`select id,owner , created_at from order where id = ?`)
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
	stmt, err := r.db.Prepare(`select id, owner, created_at from order`)
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
	_, err := r.db.Exec("delete from order where id = ?", id)
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
