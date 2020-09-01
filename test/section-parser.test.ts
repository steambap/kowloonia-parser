import fs from "fs";
import { parseSection } from "../src/section-parser";

test("parse simple config", () => {
	const file = fs.readFileSync("test/config.txt", "utf8");
	const result = parseSection(file);
	expect(result).toBeTruthy();

	const confSection = result[0];
	expect(confSection.name).toEqual("CONFIG");
	expect(confSection.order).toEqual(0);
	expect(confSection.lines.length).toEqual(12);
});

test("parse a huge default file", () => {
	const file = fs.readFileSync("test/Default.txt", "utf8");

	const result = parseSection(file);
	expect(result[0].name).toEqual("BASIC");
});

test("parse drama file", () => {
	const file = fs.readFileSync("test/Drama.txt", "utf8");

	const result = parseSection(file);
	result.forEach((section) => {
		expect(section.name).toEqual("SCRIPTEVENT");
	});
});
