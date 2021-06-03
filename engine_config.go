package main

import (
	"log"
	"strconv"
	"strings"
)

type EngineConfig struct {
	Sound     bool  `json:"sound"`
	Battle    bool  `json:"battle"`
	Drama     bool  `json:"drama"`
	AutoSave  bool  `json:"autoSave"`
	Stretch   bool  `json:"stretch"`
	Campaign  int64 `json:"campaign"`
	TextSpeed int64 `json:"textSpeed"`
}

func (conf *EngineConfig) convert(section *Section) *EngineConfig {
	if section.name != "CONFIG" {
		log.Fatalln("Not a engine config")
	}

	for _, line := range section.lines {
		pairs := strings.Split(line, "=")
		k, v := pairs[0], pairs[1]
		switch k {
		case "Sound":
			conf.Sound = v == "1"
		case "Battle":
			conf.Battle = v == "1"
		case "Drama":
			conf.Drama = v == "1"
		case "AutoSave":
			conf.AutoSave = v == "1"
		case "Stretch":
			conf.Stretch = v == "1"
		case "Campaign":
			conf.Campaign, _ = strconv.ParseInt(v, 10, 64)
		case "TextSpeed":
			conf.TextSpeed, _ = strconv.ParseInt(v, 10, 64)
		}
	}

	return conf
}
