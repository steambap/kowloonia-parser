import fs from "fs";
import { decode } from "../src/main";
import encode from "../src/encode";

describe("encode func", () => {
	test("encode parse result", () => {
		// it is a utf8 bom file actually
		const file = fs.readFileSync("test/config.txt", "utf8");

		const result = decode(file);

		const newFileContent = encode(result);
		const newFileLines = newFileContent.split(/[\r\n]+/g);

		file.split(/[\r\n]+/g).forEach((line, i) => {
			expect(line.trim()).toEqual(newFileLines[i].trim());
		});
	});

	test("encode duplicated keys", () => {
		const file = fs.readFileSync("test/slink.txt", "utf8");

		const result = decode(file);

		const newFileContent = encode(result);
		const newFileLines = newFileContent.split(/[\r\n]+/g);

		expect(newFileLines.length).toEqual(27);
	});
});
