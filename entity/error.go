package entity

import "errors"

//NotFoundError not found
var NotFoundError = errors.New("Not found")

// EmptyNameError invalid empty name
var EmptyNameError = errors.New("Invalid empty name")

// EmptyOwnerError
var EmptyOwnerError = errors.New("Invalid empty owner")

// EmptyIngredientsListError
var EmptyIngredientsListError = errors.New("Empty Ingredients list")
