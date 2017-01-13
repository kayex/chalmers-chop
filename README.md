# chalmers-chop
RSS food menu parser for restaurants near Chalmer's University. Written in Go.

**chop** `[noun]` *An individual cut or portion of meat, as mutton, lamb, veal, or pork, usually one containing a rib.*

# Usage
Chalmer's Chop exposes a simple Go API, as well a standalone binary for fetching the menus and outputting them as JSON.

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
The standalone binary comes pre-loaded with a curated list of RSS sources, which allows it to be used without any additional configuration. The fetch results can be exported as JSON and transmitted to a remote server using HTTP POST.

A full list over the RSS sources bundled with the standalone binary can be found [here](https://github.com/kayex/chalmers-chop/blob/master/config/static.go).

**Building**
```bash
$ go build github.com/kayex/chalmers-chop/cmd/chop
```

**Running**
```bash
$ ./chop
```

### Exporting the menus as JSON
By supplying the `url` command line argument, the menus are exported as JSON and transmitted to `url` via HTTP POST.

```bash
$ ./chop -url https://api.example.com/ -token my-secret-token
```
A `token` parameter may optionally be provided for authentication purposes.

### Request format

**Headers**
```http
Content-Type: application/json
```

If the `token` argument is provided, the following header will also be sent
```http
Authorization: Token my-secret-token
```

**Body**
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
