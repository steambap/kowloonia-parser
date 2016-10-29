'use strict';
import Parser from './state';

export function decode(input: string) {
	return new Parser(input).parse();
}

export function tokenizer(input: string) {
	return new Parser(input);
}
