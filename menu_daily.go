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
	d, err := parseMenuDate(feed)

	if err != nil {
		// Default to today's date
		return time.Now().Format("2006-01-02")
	}

	return d
}


// parseMenuDate parses the date of a menu.
//
// The menu date is retrieved by parsing the contents of the <guid> tag and fetching
// the last 10 runes, which will always be a date on the format YYYY-mm-dd
func parseMenuDate(feed *gofeed.Feed) (string, error) {
	if len(feed.Items) == 0 {
		return "", errors.New("Could not find menu date")
	}

	t := feed.Items[0].GUID

	return string(t[len(t)-10:]), nil
}

// parseDish parses a dish from a feed item.
//
// The dish contents, price, and allergy information is contained in a CDATA tag inside the <description> tag. For example:
//
// <description>
// <![CDATA[Beef, wheat bread, french fries@80 <br>  <img src=http://intern.chalmerskonferens.se/uploads/allergy/icon_white/1/gluten-white.png width=25 height=25 /> ><br><br>]]>
//          ^                               ^                                                                                 ^
//          contents                        price                                                                             allergen (Gluten)
//</description>
func parseDish(item *gofeed.Item) Dish {
	var dish Dish

	dish.Name = item.Title

	desc := trimCDATA(item.Description)

	dish.Contents = parseContents(desc)
	dish.Price = parsePrice(desc)
	dish.Allergens = parseAllergens(desc)

	return dish
}

// parseContents parses a dish's contents from its description text.
//
// The dish contents is all text content leading up to the @ sign.
func parseContents(desc string) string {
	return strings.Split(desc, "@")[0]
}

// parsePrice parses the price of a dish from its description text.
//
// The dish price is all text content following the first @ sign leading up to
// the first space sign.
func parsePrice(desc string) int {
	afterAtSign := strings.Split(desc, "@")[1]
	priceText := beforeSpace(afterAtSign)

	p, err := strconv.Atoi(priceText)

	if err != nil {
		return 0
	}

	return p
}

// parseAllergens parses the allergens of a dish from its description text.
//
// The dish allergens can be determined by searching the description text for
// occurrences of image sources defined in allergenImages.
//
// For example, the following dish description indicates that the dish contains the Gluten allergen ("gluten-white.png"):
//
// <![CDATA[Hamburger of the Day@80 <br>  <img src=http://intern.chalmerskonferens.se/uploads/allergy/icon_white/1/gluten-white.png width=25 height=25 /> ><br><br>]]>
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

// beforeSpace returns a substring of str, starting at the beginning of str
// and ending right before the first occurrence of a space character (0x20).
//
// Returns str unmodified if str does not contain any spaces.
func beforeSpace(str string) string {
	indexFunc := func(r rune) bool {
		return r == ' '
	}

	i := strings.IndexFunc(str, indexFunc)

	if i >= 0 && i <= len(str) {
		return str[:i]
	}

	return str
}
