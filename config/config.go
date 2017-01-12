package config

type AreaConfig struct {
	Area     string
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

func (c *Config) GetAllMenuURLs() []string {
	var urls []string

	for name, area := range c.AreaConfigs {
		area.Area = name
		urls = append(urls, area.MenuURLs...)
	}

	return urls
}
