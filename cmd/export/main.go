package main

import (
	"encoding/json"
	"fmt"
	"github.com/kayex/chalmers-chop"
	"github.com/kayex/chalmers-chop/config"
	"github.com/mmcdole/gofeed"
	"io/ioutil"
)

func main() {
	var rs []chalmers_chop.Restaurant

	conf := config.FromToml("config.toml")

	for _, restConf := range conf.Restaurants {
		var rest chalmers_chop.Restaurant
		fmt.Printf("Menus for %v\n", restConf.Name)

		fp := gofeed.NewParser()
		dailyFeed, _ := fp.ParseURL(restConf.DailyMenuURL)
		weeklyFeed, _ := fp.ParseURL(restConf.WeeklyMenuURL)

		dailyMenu := chalmers_chop.ParseDailyFeed(dailyFeed)
		weeklyMenus := chalmers_chop.ParseWeeklyFeed(weeklyFeed)

		rest.Menus = append(rest.Menus, dailyMenu)
		rest.Menus = append(rest.Menus, weeklyMenus...)

		rs = append(rs, rest)
	}

	b, err := json.Marshal(rs[0])

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("out.json", b, 0644)
}
