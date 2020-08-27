import fs from "fs";
import engineConfig from "../src/engine-config";
import { decode } from "../src/main";

describe("engine config", () => {
	test("decode config", () => {
		const file = fs.readFileSync("test/config.txt", "utf8");
		const result = decode(file);
		const configSec = result[0];
		const conf = engineConfig.decode(configSec);

		expect(conf.drama).toBeTruthy();
		expect(conf.campaign).toEqual(0);
	});
});
