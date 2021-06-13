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

	campaignContent, err := ioutil.ReadFile("./test/campaign.txt")
	if err != nil {
		panic(err)
	}
	// remove BOM EF BB BF
	campaignSections := ParseString(string(campaignContent[3:]))

	campaignList := make([]*CampaignText, 0)
	for _, section := range campaignSections {
		campaignText := &CampaignText{}
		campaignList = append(campaignList, campaignText.convert(section))
	}
	campaignData, _ := json.MarshalIndent(campaignList, "", "  ")
	ioutil.WriteFile("./dist/campaign.json", campaignData, 0666)

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

	defaultContent, err := ioutil.ReadFile("./test/Default.txt")
	if err != nil {
		panic(err)
	}
	defaultSection := ParseString(string(defaultContent[3:]))
	coreData, _ := json.MarshalIndent(readCore(defaultSection), "", "  ")
	ioutil.WriteFile("./dist/core.json", coreData, 0666)
}

// for debug
func printSection(s *Section) {
	fmt.Println("Section:", s.name, " (", s.order, ")")
	fmt.Println(s.lines)
}
