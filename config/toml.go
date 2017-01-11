package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)


func FromToml(filename string) Config {
	var conf Config

	tomlData := readFile(filename)
	tomlDataString := string(tomlData[:])

	if _, err := toml.Decode(tomlDataString, &conf); err != nil {
		panic(err)
	}

	return conf
}

func readFile(filename string) []byte {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return file
}
