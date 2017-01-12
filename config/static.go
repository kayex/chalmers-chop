package config

var restaurants map[string][]string = map[string][]string{
	"Johanneberg": {
		"http://intern.chalmerskonferens.se/view/restaurant/karrestaurangen/Veckomeny.rss",
		"http://intern.chalmerskonferens.se/view/restaurant/express/V%C3%A4nster.rss",
		"http://intern.chalmerskonferens.se/view/restaurant/j-a-pripps-pub-cafe/RSS%20Feed.rss",
		"http://intern.chalmerskonferens.se/view/restaurant/hyllan/RSS%20Feed.rss",
		"http://intern.chalmerskonferens.se/view/restaurant/linsen/RSS%20Feed.rss",
	},
	"Lindholmen": {
		"http://intern.chalmerskonferens.se/view/restaurant/l-s-kitchen/Projektor.rss",
		"http://intern.chalmerskonferens.se/view/restaurant/l-s-resto/RSS%20Feed.rss",
		"http://intern.chalmerskonferens.se/view/restaurant/kokboken/RSS%20Feed.rss",
	},
}

func staticRestaurantConfigs() []RestaurantConfig {
	var rcs []RestaurantConfig

	for area, feeds := range restaurants {
		for _, rss := range feeds {
			rc := RestaurantConfig{
				Area:    area,
				MenuURL: rss,
			}

			rcs = append(rcs, rc)
		}
	}

	return rcs
}
