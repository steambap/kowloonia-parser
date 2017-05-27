'use strict';
const Parse = require('./state.js');

/** @param {string} input */
module.exports.decode = function (input) {
	return new Parse(input).parse();
};
/** @param {string} input */
module.exports.tokenizer = function (input) {
	return new Parse(input);
};
