package entity

//Ingredient data
type Ingredient struct {
	Name      string
}

//NewIngredient create a new ingredient
func NewIngredient(name string) (*Ingredient, error) {
	if name == "" {
		return nil, EmptyNameError
	}

	i := &Ingredient{
		Name:      name,
	}
	return i, nil
}
