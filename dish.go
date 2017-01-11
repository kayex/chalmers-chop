package chalmers_chop

type Allergen int

const (
	Gluten Allergen = iota
	Egg
	Lactose
)

type Dish struct {
	Name string
	Contents string
	Price int
	Allergens []Allergen
}

func NewDish(name, contents string) *Dish {
	return &Dish{
		Name: name,
		Contents: contents,
	}
}
