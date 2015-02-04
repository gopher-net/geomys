/*
	Geomys Config Database
*/

package confdb

import (
	"encoding/json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Config struct {
	Ints    []string
	ReDist  []string
	Network []string
}

func (c Config) Start(config []byte, configchan chan Config) {
	//Convert JSON config into a map.
	var conf Config
	err := json.Unmarshal(config, &conf)
	check(err)
}
