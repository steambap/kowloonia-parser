import fs from "fs";
import { decode } from "../src/main";

describe("decode func", () => {
	test("decode a simple config file", () => {
		const file = fs.readFileSync("test/config.txt", "utf8");

		const result = decode(file);

		expect(result).toBeTruthy();

		const configSec = result[0];

		expect(configSec.order).toEqual(0);

		const map = configSec.dictionary;

		expect(map.get("Sound")).toEqual("1");
		expect(map.get("Battle")).toEqual("1");
		expect(map.get("Host")).toEqual("kowloonia.com");
	});

	test("decode a huge default file", () => {
		const file = fs.readFileSync("test/config.txt", "utf8");

		decode(file);
	});

	test("handles duplicated prop nicely", () => {
		const file = fs.readFileSync("test/slink.txt", "utf8");

		const result = decode(file);
		const citySec = result[0];
		const map = citySec.dictionary;

		expect(map.size).toEqual(22);
		expect(map.get("SLink")?.length).toEqual(3);
	});
});
