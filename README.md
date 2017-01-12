# chalmers-chop
RSS food menu parser for restaurants near Chalmer's University. Written in Go.

**chop** `[noun]` *An individual cut or portion of meat, as mutton, lamb, veal, or pork, usually one containing a rib.*

# Usage
Chalmer's Chop exposes both a binary for outputting menus as JSON as well as a Go API. Fetches both weekly menus and more detailed daily menus.

A full list of the RSS feeds used to fetch menus can be found [here](https://github.com/kayex/chalmers-chop/blob/master/config/static.go).

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

## Standalone binary
The standalone binary comes pre-loaded with a curated list of RSS feeds, which allows it to be used without any further configuration. The fetch results can be exported as JSON and transmitted to a remote server using HTTP POST.

**Building**
```bash
$ go build github.com/kayex/chalmers-chop/cmd/chop
```

**Running**
```bash
$ ./chop
```

### Exporting the menus as JSON
By supplying the `url` command line argument, the menus are exported as JSON and transmitted to `url` via HTTP POST. A `token` parameter may optionally be provided for authentication purposes. It will be included in the `Authorization` request header.

To export menus and POST as JSON
```bash
$ ./chop -url https://api.example.com/ -token my-secret-token
```

**Request headers**
```http
Content-Type: application/json
Authorization: Token {token}
```

**Request body**
```json
{
  "restaurants": [
    {
      "name": "Kårrestaurangen",
      "area": "Johanneberg",
      "menus": [
        {
          "title": "Meny Kårrestaurangen - 2017-01-09",
          "date": "2017-01-09",
          "dishes": [
            {  
              "name": "Classic Sallad",
              "contents": "Marinerad Fetaost, olivsallad, vitlöksbröd",
              "price": 80,
              "allergens": [  
                "lactose",
                "gluten"
              ]
            }
          ]
        }
      ]
    }
  ]
}
```
The `price` and `allergens` fields are optional, and may not be included depending on the completeness of the source data.

# License
MIT
