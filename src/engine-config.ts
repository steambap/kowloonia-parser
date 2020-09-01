import camelCase from "lodash.camelcase";
import { EngineConfig, Section } from "./definition";
import { parseSection } from "./section-parser";

function newEngineConfig(): EngineConfig {
	return {
		sound: 1,
		battle: 1,
		drama: true,
		autoSave: true,
		stretch: false,
		campaign: 0,
		host: "localhost",
		textSpeed: 1000,
		formLeft: 0,
		formTop: 0,
		formMaximum: 0,
	};
}

function parse(section: Section): EngineConfig {
	if (section.name !== "CONFIG") {
		throw new Error("not a engine config");
	}

	const conf = newEngineConfig();
	section.lines.forEach(line => {
		const [k, v] = line.split('=');
		switch (k) {
			case "Sound":
			case "Battle":
			case "Campaign":
			case "TextSpeed":
			case "FormLeft":
			case "FormTop":
			case "FormMaximum":
				conf[camelCase(k)] = parseInt(v, 0);
				return;
			case "Drama":
			case "AutoSave":
			case "Stretch":
				conf[camelCase(k)] = v !== "0";
				return;
			case "Host":
				conf[camelCase(k)] = v;
				return;
			default:
				conf[k] = v;
				return;
		}
	})

	return conf;
}

function load(input: string) {
	const result = parseSection(input);
	return parse(result[0]);
}

export default {
	load,
	parse,
	newEngineConfig,
};
