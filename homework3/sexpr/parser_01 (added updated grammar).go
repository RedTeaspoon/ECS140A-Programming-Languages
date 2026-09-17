package sexpr

import "errors"

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

// Old grammar:
// 		<sexpr>       ::= <atom> | <pars> | QUOTE <sexpr>
//		<atom>        ::= NUMBER | SYMBOL
// 		<pars>        ::= LPAR <dotted_list> RPAR | LPAR <proper_list> RPAR
// 		<dotted_list> ::= <proper_list> <sexpr> DOT <sexpr>
// 		<proper_list> ::= <sexpr> <proper_list> | \epsilon
// Updated grammar:
// 		<sexpr>       ::= NUMBER | SYMBOL | LPAR <list> RPAR | QUOTE <sexpr>
// 		<list>        ::= <sexpr> <tail> | \epsilon
// 		<tail>        ::= <list> | DOT <sexpr>

type Parser interface {
	Parse(string) (*SExpr, error)
}

func NewParser() Parser {
	panic("TODO: implement NewParser")
}
