package config

type RestaurantConfig struct {
	MenuURLs []string `toml:"rss"`
}

type Config struct {
	RestaurantConfig `toml:"restaurants"`
}
