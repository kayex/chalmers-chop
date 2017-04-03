package chalmers_chop

type Restaurant struct {
	Name  string `json:"name"`
	Area  string `json:"area,omitempty"`
	Menus []Menu `json:"menus"`
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
