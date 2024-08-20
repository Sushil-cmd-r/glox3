package scanner

import "github.com/sushil-cmd-r/glox/token"

type Scanner struct {
	source   []byte
	rdOffset int

	ch     byte
	offset int
}

func Init(source []byte) *Scanner {
	s := &Scanner{
		source:   source,
		rdOffset: 0,

		ch:     ' ',
		offset: 0,
	}

	s.advance()
	return s
}

func (s *Scanner) Scan() (tok token.Token, lit string) {
	return
}

func (s *Scanner) advance() {
	if s.atEnd() {
		s.ch = 0
		s.offset = len(s.source)
		return
	}

	s.offset = s.rdOffset
	s.ch = s.source[s.offset]

	s.rdOffset += 1
}

func (s *Scanner) atEnd() bool {
	return s.rdOffset >= len(s.source)
}
