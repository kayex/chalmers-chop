package chalmers_chop

type Restaurant struct {
	Name  string `json:"name"`
	Menus []Menu `json:"menus"`
}
