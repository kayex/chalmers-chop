package chalmers_chop

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
	"strings"
)

func ParseWeeklyMenu(feed *gofeed.Feed) []*Menu {
	var menus []*Menu

	for _, item := range feed.Items {
		var menu Menu

		menu.Title = item.Title
		menu.Date = parseDate(item)
		menu.Dishes = parseDishes(item)

		menus = append(menus, &menu)
	}

	return menus
}

/*
	Finds the restaurant name

	The restaurant name will be the entire contents of the <title>
	tag, except for the five first characters ("Meny ")

	For example:

	<title>Meny Kårrestaurangen</title>
	            ^
	            name starts here
*/
func parseRestaurantNameFromWeeklyFeed(feed *gofeed.Feed) string {
	t := feed.Title
	return string(t[5:])
}

func parseDate(item *gofeed.Item) string {
	// Date is always the last 10 characters of the Title property in the format
	// YYYY-mm-dd
	t := item.Title
	return string(t[len(t)-10:])
}

/*
Searches the contents of the item <description> tag for dishes

The <description> tag contains a single <table> element, for example:

<table>
  <tr>
    <td>
      <b>Hamburger of the Day</b>
    </td>
    <td>
      Beef, wheat bread, french fries
    </td>
  </tr>
</table>

Each <tr> represents a single dish.
Each dish has at least two <td>.

The first <td> contains the dish name, surrounded by a <b> tag
The second <td> contains the dish contents

Dish price and allergy information is not available in the source data
*/
func parseDishes(item *gofeed.Item) []Dish {
	var dishes []Dish
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(item.Description))

	if err != nil {
		panic(err)
	}

	doc.Find("table tr").Each(func(i int, tr *goquery.Selection) {
		var dish Dish

		td := tr.Find("td")

		dish.Name = td.Eq(0).Children().First().Text()
		dish.Contents = td.Eq(1).Text()

		dishes = append(dishes, dish)
	})

	return dishes
}
