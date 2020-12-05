import { GameSetup } from "./definition";

function newGameSetup(): GameSetup {
	return {
		mapWidth: 2400,
		mapHeight: 1800,
		startYear: 14,
		startMonth: 0,
		applicationTitle: "民國無雙",
		yearPrefix: "民國",
		battleBGM: "BGM2.mp3",
		defaultBGM: "BGM1.mp3",
		aiBGM: "BGM2.mp3",
		tpLimit: 9999,
		spLimit: 9999,
		reformLimit: 999,
		randomName1: [],
		randomName2: [],
		randomeName3: [],
	};
}
