// @flow
'use strict';
const Parse = require('./state.js');

module.exports.decode = function decode(input: string) {
	return new Parse(input).parse();
};

module.exports.tokenizer = function tokenizer(input: string) {
	return new Parse(input);
};
