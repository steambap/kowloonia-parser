package main

import (
	"log"
	"strconv"
	"strings"
)

type GrandSetup struct {
	StartYear  int64 `json:"startYear"`
	StartMonth int64 `json:"startMonth"`
	// militia buff level
	UnitBuff0 int64 `json:"unitBuff0"`
	// infantry buff level
	UnitBuff1 int64 `json:"unitBuff1"`
	// cav
	UnitBuff2 int64 `json:"unitBuff2"`
	// artillery
	UnitBuff3 int64 `json:"unitBuff3"`
	// mech
	UnitBuff4 int64 `json:"unitBuff4"`
	// buff limit
	MAXUnitBuff0 int64 `json:"maxUnitBuff0"`
	MAXUnitBuff1 int64 `json:"maxUnitBuff1"`
	MAXUnitBuff2 int64 `json:"maxUnitBuff2"`
	MAXUnitBuff3 int64 `json:"maxUnitBuff3"`
	MAXUnitBuff4 int64 `json:"maxUnitBuff4"`
	// foreign aid
	UnitAid0          int64    `json:"unitAid0"`
	UnitAid1          int64    `json:"unitAid1"`
	UnitAid2          int64    `json:"unitAid2"`
	UnitAid3          int64    `json:"unitAid3"`
	UnitAid4          int64    `json:"unitAid4"`
	YearPrefix        string   `json:"yearPrefix"`
	TP_Limit          int64    `json:"tpLimit"`
	SP_Limit          int64    `json:"spLimit"`
	REFORM_Limit      int64    `json:"reformLimit`
	UNITLVCOVER_Limit int64    `json:"unitCoverLimit"`
	MaxAP             int64    `json:"maxAP"`
	NavyLength        int64    `json:"navyLength"`
	RandomName1       []string `json:"randomName1"`
	RandomName2       []string `json:"randomName2"`
	RandomName3       []string `json:"randomName3"`
}

func (g *GrandSetup) convert(section *Section) *GrandSetup {
	if section.name != "GRAND" {
		log.Fatalln("Not a setup file grand section", section.name)
	}

	for _, line := range section.lines {
		pairs := strings.Split(line, "=")
		k, v := pairs[0], pairs[1]
		switch k {
		case "StartYear":
			g.StartYear, _ = strconv.ParseInt(v, 10, 64)
		case "StartMonth":
			g.StartMonth, _ = strconv.ParseInt(v, 10, 64)
		case "UnitBuff0":
			g.UnitBuff0, _ = strconv.ParseInt(v, 10, 64)
		case "UnitBuff1":
			g.UnitBuff1, _ = strconv.ParseInt(v, 10, 64)
		case "UnitBuff2":
			g.UnitBuff2, _ = strconv.ParseInt(v, 10, 64)
		case "UnitBuff3":
			g.UnitBuff3, _ = strconv.ParseInt(v, 10, 64)
		case "UnitBuff4":
			g.UnitBuff4, _ = strconv.ParseInt(v, 10, 64)
		case "MAXUnitBuff0":
			g.MAXUnitBuff0, _ = strconv.ParseInt(v, 10, 64)
		case "MAXUnitBuff1":
			g.MAXUnitBuff1, _ = strconv.ParseInt(v, 10, 64)
		case "MAXUnitBuff2":
			g.MAXUnitBuff2, _ = strconv.ParseInt(v, 10, 64)
		case "MAXUnitBuff3":
			g.MAXUnitBuff3, _ = strconv.ParseInt(v, 10, 64)
		case "MAXUnitBuff4":
			g.MAXUnitBuff4, _ = strconv.ParseInt(v, 10, 64)
		case "UnitAid0":
			g.UnitAid0, _ = strconv.ParseInt(v, 10, 64)
		case "UnitAid1":
			g.UnitAid1, _ = strconv.ParseInt(v, 10, 64)
		case "UnitAid2":
			g.UnitAid2, _ = strconv.ParseInt(v, 10, 64)
		case "UnitAid3":
			g.UnitAid3, _ = strconv.ParseInt(v, 10, 64)
		case "UnitAid4":
			g.UnitAid4, _ = strconv.ParseInt(v, 10, 64)
		case "YearPrefix":
			g.YearPrefix = v
		case "TP_Limit":
			g.TP_Limit, _ = strconv.ParseInt(v, 10, 64)
		case "SP_Limit":
			g.SP_Limit, _ = strconv.ParseInt(v, 10, 64)
		case "REFORM_Limit":
			g.REFORM_Limit, _ = strconv.ParseInt(v, 10, 64)
		case "UNITLVCOVER_Limit":
			g.UNITLVCOVER_Limit, _ = strconv.ParseInt(v, 10, 64)
		case "MaxAP":
			g.MaxAP, _ = strconv.ParseInt(v, 10, 64)
		case "NavyLength":
			g.NavyLength, _ = strconv.ParseInt(v, 10, 64)
		case "RandomName1":
			if strings.TrimSpace(v) != "" {
				g.RandomName1 = append(g.RandomName1, v)
			}
		case "RandomName2":
			if strings.TrimSpace(v) != "" {
				g.RandomName2 = append(g.RandomName2, v)
			}
		case "RandomName3":
			if strings.TrimSpace(v) != "" {
				g.RandomName3 = append(g.RandomName3, v)
			}
		}
	}

	return g
}
