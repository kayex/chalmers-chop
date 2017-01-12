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
	defer timeTrack(time.Now())

	conf := config.FromToml("config.toml")
	rssURLs := conf.RestaurantConfig.MenuURLs

	rs := make(chan *chalmers_chop.Restaurant, len(rssURLs))
	var wg sync.WaitGroup

	wg.Add(len(rssURLs))

	for _, rss := range conf.RestaurantConfig.MenuURLs {
		go func(rss string) {
			defer wg.Done()
			rs <- chalmers_chop.FetchFromRSS(rss)

		}(rss)
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
	export(json, conf.ExportConfig)
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
	exporter := chalmers_chop.NewPOSTExporter(conf.URL, conf.Token)
	err := exporter.Export(json)

	if err != nil {
		panic(err)
	}
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}
