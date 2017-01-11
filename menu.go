package chalmers_chop

type Menu struct {
	Title  string
	Date   string
	Dishes []Dish
}

func NewMenu(title string, date string) *Menu {
	return &Menu{
		Title: title,
		Date:  date,
	}
}

func (m *Menu) AddDish(dish Dish) {
	m.Dishes = append(m.Dishes, dish)
}
