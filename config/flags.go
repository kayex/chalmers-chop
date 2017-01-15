package config

import "flag"

/*
	Gets config from command line arguments
*/
func FromFlags() *Config {
	var expConf ExportConfig

	flag.StringVar(&expConf.URL, "url", "", "Export target")
	flag.StringVar(&expConf.Token, "token", "", "Export authentication token")
	flag.Parse()

	return &Config{
		Export:      expConf,
		Restaurants: staticRestaurantConfigs(),
	}
}
