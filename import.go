package chalmers_chop

import "github.com/mmcdole/gofeed"

func FetchFromRSS(url string) *Restaurant {
	var rest Restaurant

	fp := gofeed.NewParser()
	dailyFeed, _ := fp.ParseURL(dailyMenuURL(url))
	weeklyFeed, _ := fp.ParseURL(weeklyMenuURL(url))

	rest.Name = parseRestaurantNameFromWeeklyFeed(weeklyFeed)

	dailyMenu := ParseDailyMenu(dailyFeed)
	weeklyMenus := ParseWeeklyMenu(weeklyFeed)

	// Let the daily menu overwrite any weekly menu entry for the same day
	rest.addMenus(weeklyMenus...)
	rest.addMenus(dailyMenu)

	return &rest
}

func dailyMenuURL(feedURL string) string {
	return feedURL + "?today=true"
}

func weeklyMenuURL(feedURL string) string {
	return feedURL
}
