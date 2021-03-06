package repository

import (
	"database/sql"
	"github.com/lhonda/clean-architecture-go-version/entity"
	"time"
)

//PizzaMySQL mysql repo
type PizzaMySQL struct {
	db *sql.DB
}

//NewPizzaMySQL create new repository
func NewPizzaMySQL(db *sql.DB) *PizzaMySQL {
	return &PizzaMySQL{
		db: db,
	}
}

//Create a pizza
func (r *PizzaMySQL) Create(e *entity.Pizza) (*entity.Pizza, error) {
	stmt, err := r.db.Prepare(`insert into pizza (id, name ,ingredients, created_at)
		values(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Name,
		e.GetIngredientsAsString(),
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return e, nil
}

//Get a pizza
func (r *PizzaMySQL) Get(id entity.ID) (*entity.Pizza, error) {
	stmt, err := r.db.Prepare(`select id,name, ingredients , created_at from pizza where id = ?`)
	if err != nil {
		return nil, err
	}
	var b entity.Pizza
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&b.ID, &b.Name, &b.Ingredients, &b.CreatedAt)
	}
	return &b, nil
}

// GetByName Get a pizza by name
func (r *PizzaMySQL) GetByName(name string) (*entity.Pizza, error) {
	stmt, err := r.db.Prepare(`select id,name, ingredients ,created_at from pizza where name = ?`)
	if err != nil {
		return nil, err
	}
	var b entity.Pizza
	rows, err := stmt.Query(name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&b.ID, &b.Name, &b.Ingredients, &b.CreatedAt)
	}
	return &b, nil
}

//List pizzas
func (r *PizzaMySQL) List() ([]*entity.Pizza, error) {
	stmt, err := r.db.Prepare(`select id, name, ingredients,created_at from pizza`)
	if err != nil {
		return nil, err
	}
	var pizzas []*entity.Pizza
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p entity.Pizza
		var ingredients *string
		var ts = new(string)
		err = rows.Scan(&p.ID, &p.Name, &ingredients, ts)
		if err != nil {
			return nil, err
		}

		if *ts != "" {
			t, _ := time.Parse("2006-01-02T15:04:05Z", *ts)
			p.CreatedAt = t
		}
		p.Ingredients = []entity.Ingredient{entity.Ingredient(*ingredients)}
		pizzas = append(pizzas, &p)
	}
	return pizzas, nil
}

//Delete a pizza
func (r *PizzaMySQL) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from pizza where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
