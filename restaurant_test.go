package chalmers_chop

import (
	"reflect"
	"testing"
)

func TestRestaurant_addMenus(t *testing.T) {
	cases := []struct {
		r   Restaurant
		m   Menu
		exp Restaurant
	}{
		{
			r: Restaurant{
				Menus: []Menu{
					Menu{
						Title: "Old Menu",
						Date:  "2017-01-02",
					},
				},
			},
			m: Menu{
				Title: "New Menu",
				Date:  "2017-01-02",
			},
			exp: Restaurant{
				Menus: []Menu{
					Menu{
						Title: "New Menu",
						Date:  "2017-01-02",
					},
				},
			},
		},
	}

	for _, c := range cases {
		c.r.addMenus(&c.m)

		if !reflect.DeepEqual(c.r, c.exp) {
			t.Errorf("Expected addMenus(%v) to result in %v, got %v", c.m, c.exp, c.r)
		}
	}
}
