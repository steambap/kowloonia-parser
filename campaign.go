package main

import (
	"log"
	"strings"

	"github.com/siongui/gojianfan"
)

type CampaignText struct {
	Name        string   `json:"name"`
	SubTitle    string   `json:"subTitle"`
	Picture     string   `json:"picture"`
	Description []string `json:"description"`
}

func (c *CampaignText) convert(section *Section) *CampaignText {
	if section.name != "CAMPAIGN" {
		log.Fatalln("Not a campaign text file")
	}

	for _, line := range section.lines {
		pairs := strings.Split(line, "=")
		k, v := pairs[0], pairs[1]
		switch k {
		case "Name":
			c.Name = gojianfan.T2S(v)
		case "SubTitle":
			c.SubTitle = gojianfan.T2S(v)
		case "Picture":
			c.Picture = gojianfan.T2S(v)
		case "Description0":
			fallthrough
		case "Description1":
			fallthrough
		case "Description2":
			fallthrough
		case "Description3":
			fallthrough
		case "Description4":
			fallthrough
		case "Description5":
			fallthrough
		case "Description6":
			fallthrough
		case "Description7":
			c.Description = append(c.Description, gojianfan.T2S(v))
		}
	}

	return c
}
