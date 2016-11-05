// @flow
'use strict';
const Parse = require('./state.js');

module.exports.decode = function (input: string) {
	return new Parse(input).parse();
};

module.exports.tokenizer = function (input: string) {
	return new Parse(input);
};
