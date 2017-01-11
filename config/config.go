package config

type RestaurantConfig struct {
	Name string
	DailyMenuURL string `toml:"daily_menu"`
	WeeklyMenuURL string `toml:"weekly_menu"`
}

type Config struct {
	Restaurants map[string]RestaurantConfig
}
