package lexer

import "fmt"

// Custom error type
type LexError struct {
	position LexPosition
	message  string
}

func (e *LexError) Error() string {
	return fmt.Sprintf("[at position %d] %s", e.position.Index, e.message)
}
