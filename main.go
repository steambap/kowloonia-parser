package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	confContent, err := ioutil.ReadFile("./test/config.txt")
	if err != nil {
		panic(err)
	}
	// remove BOM EF BB BF
	confSections := ParseString(string(confContent[3:]))

	engineConf := &EngineConfig{}
	confData, _ := json.MarshalIndent(engineConf.convert(confSections[0]), "", "  ")
	ioutil.WriteFile("./dist/config.json", confData, 0666)

	setupContent, err := ioutil.ReadFile("./test/Setup.txt")
	if err != nil {
		panic(err)
	}
	setupSections := ParseString(string(setupContent[3:]))

	grandSetup := &GrandSetup{}
	for _, section := range setupSections {
		if section.name == "GRAND" {
			gData, _ := json.MarshalIndent(grandSetup.convert(section), "", "  ")
			ioutil.WriteFile("./dist/grand.json", gData, 0666)
		}
	}
}

// for debug
func printSection(s *Section) {
	fmt.Println("Section:", s.name, " (", s.order, ")")
	fmt.Println(s.lines)
}
