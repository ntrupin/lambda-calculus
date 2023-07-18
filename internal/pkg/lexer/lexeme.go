package lexer

// Pseudo-enum type
type LexemeType int8

const (
	NULL LexemeType = iota
	EOF
	LPAREN
	RPAREN
	LAMBDA
	DOT
	IDENT
)

// Lexeme struct
type Lexeme struct {
	Value    string
	Position LexPosition
	Vtype    LexemeType
}

// Create a new lexeme
func newLexeme(vtype LexemeType, value string, position LexPosition) Lexeme {
	return Lexeme{
		Value:    value,
		Position: position,
		Vtype:    vtype,
	}
}

// Default lexeme initializer
func emptyLexeme() Lexeme {
	return Lexeme{
		Value:    "",
		Position: NewLexPosition(0, 0, 0),
		Vtype:    NULL,
	}
}
