'use strict';
const assert = require('assert');
const fs = require('fs');
const decode = require('../lib/decode.js').decode;

const is = assert.equal;

describe('decode func', () => {
	it('decode a simple config file', () => {
		const file = fs.readFileSync('test/config.txt', 'utf8');

		const result = decode(file);

		assert(result);

		const configSec = result[0];

		assert(configSec.order === 0);

		const map = configSec.dictionary;

		assert(map.get('Sound'), '1');
		assert(map.get('Battle'), '1');
		assert(map.get('Host'), 'kowloonia.com');
	});

	it('decode a huge default file', () => {
		const file = fs.readFileSync('test/config.txt', 'utf8');

		assert(decode(file));
	});

	it('handles duplicated prop nicely', () => {
		const file = fs.readFileSync('test/slink.txt', 'utf8');

		const result = decode(file);
		const citySec = result[0];
		const map = citySec.dictionary;

		is(map.size, 22);
		is(map.get('SLink').length, 3);
	});
});
