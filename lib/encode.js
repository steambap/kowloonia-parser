// @flow
'use strict';
import type {Section} from './state.js';

function encode(node: Section[]): string {
	const lines: string[] = [];
	node.forEach(section => {
		const sectionNum = Number.isFinite(section.order)
			? String(section.order) : '';
		// section opening
		lines.push(`[${section.name}]${sectionNum}`);
		// each key-value pair
		section.dictionary.forEach((val, key) => {
			lines.push(` ${key}=${val}`);
		});
		// section closing
		lines.push(`[/${section.name}]`);
	});
	// trailing line
	lines.push('');

	return lines.join('\n');
}

module.exports = encode;
