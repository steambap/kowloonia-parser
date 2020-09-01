import fs from "fs";
import engineConfig from "../src/engine-config";

test("parse a config", () => {
	const file = fs.readFileSync("test/config.txt", "utf8");
	const conf = engineConfig.load(file);

	expect(conf.sound).toEqual(1);
	expect(conf.autoSave).toEqual(true);
	expect(conf.host).toEqual("kowloonia.com");
	expect(conf.formLeft).toEqual(8);
	expect(conf.formTop).toEqual(9);
});
