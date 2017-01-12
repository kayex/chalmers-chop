package chalmers_chop

import "time"

type Menu struct {
	Title  string `json:"title"`
	Date   string `json:"date"`
	Dishes []Dish `json:"dishes"`
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

func (r *Restaurant) TodaysMenu() *Menu {
	found, menu := r.findMenuByTime(time.Now())

	if !found {
		return nil
	}

	return menu
}

func (r *Restaurant) findMenuByTime(time time.Time) (bool, *Menu) {
	date := time.Format("2006-01-02")

	for _, m := range r.Menus {
		if m.Date == date {
			return true, &m
		}
	}

	return false, nil
}
