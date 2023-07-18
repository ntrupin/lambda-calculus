package lambda

import (
	"flag"
	"fmt"

	"ntrupin.com/lambda/internal/pkg/lexer"
	"ntrupin.com/lambda/internal/pkg/misc"
)

type VM struct {
	lexer   *lexer.Lexer
	context map[string]int
	funcs   map[string][]rune
	level   int
}

/*func (vm *VM) shiftFreeVariable(cutoff int, inc int) {
	for k, v := range vm.vars {
		if v >= cutoff {
			vm.vars[k] += inc
		}
	}
}*/

func NewVM(lexer *lexer.Lexer) *VM {
	return &VM{
		lexer:   lexer,
		context: map[string]int{},
		funcs:   map[string][]rune{},
		level:   0,
	}
}

func (vm *VM) eval() {
	for {
		token, err := vm.lexer.Next()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		fmt.Printf("%s", token.Value)
		switch token.Vtype {
		case lexer.LAMBDA:
			ident, err := vm.lexer.Next()
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			} else if ident.Vtype != lexer.IDENT {
				fmt.Printf("Wrong Vtype\n")
				return
			}

			vm.context[string(ident.Value)] = len(vm.context)
		case lexer.EOF:
			return
		}
	}
}

func Run() {
	filePtr := flag.String("file", "", "file to read")

	flag.Parse()

	// Check for unnamed arg
	if !misc.FlagExists("file") {
		if len(flag.Args()) > 0 {
			*filePtr = flag.Arg(0)
		} else {
			fmt.Println("no file provided")
			return
		}
	}

	file := misc.OpenFile(*filePtr)
	defer misc.CloseFile(file)

	NewVM(lexer.NewLexerForString("\\x.\\y.x")).eval()
}
