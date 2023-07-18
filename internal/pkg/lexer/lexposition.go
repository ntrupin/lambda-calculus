package lexer

// Lexer position information
type LexPosition struct {
	Index int
	Line  int
	Col   int
}

// Create a new LexPosition
func NewLexPosition(index int, line int, col int) LexPosition {
	return LexPosition{
		Index: index,
		Line:  line,
		Col:   col,
	}
}
