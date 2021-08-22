package entity

import "errors"

//NotFoundError not found
var NotFoundError = errors.New("not found")

// EmptyNameError invalid empty name
var EmptyNameError = errors.New("invalid empty name")

// EmptyOwnerError empty owner
var EmptyOwnerError = errors.New("invalid empty owner")

// EmptyIngredientsListError empty ingredients
var EmptyIngredientsListError = errors.New("empty Ingredients list")
