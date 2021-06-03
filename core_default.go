package main

type BasicDefault struct {
	Turn          int64 `json:"turn"`
	CurrentPlayer int64 `json:"currentPlayer"`
	AICheat       int64 `json:"aiCheat"`
}

type Faction struct {
	Name         string `json:"name"`
	Player       string `json:"player"`
	CountryState string `json:"countryState"`
}

type City struct {
	Name  string `json:"name"`
	Graph string `json:"graph"`
	PosX  int64  `json:"posX"`
	// to convert from 2d to 3d
	// map (x, y) to (x, z)
	PosY    int64  `json:"poxZ"`
	Faction string `json:"faction"`
	Owner   string `json:"owner"`
	Index   int64  `json:"index"`
	SP      int64  `json:"supplyPt"`
	TP      int64  `json:"tradePt"`
	PP      int64  `json:"PoliticalPt"`
	NavyPt  int64  `json:"navyPt"`
	Police  int64  `json:"police"`
	Terrain int64  `json:"terrain"`
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

type CoreDefault struct {
	Basic   BasicDefault `json:"basic"`
	Faction []Faction    `json:"faction"`
	City    []City       `json:"city"`
}
