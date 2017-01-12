package main

import (
	"encoding/json"
	"github.com/kayex/chalmers-chop"
	"github.com/kayex/chalmers-chop/config"
	"io/ioutil"
)

func main() {
	var rs []*chalmers_chop.Restaurant

	conf := config.FromToml("config.toml")

	for _, rss := range conf.MenuURLs {
		rest := chalmers_chop.FetchFromRSS(rss)
		rs = append(rs, rest)
	}

	b, err := toJson(rs)

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("out.json", b, 0644)
}

type OutputJson struct {
	Restaurants []*chalmers_chop.Restaurant `json:"restaurants"`
}

func toJson(rest []*chalmers_chop.Restaurant) ([]byte, error) {
	out := OutputJson{
		Restaurants: rest,
	}

	b, err := json.Marshal(out)

	if err != nil {
		return nil, err
	}

	return b, nil
}
