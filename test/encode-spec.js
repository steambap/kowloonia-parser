'use strict';
const assert = require('assert');
const fs = require('fs');
const decode = require('../lib/decode.js').decode;
const encode = require('../lib/encode.js');

const is = assert.equal;

describe('encode func', () => {
	it('encode parse result', () => {
		// it is a utf8 bom file actually
		const file = fs.readFileSync('test/config.txt', 'utf8').slice(1);

		const result = decode(file);

		const newFileContent = encode(result);
		const newFileLines = newFileContent.split(/[\r\n]+/g);

		file.split(/[\r\n]+/g).forEach((line, i) => {
			is(line.trim(), newFileLines[i].trim(), `line: ${i} error`);
		});
	});
});
