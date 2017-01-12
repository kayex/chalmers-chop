package config

type RestaurantConfig struct {
	Area    string
	MenuURL string
}

type ExportConfig struct {
	URL   string
	Token string
}

type Config struct {
	Export      ExportConfig
	Restaurants []RestaurantConfig
}
