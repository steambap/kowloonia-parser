import Parser from './parser';

export function decode(txt: string) {
  return new Parser(txt).parse()
}
