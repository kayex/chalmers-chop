package chalmers_chop

import "github.com/mmcdole/gofeed"

func FetchFromRSS(rssURL string) *Restaurant {
	var rest Restaurant

	fp := gofeed.NewParser()
	dailyFeed, _ := fp.ParseURL(dailyMenu(rssURL))
	weeklyFeed, _ := fp.ParseURL(weeklyMenu(rssURL))

	rest.Name = parseRestaurantNameFromWeeklyFeed(weeklyFeed)

	dailyMenu := ParseDailyMenu(dailyFeed)
	weeklyMenus := ParseWeeklyMenu(weeklyFeed)

	// Let the daily menu overwrite any weekly menu entry for the same day
	rest.addMenus(weeklyMenus...)
	rest.addMenus(dailyMenu)

	return &rest
}

// Get the daily menu RSS URL
func dailyMenu(rssURL string) string {
	return rssURL + "?today=true"
}

// Get the weekly menu RSS URL
func weeklyMenu(rssURL string) string {
	return rssURL
}
