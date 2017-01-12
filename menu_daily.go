package chalmers_chop

import (
	"errors"
	"github.com/mmcdole/gofeed"
	"strconv"
	"strings"
	"time"
)

// Defines the mapping between <img> sources and allergens
var allergenImages map[string]Allergen = map[string]Allergen{
	"gluten-white.png": Gluten,
	"gluten-black.png": Gluten,
	"egg-white.png":    Egg,
	"egg-black.png":    Egg,
	"lactose-white":    Lactose,
	"lactose-black":    Lactose,
	"FISH_black.png":   Fish,
	"FISH_white.png":   Fish,
}

func ParseDailyMenu(feed *gofeed.Feed) *Menu {
	var menu Menu

	menu.Title = feed.Title
	menu.Date = getMenuDate(feed)

	for _, item := range feed.Items {
		menu.AddDish(parseDish(item))
	}

	return &menu
}

func getMenuDate(feed *gofeed.Feed) string {
	/*
		We could assume that all "daily menus" should have today's date and generate it ourselves,
		however, to avoid trouble with timezones and other types of clock skew we prefer fetching
		it from the feed, if possible.
	*/
	d, err := parseMenuDate(feed)

	// Default to today's date
	if err != nil {
		return time.Now().Format("2006-01-02")
	}

	return d
}

/*
The menu date is retrieved by parsing the <guid> and fetching the last 10 runes, which
will always be a date on the format YYYY-mm-dd
*/
func parseMenuDate(feed *gofeed.Feed) (string, error) {
	if len(feed.Items) == 0 {
		return "", errors.New("Could not find menu date")
	}

	t := feed.Items[0].GUID

	return string(t[len(t)-10:]), nil
}

/*
Parses a dish from an item

The dish contents, price, and allergy information is contained in a CDATA tag inside the <description> tag. For example:

<description>
	<![CDATA[Beef, wheat bread, french fries@80 <br>  <img src=http://intern.chalmerskonferens.se/uploads/allergy/icon_white/1/gluten-white.png width=25 height=25 /> ><br><br>]]>
		 ^                               ^                                                                                 ^
		 contents                        price                                                                             allergen
 </description>

*/
func parseDish(item *gofeed.Item) Dish {
	var dish Dish

	dish.Name = item.Title

	desc := trimCDATA(item.Description)

	dish.Contents = parseContents(desc)
	dish.Price = parsePrice(desc)
	dish.Allergens = parseAllergens(desc)

	return dish
}

// The dish contents is all text content of the <description> tag leading up to the @ sign
func parseContents(desc string) string {
	return strings.Split(desc, "@")[0]
}

// The dish price is all text content of the <description> tag following the first @ sign, up until the first space sign
func parsePrice(desc string) int {
	afterAtSign := strings.Split(desc, "@")[1]
	priceText := beforeSpace(afterAtSign)

	p, err := strconv.Atoi(priceText)

	if err != nil {
		return 0
	}

	return p
}

/*
The dish allergens can be determined by checking the <description> tag for <img> tags.

For example, the following dish contains the allergen "Gluten":
	<![CDATA[Hamburger of the Day@80 <br>  <img src=http://intern.chalmerskonferens.se/uploads/allergy/icon_white/1/gluten-white.png width=25 height=25 /> ><br><br>]]>
*/
func parseAllergens(desc string) []Allergen {
	var al []Allergen

	for k, a := range allergenImages {
		if strings.Contains(desc, k) {
			al = append(al, a)
		}
	}

	return al
}

func trimCDATA(text string) string {
	if !strings.HasPrefix(text, "<![CDATA") || !strings.HasSuffix(text, "]]>") {
		return text
	}

	return strings.TrimSuffix(strings.TrimPrefix(text, "<![CDATA"), "]]>")
}

func beforeSpace(str string) string {
	indexFunc := func(r rune) bool {
		return r == ' '
	}

	value := strings.IndexFunc(str, indexFunc)

	if value >= 0 && value <= len(str) {
		return str[:value]
	}

	return str
}
