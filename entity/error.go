package entity

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("not found")

// ErrEmptyName invalid empty name
var ErrEmptyName = errors.New("invalid empty name")

// ErrEmptyOwner empty owner
var ErrEmptyOwner = errors.New("invalid empty owner")

// ErrEmptyIngredientsList empty ingredients
var ErrEmptyIngredientsList = errors.New("Empty ingredients list")
