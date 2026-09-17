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
// 		<start>       ::= <sexpr>
// 		<sexpr>       ::= NUMBER | SYMBOL | LPAR <list> RPAR | QUOTE <sexpr>
// 		<list>        ::= <sexpr> <tail> | \epsilon
// 		<tail>        ::= <sexpr> <tail> | \epsilon | DOT <sexpr>

type ParserImpl struct {
	lex     *lexer
	peekTok *token
	// createdSExpr []*SExpr
}

// Helper function which returns the next token.
func (p *ParserImpl) nextToken() (*token, error) {
	if tok := p.peekTok; tok != nil {
		p.peekTok = nil
		return tok, nil
	}

	tok, err := p.lex.next()
	if err != nil {
		return nil, ErrParser
	}

	return tok, nil
}

// Helper function which puts a token back as the next token.
func (p *ParserImpl) backToken(tok *token) {
	p.peekTok = tok
}

// Helper function to peek the next token.
func (p *ParserImpl) peekToken() (*token, error) {
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}
	p.backToken(tok)

	return tok, nil
}

// Helper function to determine if 2 terms are equal
// func termsEqual(a, b *SExpr) bool {
// 	// if a == b {
// 	// 	return true
// 	// }
// 	if a.atom != b.atom || a.car != b.car || a.cdr != b.cdr {
// 		return false
// 	}
// 	// for i, _ := range a.cdr {
// 	// 	if a.Args[i] != b.Args[i] {
// 	// 		return false
// 	// 	}
// 	// }
// 	return true
// }

// Helper function to determine a term with the same values as newTerm has already been created
// func (p *ParserImpl) termExists(newTerm *SExpr) *SExpr {
// 	for _, term := range p.createdSExpr {
// 		if termsEqual(newTerm, term) {
// 			return term
// 		}
// 	}
// 	// add newTerm to createdTerms array
// 	p.createdSExpr = append(p.createdSExpr, newTerm)
// 	return newTerm
// }

type Parser interface {
	Parse(string) (*SExpr, error)
}

func NewParser() Parser {
	return &ParserImpl{}
}

// Parse the non-terminal <start>
// <start> ::= <sexpr>
func (parser *ParserImpl) Parse(input string) (*SExpr, error) {
	parser.lex = newLexer(input)

	// peek next token to determine case to do
	tok, err := parser.peekToken()
	if err != nil {
		return nil, ErrParser
	}

	// Testing termExists() implementation works
	/* println(parser.termExists(&Term{Typ: TermNumber, Literal: "1"}) == parser.termExists(&Term{Typ: TermNumber, Literal: "1"}))
	println(len(parser.createdTerms))
	return nil, nil */

	// FIRST(<start>) = {NUM, SYMBOL, LPAR, QUOTE}
	// FOLLOW(<start>) = {EOF}
	switch tok.typ {
	case tokenNumber, tokenSymbol, tokenLpar, tokenQuote:
		// Parse <sexpr>
		sexpr, err := parser.sexprNT()
		if err != nil {
			return nil, ErrParser
		}

		// // FOLLOW(<start>) = {EOF}
		// // check that next token is EOF, otherwise return error
		// if nextTok, err := parser.nextToken(); err != nil || nextTok.typ != tokenEOF {
		// 	return nil, ErrParser
		// }

		return sexpr, nil

	default:
		return nil, ErrParser
	}
}

// Parse the non-terminal <sexpr>
// <sexpr> ::= NUMBER | SYMBOL | LPAR <list> RPAR | QUOTE <sexpr>
func (parser *ParserImpl) sexprNT() (*SExpr, error) {
	var sexpr *SExpr = &SExpr{atom: nil, car: nil, cdr: nil}

	// peek next token to determine case to do
	tok, err := parser.peekToken()
	if err != nil {
		return nil, ErrParser
	}

	// FIRST(<sexpr>) = {NUM, SYMBOL, LPAR, QUOTE}
	// need some check for EOF
	switch tok.typ {
	case tokenNumber:
		// Consume NUM token
		tok, _ := parser.nextToken()
		sexpr.atom = mkTokenNumber(tok.num.String())
		return sexpr, nil

	case tokenSymbol:
		// Consume NUM token
		tok, _ := parser.nextToken()
		return &SExpr{atom: mkTokenSymbol(tok.literal)}, nil

	// <sexpr> ::= LPAR <list> RPAR
	case tokenLpar:
		// Consume LPAR; Checking ensured by peekToken()
		_, _ = parser.nextToken()

		// Parse <list>
		list, err := parser.listNT()
		if err != nil {
			return nil, ErrParser
		}

		// Consume RPAR; Checking ensured by listNT()
		_, _ = parser.nextToken()

		// determine sexpr format
		sexpr.car = list

		return sexpr, nil

	// <sexpr> ::= QUOTE <sexpr>
	case tokenQuote:
		// Consume QUOTE; Checking ensured by peekToken()
		_, _ = parser.nextToken()

		// Parse <sexpr>
		innerSExpr, err := parser.sexprNT()
		if err != nil {
			return nil, ErrParser
		}

		sexpr.car = innerSExpr

		return sexpr, nil

	default:
		return nil, ErrParser
	}
}

// <list> ::= <sexpr> <tail> | \epsilon
func (parser *ParserImpl) listNT() (*SExpr, error) {
	// peek next token to determine case to do
	tok, err := parser.peekToken()
	if err != nil {
		return nil, ErrParser
	}

	// FIRST(<sexpr>) = {NUM, SYMBOL, LPAR, QUOTE}
	// need some check for EOF
	switch tok.typ {
	case tokenNumber, tokenSymbol, tokenLpar, tokenQuote:
		// Parse <sexpr>
		sexpr, err := parser.sexprNT()
		if err != nil {
			return nil, ErrParser
		}

		// peek next token to determine how to format
		// {"(a b . c)", "(A . (B . C))"},
		// {"(a b c d)", "(A . (B . (C . (D . NIL))))"},
		tok, err := parser.peekToken()
		if err != nil {
			return nil, ErrParser
		}

		// Parse <tail>
		tail, err := parser.tailNT()
		if err != nil {
			return nil, ErrParser
		}

		println("Testing")

		// set sexpr.cdr to <tail> if it doesn't return nil
		if tok.typ == tokenDot || tail == nil {
			return mkConsCell(sexpr, tail), nil

		} else {
			cdrCell := mkConsCell(tail, nil)
			return mkConsCell(sexpr, cdrCell), nil
		}

	// <list> ::= \epsilon
	case tokenRpar:
		return nil, nil

	default:
		return nil, ErrParser
	}
}

// <tail> ::= <sexpr> <tail> | \epsilon | DOT <sexpr>
func (parser *ParserImpl) tailNT() (*SExpr, error) {
	// peek next token to determine case to do
	tok, err := parser.peekToken()
	if err != nil {
		return nil, ErrParser
	}

	// FIRST(<tail>) = {NUM, SYMBOL, LPAR, QUOTE, DOT}
	switch tok.typ {
	case tokenNumber, tokenSymbol, tokenLpar, tokenQuote:
		// Parse <sexpr>
		sexpr, err := parser.sexprNT()
		if err != nil {
			return nil, ErrParser
		}

		// peek next token to determine how to format
		// {"(a b . c)", "(A . (B . C))"},
		// {"(a b c d)", "(A . (B . (C . (D . NIL))))"},
		tok, err := parser.peekToken()
		if err != nil {
			return nil, ErrParser
		}

		// Parse <tail>
		tail, err := parser.tailNT()
		if err != nil {
			return nil, ErrParser
		}

		// set sexpr.cdr to <tail> if it doesn't return nil
		if tok.typ == tokenDot || tail == nil {
			return mkConsCell(sexpr, tail), nil

		} else {
			cdrCell := mkConsCell(tail, nil)
			return mkConsCell(sexpr, cdrCell), nil
		}

	case tokenDot:
		// Consume DOT; Checking ensured by peekToken()
		_, _ = parser.nextToken()

		// Parse <sexpr>
		sexpr, err := parser.sexprNT()
		if err != nil {
			return nil, ErrParser
		}

		// determine sexpr format

		return sexpr, nil

	// <tail> ::= \epsilon
	case tokenRpar:
		return nil, nil

	default:
		return nil, ErrParser
	}
}
