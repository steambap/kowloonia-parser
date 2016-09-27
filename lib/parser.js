'use strict';

function parse(input) {
	const lines = input.split('\n');

	const result = {};
	let curSection = '';
	lines.forEach(line => {
		line = line.trim();
		if (line === '') {
			return;
		}

		let match;
		if (curSection === '') {
			match = /\[(\w+)\](\d+)?/;
			if (!match) throw new Error('invalid section');

			curSection = match[0];
			result[curSection] = {};

			return;
		}

		if (line[0] === '[') {
			curSection = '';

			return;
		}

		const [key, value] = line.split('=');

		result[curSection][key] = value;
	});

	return result;
}

module.exports = parse;
