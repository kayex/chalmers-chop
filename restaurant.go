package chalmers_chop

type Restaurant struct {
	Name  string `json:"name"`
	Area  string `json:"area,omitempty"`
	Menus []Menu `json:"menus"`
}

/*
	Add a menu to the restaurant. Adding a menu with the Date field
	set to the same value as that of an existing menu in the slice, will
	replace the existing menu with the new one.
*/
func (r *Restaurant) addMenus(menus ...*Menu) {
	// If a menu with the same date already exists, overwrite it
	for _, menu := range menus {
		for i, m := range r.Menus {
			if m.Date == menu.Date {
				r.Menus[i] = r.Menus[len(r.Menus)-1]
				r.Menus = r.Menus[:len(r.Menus)-1]

				break
			}
		}

		r.Menus = append(r.Menus, *menu)
	}
}
