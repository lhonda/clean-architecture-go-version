package entity

import "errors"

//NotFoundError not found
var NotFoundError = errors.New("Not found")

// EmptyNomeError invality empty nome
var EmptyNomeError = errors.New("Invalid empty nome")

// EmptyNomeError
var EmptyOwnerError = errors.New("Invalid empty owner")

// EmptyIngredientsListError
var EmptyIngredientsListError = errors.New("Empty Ingredients list")
