import { Section, SectionBody } from './definition';

class Parser {
	lines: string[];
	line: string;
	inSection: boolean;
	constructor(input: string) {
		this.lines = input.slice(1).split(/[\r\n]+/g);
		this.line = '';
		this.inSection = false;
	}

	parse() {
		const out: Section[] = [];
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
			const sectionBody: SectionBody = node[node.length - 1].dictionary;

			return this.parseSection(sectionBody);
		}

		const sectionHeader = this.parseSecHeader();
		node.push(sectionHeader);
	}

	parseSecHeader(): Section {
		const match = /^\[([^\]]*)](\d*)?$/.exec(this.line);
		if (!match) {
			throw new SyntaxError('no header');
		}

		this.inSection = true;
		const name = match[1];
		const order = Number(match[2]);

		return {name, order, dictionary: new Map()};
	}

	parseSecClosing() {
		const match = /^\[\/([^\]]*)]$/.exec(this.line);
		if (!match) {
			throw new SyntaxError('invalid section closing');
		}
		this.inSection = false;
	}

	parseSection(sectionBody: SectionBody) {
		const eqPosition = this.line.indexOf('=');
		if (eqPosition === -1) {
			return this.parseSecClosing();
		}

		const key = this.line.slice(0, eqPosition);
		const value = this.line.slice(eqPosition + 1);

		// key 可以是重複的，比如slink
		if (sectionBody.has(key)) {
			const prevVal = <string | string[]>(sectionBody.get(key));
			sectionBody.set(key, ([] as string[]).concat(prevVal, value));
		} else {
			sectionBody.set(key, value);
		}
	}
}

export default Parser;