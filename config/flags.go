package config

import "flag"

func FromFlags() *Config {
	var expConf ExportConfig

	flag.StringVar(&expConf.URL, "url", "", "Export target")
	flag.StringVar(&expConf.Token, "token", "", "Authentication token")
	flag.Parse()

	return &Config{
		Export:      expConf,
		Restaurants: staticRestaurantConfigs(),
	}
}