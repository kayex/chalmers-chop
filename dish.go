package chalmers_chop

type Allergen string

const (
	Gluten  Allergen = "gluten"
	Egg              = "egg"
	Lactose          = "lactose"
	Fish             = "fish"
)

type Dish struct {
	Name      string     `json:"name"`
	Contents  string     `json:"contents"`
	Price     int        `json:"price,omitempty"`
	Allergens []Allergen `json:"allergens"`
}

func NewDish(name, contents string) *Dish {
	return &Dish{
		Name:     name,
		Contents: contents,
	}
}
