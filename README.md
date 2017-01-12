# chalmers-chop
RSS food menu parser for restaurants near Chalmer's University. Written in Go.

**chop** `[noun]` *An individual cut or portion of meat, as mutton, lamb, veal, or pork, usually one containing a rib.*

# Usage
Chalmer's Chop exposes both a binary for outputting menus as JSON as well as a Go API.

A full list of restaurant RSS feeds can be found [here](http://chalmerskonferens.se/en/rss-2/).

## Go API
```go
import (
	"github.com/kayex/chalmers-chop"
	"fmt"
)

rss := "http://intern.chalmerskonferens.se/view/restaurant/karrestaurangen/Veckomeny.rss"
restaurant := chalmers_chop.FetchFromRSS(rss)

fmt.Println(restaurant.Name)

for _, d := range restaurant.TodaysMenu().Dishes {
	fmt.Printf("%v (%v) - %v %v\n", d.Name, d.Contents, d.Price, d.Allergens)
}
```

### Types
```go
type Allergen string

const (
	Gluten  Allergen = "gluten"
	Egg              = "egg"
	Lactose          = "lactose"
	Fish             = "fish"
)

type Dish struct {
	Name      string
	Contents  string
	Price     int
	Allergens []Allergen
}

type Menu struct {
	Title  string
	Date   string
	Dishes []Dish
}

type Restaurant struct {
	Name  string
	Menus []Menu
}

```

# License
MIT
