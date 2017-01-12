package config

type RestaurantConfig struct {
	MenuURLs []string `toml:"rss"`
}

type ExportConfig struct {
	URL   string `toml:"url"`
	Token string
}

type Config struct {
	ExportConfig ExportConfig          `toml:"export"`
	AreaConfigs  map[string]AreaConfig `toml:"restaurants"`
}
