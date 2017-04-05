package chalmers_chop

import "time"

type Restaurant struct {
	Name  string `json:"name"`
	Area  string `json:"area,omitempty"`
	Menus []Menu `json:"menus"`
}

type Menu struct {
	Title  string `json:"title"`
	Date   string `json:"date"`
	Dishes []Dish `json:"dishes"`
}

func (m *Menu) AddDish(dish Dish) {
	m.Dishes = append(m.Dishes, dish)
}

// addMenus adds one or more menus to a restaurant.
//
// Adding a menu with the same date as an existing menu will replace the
// existing menu with the new one.
func (r *Restaurant) addMenus(menus ...*Menu) {
Adding:
	for _, menu := range menus {
		for i, existing := range r.Menus {
			if menu.Date == existing.Date {
				r.Menus[i] = *menu
				continue Adding
			}
		}

		r.Menus = append(r.Menus, *menu)
	}
}

func (r *Restaurant) TodaysMenu() *Menu {
	found, menu := r.findMenuByTime(time.Now())

	if !found {
		return nil
	}

	return menu
}

func (r *Restaurant) findMenuByTime(t time.Time) (bool, *Menu) {
	date := t.Format("2006-01-02")

	for _, m := range r.Menus {
		if m.Date == date {
			return true, &m
		}
	}

	return false, nil
}
