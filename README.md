# chalmers-chop
RSS food menu parser for restaurants near Chalmer's University. Written in Go.

**chop** `[noun]` *An individual cut or portion of meat, as mutton, lamb, veal, or pork, usually one containing a rib.*

# Usage
Chalmer's Chop exposes both a binary for outputting menus as JSON as well as a Go API. Fetches both weekly menus and more detailed daily menus.

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

## Standalone binary
The standalone binary offers various ways of exporting the menu data as JSON. Currently the only supported export method is a `POST`-request, with an optional authentication header.

###
**Building**
```bash
$ go build github.com/kayex/chalmers-chop/cmd/chop
```

**Running**
```bash
$ ./chop
```

### Config
The program needs a `config.toml` file in the same directory to run. You can find an example config file [here](https://github.com/kayex/chalmers-chop/blob/master/config.toml.example).

**config.toml**
```toml
[export]
url = 'https://api.example.com'
token = 'secret-token'

[restaurants]
rss = [
    'http://intern.chalmerskonferens.se/view/restaurant/karrestaurangen/Veckomeny.rss' # One URL per restaurant
]
```

### POST-request
This is the structure of the `POST`-request being sent. The application requires the export target to properly reply with a valid `2XX` status code.

```http
Authorization: Token {token}
```

```json
{
  "restaurants": [
    {
      "name": "Kårrestaurangen",
      "menus": [
        {
          "title": "Meny Kårrestaurangen - 2017-01-09",
          "date": "2017-01-09",
          "dishes": [
            {  
              "name":"Classic Sallad",
              "contents":"Marinerad Fetaost, olivsallad, vitlöksbröd",
              "price":80,
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
