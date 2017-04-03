package main

import (
	"encoding/json"
	"fmt"
	"github.com/kayex/chalmers-chop"
	"github.com/kayex/chalmers-chop/config"
	"sync"
	"time"
)

func main() {
	defer trackExeTime(time.Now())

	conf := config.FromFlags()

	rs := make(chan *chalmers_chop.Restaurant, len(conf.Restaurants))
	var wg sync.WaitGroup

	wg.Add(len(conf.Restaurants))

	for _, rest := range conf.Restaurants {
		go func(rss string) {
			defer wg.Done()
			r := chalmers_chop.FetchFromRSS(rss)
			rs <- r

		}(rest.MenuURL)
	}

	var restaurants []*chalmers_chop.Restaurant

	wg.Wait()
	close(rs)

	for r := range rs {
		restaurants = append(restaurants, r)
	}

	numRest := 0
	numMenu := 0
	numDish := 0

	for _, restu := range restaurants {
		numRest++
		numMenu += len(restu.Menus)

		for _, menu := range restu.Menus {
			numMenu++
			numDish += len(menu.Dishes)
		}
	}

	fmt.Printf("Restaurants: %v\n", numRest)
	fmt.Printf("Menus: %v\n", numMenu)
	fmt.Printf("Dishes: %v\n", numDish)

	json := toJson(restaurants)
	export(json, conf.Export)
}

type OutputJson struct {
	Restaurants []*chalmers_chop.Restaurant `json:"restaurants"`
}

func toJson(rest []*chalmers_chop.Restaurant) []byte {
	out := OutputJson{
		Restaurants: rest,
	}

	b, err := json.Marshal(out)

	if err != nil {
		panic(err)
	}

	return b
}

func export(json []byte, conf config.ExportConfig) {
	if conf.URL == "" {
		return
	}

	exporter := chalmers_chop.NewPOSTExporter(conf.URL, conf.Token)
	err := exporter.Export(json)

	if err != nil {
		panic(err)
	}
}

func trackExeTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}
