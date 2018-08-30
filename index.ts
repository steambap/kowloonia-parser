export * from './lib/encode';
import Parser from './lib/state';

export function decode(input: string) {
	return new Parser(input).parse();
}

export function tokenizer(input: string) {
	return new Parser(input);
}
