import { Section } from "./definition";

class Parser {
	lines: string[];
	line: string;
	inSection: boolean;
	constructor(input: string) {
		this.lines = input.slice(1).split(/[\r\n]+/g);
		this.line = "";
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

	parseLine(out: Section[]) {
		if (this.inSection) {
			const header = out[out.length - 1];

			return this.findClose(header);
		}

		const sectionHeader = this.parseSecHeader();
		out.push(sectionHeader);
	}

	parseSecHeader(): Section {
		const match = /^\[([A-Z]+)](\d+)$/.exec(this.line);
		if (!match) {
			throw new SyntaxError("no header");
		}

		this.inSection = true;
		const name = match[1];
		const order = Number(match[2]);

		return { name, order, lines: [] };
	}

	findClose(header: Section) {
		const closing = `[/${header.name}]`;
		if (this.line === closing) {
			this.inSection = false;
		} else {
			header.lines.push(this.line);
		}
	}
}

export function parseSection(input: string) {
	return new Parser(input).parse();
}

export default Parser;
