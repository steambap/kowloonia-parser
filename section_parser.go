package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct {
	lines     []string
	line      string
	inSection bool
}

type Section struct {
	name  string
	order int64
	lines []string
}

func ParseString(input string) []*Section {
	parser := &Parser{
		lines: regexp.MustCompile("[\r\n]+").Split(input, -1),
	}

	return parser.Parse()
}

var emptyLineTest = regexp.MustCompile(`^\s*[;#]`)

func (p *Parser) Parse() []*Section {
	out := make([]*Section, 0)
	length := len(p.lines)
	for i := 0; i < length; i++ {
		line := p.lines[i]
		if line == "" || emptyLineTest.MatchString(line) {
			continue
		}
		p.line = strings.TrimSpace(line)

		out = p.ParseLine(out)
	}

	return out
}

func (p *Parser) ParseLine(out []*Section) []*Section {
	if p.inSection {
		header := out[len(out)-1]
		p.FindClose(header)
		return out
	}

	sectionHeader := p.ParseSecHeader()
	out = append(out, sectionHeader)
	return out
}

var sectionHeaderTest = regexp.MustCompile(`^\[([A-Z_]+)](\d+)$`)

func (p *Parser) ParseSecHeader() *Section {
	if !sectionHeaderTest.MatchString(p.line) {
		log.Fatalln("No Header:", p.line, "len", len(p.line))
	}
	section := &Section{}

	p.inSection = true
	section.lines = make([]string, 0)
	match := sectionHeaderTest.FindStringSubmatch(p.line)
	section.name = match[1]
	section.order, _ = strconv.ParseInt(match[2], 10, 64)
	return section
}

func (p *Parser) FindClose(header *Section) {
	closing := "[/" + header.name + "]"
	if p.line == closing {
		p.inSection = false
	} else {
		header.lines = append(header.lines, p.line)
	}
}
