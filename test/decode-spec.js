'use strict';
const assert = require('assert');
const fs = require('fs');
const decode = require('../lib/decode.js').decode;

describe('decode func', () => {
	it('decode a simple config file', () => {
		const file = fs.readFileSync('test/config.txt', 'utf8');

		const result = decode(file);

		assert(result);
	});
});
