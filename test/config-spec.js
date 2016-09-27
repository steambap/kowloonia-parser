'use strict';
const assert = require('assert');
const fs = require('fs');
const parse = require('../lib/parser.js');

describe('parse', () => {
	it('config', () => {
		const file = fs.readFileSync('test/config.txt', 'utf8');

		const result = parse(file);

		assert(result);
	});
});
