'use strict';

export default class Parser {
	readonly lines: string[];
	line: string;
	inSection: boolean;
	constructor(input: string) {
		this.lines = input.split(/[\r\n]/g);
		this.line = '';
		this.inSection = false;
	}

	parse() {
		const out = [];
		const len = this.lines.length;

		for (let i = 0; i < len; i++) {
			const line = this.lines[i];
			if (!line || line.match(/^\s*[;#]/)) {
				continue;
			}
			this.line = line.trim();

			this.parseLine(out);
		}

		return out;
	}

	parseLine(node: Section[]) {
		if (this.inSection) {
			const sectionBody = node[node.length - 1].dictionary;

			return this.parseSection(sectionBody);
		}

		const sectionHeader = this.parseSecHeader();
		node.push(sectionHeader);
	}

	parseSecHeader() {
		const match = /^\[([^\]]*)\](\d*)?$/.exec(this.line);
		if (!match) {
			throw new SyntaxError('no header');
		}

		this.inSection = true;
		const name = match[1];
		const order = Number(match[2]);

		return {name, order, dictionary: {}};
	}

	parseSecClosing() {
		const match = /^\[\/([^\]]*)\]$/.exec(this.line);
		if (!match) {
			throw new SyntaxError('invalid section closing');
		}
		this.inSection = false;
	}

	parseSection(sectionBody: Dictionary<string>) {
		const eqPosition = this.line.indexOf('=');
		if (eqPosition === -1) {
			return this.parseSecClosing();
		}

		const key = this.line.slice(0, eqPosition);
		const value = this.line.slice(eqPosition + 1);

		sectionBody[key] = value;
	}
}

interface Section {
	name: string;
	order?: number;
	dictionary: Dictionary<string>;
}

interface Dictionary<T> {
	[K: string]: T;
}
