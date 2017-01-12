package config

type RestaurantConfig struct {
	MenuURLs []string `toml:"rss"`
}

type Config struct {
	RestaurantConfig RestaurantConfig `toml:"restaurants"`
}
