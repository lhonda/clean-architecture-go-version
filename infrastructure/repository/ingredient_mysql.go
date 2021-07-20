package repository

import (
	"database/sql"
	"time"

	"github.com/lhonda/clean-architecture-go-version/entity"
)

//IngredientMySQL mysql repo
type IngredientMySQL struct {
	db *sql.DB
}

//NewIngredientMySQL create new repository
func NewIngredientMySQL(db *sql.DB) *IngredientMySQL {
	return &IngredientMySQL{
		db: db,
	}
}

//Create an ingredient
func (r *IngredientMySQL) Create(e *entity.Ingredient) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into ingredient (ID, name, created_at) 
		values(?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		entity.NewID(),
		e.Name,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

//Get a ingredient
func (r *IngredientMySQL) Get(id entity.ID) (*entity.Ingredient, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book where id = ?`)
	if err != nil {
		return nil, err
	}
	var b entity.Ingredient
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&b.ID, &b.Name, &b.CreatedAt)
	}
	return &b, nil
}

//Update an ingredient
func (r *IngredientMySQL) Update(e *entity.Ingredient) error {
	_, err := r.db.Exec("update ingredient set name = ? where id = ?", e.Name, e.ID)
	if err != nil {
		return err
	}
	return nil
}



//List ingredients
func (r *IngredientMySQL) List() ([]*entity.Ingredient, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book`)
	if err != nil {
		return nil, err
	}
	var ingredients []*entity.Ingredient
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b entity.Ingredient
		err = rows.Scan(&b.ID, &b.Name, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, &b)
	}
	return ingredients, nil
}

//Delete an ingredient
func (r *IngredientMySQL) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from ingredient where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
