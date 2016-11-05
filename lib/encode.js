// @flow
'use strict';
import type {Section} from './state';

function encode(node: Section[]): string {
	const lines: string[] = [];
	node.forEach(section => {
		const sectionNum = Number.isFinite(section.order) ?
			String(section.order) : '';
		// section opening
		lines.push(`[${section.name}]${sectionNum}`);
		// each key-value pair
		section.dictionary.forEach((val, key) => {
			if (typeof val === 'string') {
				lines.push(` ${key}=${val}`);
			} else {
				const linesWithSameKey = val.map(v => ` ${key}=${v}`);
				lines.push(...linesWithSameKey);
			}
		});
		// section closing
		lines.push(`[/${section.name}]`);
	});
	// trailing line
	lines.push('');

	return lines.join('\n');
}

module.exports = encode;
