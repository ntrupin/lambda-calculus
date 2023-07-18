package lexer

import (
	"errors"
	"io"
	"os"

	"ntrupin.com/lambda/internal/pkg/misc"
)

// Lexer struct
type Lexer struct {
	input    string
	file     *os.File
	position LexPosition
}

// Create a new lexer pointer
func NewLexerForString(input string) *Lexer {
	return &Lexer{
		input:    input,
		position: NewLexPosition(0, 1, 1),
	}
}

// Create a new lexer pointer
func NewLexerForFile(file *os.File) *Lexer {
	if file == nil {
		return nil
	}
	return &Lexer{
		file:     file,
		position: NewLexPosition(0, 1, 1),
	}
}

// Create a new error
func (l *Lexer) error(message string) *LexError {
	return &LexError{
		position: l.position,
		message:  message,
	}
}

// Get current character
func (l *Lexer) current() (byte, error) {
	if l.file != nil {
		// If reader
		buf := make([]byte, 1)
		n, err := io.ReadFull(l.file, buf)
		if err != nil || n < 1 {
			return 0, l.error("failed to read file")
		}
		return buf[0], nil
	} else {
		// If stream
		if l.position.Index >= len(l.input) {
			return 0, nil
		}
		return l.input[l.position.Index], nil
	}
}

// Advance the lexer
func (l *Lexer) advance() error {
	cur, err := l.current()
	if err != nil {
		return err
	}

	if cur == '\n' {
		l.position.Line += 1
		l.position.Col = 1
	} else {
		l.position.Col += 1
	}
	l.position.Index += 1

	// Discard from reader
	if l.file != nil {
		l.file.Read(make([]byte, 1))
	}

	return nil
}

// Make a new lexeme
func (l *Lexer) newLexeme(vtype LexemeType, start LexPosition) (Lexeme, error) {
	if l.file != nil {
		buf := make([]byte, l.position.Index-start.Index)
		n, err := io.ReadFull(l.file, buf)
		if err != nil || n < 1 {
			return emptyLexeme(), l.error("failed to read file")
		}
		return newLexeme(vtype, string(buf), start), nil
	} else {
		return newLexeme(vtype, l.input[start.Index:l.position.Index], start), nil
	}
}

// Read the next token
func (l *Lexer) Next() (Lexeme, error) {
	// Make sure we don't have a nil pointer
	if l == nil {
		return emptyLexeme(), errors.New("no lexer provided")
	}

	start_pos := l.position
	if l.position.Index >= len(l.input) {
		return l.newLexeme(EOF, start_pos)
	}

	cur, err := l.current()
	if err != nil {
		return emptyLexeme(), err
	}

	if misc.IsIdent(cur) {
		for ok := true; ok; ok = misc.IsIdent(cur) {
			l.advance()
			cur, err = l.current()
			if err != nil {
				return emptyLexeme(), err
			}
		}
		return l.newLexeme(IDENT, start_pos)
	} else {
		// Single-character tokens
		singles := map[byte]LexemeType{
			'(': LPAREN, ')': RPAREN,
			'\\': LAMBDA,
			'.':  DOT, 0: EOF,
		}
		if vtype, ok := singles[cur]; ok {
			l.advance()
			return l.newLexeme(vtype, start_pos)
		}
	}

	return emptyLexeme(), l.error("unknown error")
}
