package main

import (
	"strconv"
	"strings"

	"github.com/siongui/gojianfan"
)

type BasicDefault struct {
	Turn          int64 `json:"turn"`
	CurrentPlayer int64 `json:"currentPlayer"`
	AICheat       int64 `json:"aiCheat"`
}

func readBasic(s *Section) *BasicDefault {
	basic := &BasicDefault{}
	for _, line := range s.lines {
		pairs := strings.Split(line, "=")
		k, v := pairs[0], pairs[1]
		switch k {
		case "Turn":
			basic.Turn, _ = strconv.ParseInt(v, 10, 64)
		case "CurrentPlayer":
			basic.CurrentPlayer, _ = strconv.ParseInt(v, 10, 64)
		case "AICheat":
			basic.AICheat, _ = strconv.ParseInt(v, 10, 64)
		}
	}

	return basic
}

type Faction struct {
	Name          string           `json:"name"`
	Player        string           `json:"player"`
	CountryState  string           `json:"countryState"`
	CountryType   string           `json:"countryType"`
	BGM1          string           `json:"bgm1"`
	EmpireName    string           `json:"empireName"`
	ChinaName     string           `json:"chinaName"`
	NationName    string           `json:"nationName"`
	BattleBonus   int64            `json:"battleBonus"`
	AP            int64            `json:"actionPt"`
	SP            int64            `json:"supplyPt"`
	TP            int64            `json:"tradePt"`
	BanWeapon     bool             `json:"banWeapon"`
	OffboardSP    int64            `json:"offBoardSP"`
	OffboardIP    int64            `json:"offBoardIP"`
	OffboardTP    int64            `json:"offBoardTP"`
	Faction       string           `json:"power"`
	Flag          int64            `json:"flag"`
	WarTry        int64            `json:"warTry"`
	MoveTry       int64            `json:"moveTry"`
	StrikeTry     int64            `json:"strikeTry"`
	TrainTry      int64            `json:"trainTry"`
	ReformTry     int64            `json:"reformTry"`
	DipTry        int64            `json:"dipTry"`
	SPTry         int64            `json:"spTry"`
	Culture       int64            `json:"culture"`
	Neutral       int64            `json:"neutral"`
	Rubbish       int64            `json:"rubbish"`
	TopEquipment  string           `json:"topEquipment"`
	Government    string           `json:"government"`
	Economy       string           `json:"economy"`
	Value         string           `json:"value"`
	Strategy      string           `json:"strategy"`
	Target        string           `json:"target"`
	Master        string           `json:"master"`
	Commander     string           `json:"commander"`
	Conscription  int64            `json:"conscription0"`
	Conscription1 int64            `json:"conscription1"`
	Conscription2 int64            `json:"conscription2"`
	Conscription3 int64            `json:"conscription3"`
	Conscription4 int64            `json:"conscription4"`
	Reform        string           `json:"reform"`
	ReformPoint   int64            `json:"reformPoint"`
	Agenda0       string           `json:"agenda0"`
	Agenda1       string           `json:"agenda1"`
	Agenda2       string           `json:"agenda2"`
	Agenda3       string           `json:"agenda3"`
	Agenda4       string           `json:"agenda4"`
	Description0  string           `json:"description0"`
	Description1  string           `json:"description1"`
	Description2  string           `json:"description2"`
	Description3  string           `json:"description3"`
	HeadQuarter   string           `json:"headQuarter"`
	SortList      []int64          `json:"unitList"`
	Diplomacy     map[string]int64 `json:"diplomacy"`
	Friendship    map[string]int64 `json:"friendship"`
	Market        map[string]int64 `json:"market"`
}

func readFaction(s *Section) *Faction {
	faction := &Faction{
		SortList:   make([]int64, 0),
		Diplomacy:  make(map[string]int64),
		Friendship: make(map[string]int64),
		Market:     make(map[string]int64),
	}
	for _, line := range s.lines {
		pairs := strings.Split(line, "=")
		k, v := pairs[0], pairs[1]
		switch k {
		case "Name":
			faction.Name = gojianfan.T2S(v)
		case "Player":
			faction.Player = v
		case "CountryState":
			faction.CountryState = v
		case "CountryType":
			faction.CountryType = v
		case "BGM1":
			faction.BGM1 = v
		case "EmpireName":
			faction.EmpireName = gojianfan.T2S(v)
		case "ChinaName":
			faction.ChinaName = gojianfan.T2S(v)
		case "NationName":
			faction.NationName = gojianfan.T2S(v)
		case "BattleBonus":
			faction.BattleBonus, _ = strconv.ParseInt(v, 10, 64)
		case "AP":
			faction.AP, _ = strconv.ParseInt(v, 10, 64)
		case "SP":
			faction.SP, _ = strconv.ParseInt(v, 10, 64)
		case "TP":
			faction.TP, _ = strconv.ParseInt(v, 10, 64)
		case "BanWeapon":
			faction.BanWeapon = v == "1"
		case "OffboardSP":
			faction.OffboardSP, _ = strconv.ParseInt(v, 10, 64)
		case "OffboardIP":
			faction.OffboardIP, _ = strconv.ParseInt(v, 10, 64)
		case "OffboardTP":
			faction.OffboardTP, _ = strconv.ParseInt(v, 10, 64)
		case "Faction":
			faction.Faction = gojianfan.T2S(v)
		case "Flag":
			faction.Flag, _ = strconv.ParseInt(v, 10, 64)
		case "WarTry":
			faction.WarTry, _ = strconv.ParseInt(v, 10, 64)
		case "MoveTry":
			faction.MoveTry, _ = strconv.ParseInt(v, 10, 64)
		case "StrikeTry":
			faction.StrikeTry, _ = strconv.ParseInt(v, 10, 64)
		case "TrainTry":
			faction.TrainTry, _ = strconv.ParseInt(v, 10, 64)
		case "ReformTry":
			faction.ReformTry, _ = strconv.ParseInt(v, 10, 64)
		case "DipTry":
			faction.DipTry, _ = strconv.ParseInt(v, 10, 64)
		case "SPTry":
			faction.SPTry, _ = strconv.ParseInt(v, 10, 64)
		case "Culture":
			faction.Culture, _ = strconv.ParseInt(v, 10, 64)
		case "Neutral":
			faction.Neutral, _ = strconv.ParseInt(v, 10, 64)
		case "Rubbish":
			faction.Rubbish, _ = strconv.ParseInt(v, 10, 64)
		case "TopEquipment":
			faction.TopEquipment = v
		case "Government":
			faction.Government = v
		case "Economy":
			faction.Economy = v
		case "Value":
			faction.Value = v
		case "Strategy":
			faction.Strategy = v
		case "Target":
			faction.Target = gojianfan.T2S(v)
		case "Master":
			faction.Master = gojianfan.T2S(v)
		case "Commander":
			faction.Commander = gojianfan.T2S(v)
		case "Conscription":
			faction.Conscription, _ = strconv.ParseInt(v, 10, 64)
		case "Conscription1":
			faction.Conscription1, _ = strconv.ParseInt(v, 10, 64)
		case "Conscription2":
			faction.Conscription2, _ = strconv.ParseInt(v, 10, 64)
		case "Conscription3":
			faction.Conscription3, _ = strconv.ParseInt(v, 10, 64)
		case "Conscription4":
			faction.Conscription4, _ = strconv.ParseInt(v, 10, 64)
		case "Reform":
			faction.Reform = v
		case "ReformPoint":
			faction.ReformPoint, _ = strconv.ParseInt(v, 10, 64)
		case "Agenda0":
			faction.Agenda0 = v
		case "Agenda1":
			faction.Agenda1 = v
		case "Agenda2":
			faction.Agenda2 = v
		case "Agenda3":
			faction.Agenda3 = v
		case "Agenda4":
			faction.Agenda4 = v
		case "Description0":
			faction.Description0 = gojianfan.T2S(v)
		case "Description1":
			faction.Description1 = gojianfan.T2S(v)
		case "Description2":
			faction.Description2 = gojianfan.T2S(v)
		case "Description3":
			faction.Description3 = gojianfan.T2S(v)
		case "HeadQuarter":
			faction.HeadQuarter = gojianfan.T2S(v)
		case "SortList":
			unitIndex, _ := strconv.ParseInt(v, 10, 64)
			faction.SortList = append(faction.SortList, unitIndex)
		case "Diplomacy":
			keyValue := strings.Split(v, ",")
			facName := gojianfan.T2S(keyValue[0])
			facVal, _ := strconv.ParseInt(keyValue[1], 10, 64)
			faction.Diplomacy[facName] = facVal
		case "Friendship":
			keyValue := strings.Split(v, ",")
			facName := gojianfan.T2S(keyValue[0])
			facVal, _ := strconv.ParseInt(keyValue[1], 10, 64)
			faction.Friendship[facName] = facVal
		case "Market":
			keyValue := strings.Split(v, ",")
			facName := gojianfan.T2S(keyValue[0])
			facVal, _ := strconv.ParseInt(keyValue[1], 10, 64)
			faction.Market[facName] = facVal
		}
	}

	return faction
}

type City struct {
	Name    string `json:"name"`
	Graph   string `json:"graph"`
	PosX    int64  `json:"posX"`
	PosY    int64  `json:"posY"`
	Faction string `json:"faction"`
	Owner   string `json:"owner"`
	Index   int64  `json:"index"`
	SP      int64  `json:"supplyPt"`
	TP      int64  `json:"tradePt"`
	PP      int64  `json:"politicalPt"`
	NavyPt  int64  `json:"navyPt"`
	Police  int64  `json:"police"`
	Terrain string `json:"terrain"`
	// 0港口 1外商 2馬匹 3礦產 4石油 5都市
	Resource0   bool     `json:"resource0"`
	Resource1   bool     `json:"resource1"`
	Resource2   bool     `json:"resource2"`
	Resource3   bool     `json:"resource3"`
	Resource4   bool     `json:"resource4"`
	Resource5   bool     `json:"resource5"`
	SLink       []string `json:"sLink"`
	Improvement []string `json:"improvement"`
}

func readCity(s *Section) *City {
	city := &City{
		SLink:       make([]string, 0),
		Improvement: make([]string, 0),
	}
	for _, line := range s.lines {
		pairs := strings.Split(line, "=")
		k, v := pairs[0], pairs[1]
		switch k {
		case "Name":
			city.Name = gojianfan.T2S(v)
		case "Graph":
			city.Graph = gojianfan.T2S(v)
		case "PosX":
			city.PosX, _ = strconv.ParseInt(v, 10, 64)
		case "PosY":
			city.PosY, _ = strconv.ParseInt(v, 10, 64)
		case "Faction":
			city.Faction = gojianfan.T2S(v)
		case "Owner":
			city.Owner = gojianfan.T2S(v)
		case "Index":
			city.Index, _ = strconv.ParseInt(v, 10, 64)
		case "SP":
			city.SP, _ = strconv.ParseInt(v, 10, 64)
		case "TP":
			city.TP, _ = strconv.ParseInt(v, 10, 64)
		case "PP":
			city.PP, _ = strconv.ParseInt(v, 10, 64)
		case "NavyPt":
			city.NavyPt, _ = strconv.ParseInt(v, 10, 64)
		case "Police":
			city.Police, _ = strconv.ParseInt(v, 10, 64)
		case "Terrain":
			city.Terrain = v
		case "Resource0":
			city.Resource0 = v == "1"
		case "Resource1":
			city.Resource1 = v == "1"
		case "Resource2":
			city.Resource2 = v == "1"
		case "Resource3":
			city.Resource3 = v == "1"
		case "Resource4":
			city.Resource4 = v == "1"
		case "Resource5":
			city.Resource5 = v == "1"
		case "SLink":
			city.SLink = append(city.SLink, gojianfan.T2S(v))
		case "Improvement":
			city.Improvement = append(city.Improvement, v)
		}
	}

	return city
}

type CoreDefault struct {
	Basic   *BasicDefault `json:"basic"`
	Faction []*Faction    `json:"faction"`
	City    []*City       `json:"city"`
}

func readCore(s []*Section) *CoreDefault {
	core := &CoreDefault{
		Basic:   &BasicDefault{},
		Faction: make([]*Faction, 0),
		City:    make([]*City, 0),
	}

	for _, section := range s {
		if section.name == "BASIC" {
			core.Basic = readBasic(section)
		} else if section.name == "FACTION" {
			core.Faction = append(core.Faction, readFaction(section))
		} else if section.name == "CITY" {
			core.City = append(core.City, readCity(section))
		}
	}

	return core
}
