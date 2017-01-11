package main

import (
	"encoding/json"
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

		rest.Name = restConf.Name

		fp := gofeed.NewParser()
		dailyFeed, _ := fp.ParseURL(restConf.DailyMenuURL)
		weeklyFeed, _ := fp.ParseURL(restConf.WeeklyMenuURL)

		dailyMenu := chalmers_chop.ParseDailyFeed(dailyFeed)
		weeklyMenus := chalmers_chop.ParseWeeklyFeed(weeklyFeed)

		rest.Menus = append(rest.Menus, dailyMenu)
		rest.Menus = append(rest.Menus, weeklyMenus...)

		rs = append(rs, rest)
	}

	b, err := toJson(rs)

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("out.json", b, 0644)
}

type OutputJson struct {
	Restaurants []chalmers_chop.Restaurant `json:"restaurants"`
}

func toJson(rest []chalmers_chop.Restaurant) ([]byte, error) {
	out := OutputJson{
		Restaurants: rest,
	}

	b, err := json.Marshal(out)

	if err != nil {
		return nil, err
	}

	return b, nil
}
