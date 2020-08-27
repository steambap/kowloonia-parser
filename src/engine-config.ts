import { EngineConfig, Section, SectionBody } from "./definition";

function newEngineConfig(): EngineConfig {
	return {
		sound: 1,
		battle: 1,
		drama: true,
		autosave: true,
		corp: "1",
		stretch: false,
		campaign: 0,
		host: "kowloonia.com",
		textSpeed: 1000,
		formLeft: 8,
		formTop: 9,
		formMaximum: 0,
	};
}

function decode(section: Section): EngineConfig {
	if (section.name !== "CONFIG") {
		throw new Error("not a engine config");
	}

	const body = section.dictionary;
	const sound = parseInt(body.get("Sound") as string, 0);
	const battle = parseInt(body.get("Battle") as string, 0);
	const drama = body.get("Drama") !== "0";
	const autosave = body.get("AutoSave") !== "0";
	const corp = body.get("Corp") as string;
	const stretch = body.get("Stretch") !== "0";
	const campaign = parseInt(body.get("Campaign") as string, 0);
	const host = body.get("Host") as string;
	const textSpeed = parseInt(body.get("TextSpeed") as string, 0);
	const formLeft = parseInt(body.get("FormLeft") as string, 0);
	const formTop = parseInt(body.get("FormTop") as string, 0);
	const formMaximum = parseInt(body.get("FormMaximum") as string, 0);

	return {
		sound,
		battle,
		drama,
		autosave,
		corp,
		stretch,
		campaign,
		host,
		textSpeed,
		formLeft,
		formTop,
		formMaximum,
	};
}

function encode(config: EngineConfig): Section {
	const dictionary: SectionBody = new Map();
	dictionary.set("Sound", config.sound.toString());
	dictionary.set("Battle", config.battle.toString());
	dictionary.set("Drama", config.drama ? "1" : "0");
	dictionary.set("AutoSave", config.autosave ? "1" : "0");
	dictionary.set("Corp", config.corp);
	dictionary.set("Stretch", config.stretch ? "1" : "0");
	dictionary.set("Campaign", config.campaign.toString());
	dictionary.set("Host", config.host);
	dictionary.set("TextSpeed", config.textSpeed.toString());
	dictionary.set("FormLeft", config.formLeft.toString());
	dictionary.set("FormTop", config.formTop.toString());
	dictionary.set("FormMaximum", config.formMaximum.toString());

	return {
		name: "CONFIG",
		order: 0,
		dictionary,
	};
}

export default {
	decode,
	encode,
	newEngineConfig,
};
