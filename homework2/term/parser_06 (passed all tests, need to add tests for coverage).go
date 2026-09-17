package term

import (
	"errors"
	// "strconv"
)

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

// <start>		::= <term> | \epsilon
// <term>		::= ATOM <pars> | NUM | VAR
// <args>		::= <term> <otherargs>
// <pars>		::= LPAR <args> RPAR | \epsilon
// <otherargs>	::= COMMA <args> | \epsilon

type ParserImpl struct {
	lex          *lexer
	peekTok      *Token
	createdTerms []*Term
}

// Helper function which returns the next token.
func (p *ParserImpl) nextToken() (*Token, error) {
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
func (p *ParserImpl) backToken(tok *Token) {
	p.peekTok = tok
}

// Helper function to peek the next token.
func (p *ParserImpl) peekToken() (*Token, error) {
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}
	p.backToken(tok)

	return tok, nil
}

// Helper function to determine if 2 terms are equal
func termsEqual(a, b *Term) bool {
	// if a == b {
	// 	return true
	// }
	if a.Typ != b.Typ || a.Literal != b.Literal || a.Functor != b.Functor || len(a.Args) != len(b.Args) {
		return false
	}
	for i, _ := range a.Args {
		if a.Args[i] != b.Args[i] {
			return false
		}
	}
	return true
}

// Helper function to determine a term with the same values as newTerm has already been created
func (p *ParserImpl) termExists(newTerm *Term) *Term {
	for _, term := range p.createdTerms {
		if termsEqual(newTerm, term) {
			return term
		}
	}
	// add newTerm to createdTerms array
	p.createdTerms = append(p.createdTerms, newTerm)
	return newTerm
}

// Parser is the interface for the term parser.
// Do not change the definition of this interface.
type Parser interface {
	Parse(string) (*Term, error)
}

// NewParser creates a struct of a type that satisfies the Parser interface.
func NewParser() Parser {
	return &ParserImpl{}
}

// Parse the non-terminal <start>
// <start> ::= <term> | \epsilon
func (parser *ParserImpl) Parse(input string) (*Term, error) {
	var dagTree *Term
	parser.lex = newLexer(input)

	// Testing termExists() implementation works
	/* println(parser.termExists(&Term{Typ: TermNumber, Literal: "1"}) == parser.termExists(&Term{Typ: TermNumber, Literal: "1"}))
	println(len(parser.createdTerms))
	return nil, nil */

	// peek next token to determine case to do
	tok, err := parser.peekToken()
	if err != nil {
		return nil, ErrParser
	}

	// FIRST(<start>) = {ATOM, NUM, VAR, \epsilon}
	// FOLLOW(<start>) = {EOF}
	switch tok.typ {
	case tokenAtom, tokenNumber, tokenVariable:
		dagTree, err = parser.termNT()
		if err != nil {
			return nil, ErrParser
		}

		// FOLLOW(<start>) = {EOF}
		// check that next token is EOF, otherwise return error
		if nextTok, err := parser.nextToken(); err != nil || nextTok.typ != tokenEOF {
			return nil, ErrParser
		}
		return dagTree, nil

	case tokenEOF:
		return nil, nil

	default:
		return nil, ErrParser
	}
}

// Parse the non-terminal <term>
// <term> ::= ATOM <pars> | NUM | VAR
func (parser *ParserImpl) termNT() (*Term, error) {
	// peek next token to determine case to do
	tok, err := parser.peekToken()
	if err != nil {
		return nil, ErrParser
	}

	// FIRST(<term>) = {ATOM, NUM, VAR}
	switch tok.typ {
	// <term> ::= ATOM <pars>
	case tokenAtom:
		// Consume ATOM token
		tok, _ := parser.nextToken()
		atom := parser.termExists(&Term{Typ: TermAtom, Literal: tok.literal, Functor: nil, Args: nil})

		// Parse <pars>
		pars, err := parser.parsNT()
		if err != nil {
			return nil, ErrParser
		}

		// if <pars> returns nil (next token is EOF), return ATOM, otherwise make TermCompound from ATOM
		if pars == nil {
			return atom, nil
		} else {
			return parser.termExists(&Term{Typ: TermCompound, Functor: atom, Args: pars.Args}), nil
		}

	case tokenNumber:
		// Consume NUM token
		tok, _ := parser.nextToken()
		return parser.termExists(&Term{Typ: TermNumber, Literal: tok.literal}), nil

	case tokenVariable:
		// Consume VAR token
		tok, _ := parser.nextToken()
		return parser.termExists(&Term{Typ: TermVariable, Literal: tok.literal}), nil

	default:
		return nil, ErrParser

	}
}

// Parse the non-terminal <args>
// <args> ::= <term> <otherargs>
func (parser *ParserImpl) argsNT() (*Term, error) {
	var args *Term = &Term{Typ: TermAtom, Literal: ""}
	// Parse <term>
	term, err := parser.termNT()
	if err != nil {
		return nil, ErrParser
	}

	// Parse <otherargs>
	otherargs, err := parser.otherargsNT()
	if err != nil {
		return nil, ErrParser
	}

	// append <otherargs> Args to <term> Args if not nil
	if otherargs == nil {
		args.Args = []*Term{term}
		// check if otherargs is just NUM | VAR | ATOM | ATOM <pars>
	} else if otherargs.Typ != TermAtom || len(otherargs.Args) == 0 {
		args.Args = []*Term{term, otherargs}
		// otherargs being used to pass multiple args
	} else if otherargs.Typ == TermAtom && len(otherargs.Args) > 0 {
		otherargs.Args = append([]*Term{term}, otherargs.Args...)
		args.Args = otherargs.Args
	}

	return args, nil
}

// Parse the non-terminal <pars>
// <pars> ::= LPAR <args> RPAR | \epsilon
func (parser *ParserImpl) parsNT() (*Term, error) {
	// peek next token to determine case to do
	tok, err := parser.peekToken()
	if err != nil {
		return nil, ErrParser
	}

	// FIRST(<pars>) = {LPAR, \epsilon}
	// FOLLOW(<start>) = {COMMA, RPAR, EOF}
	switch tok.typ {
	// <term> ::= LPAR <args> RPAR
	case tokenLpar:
		// Consume LPAR
		tok, err := parser.nextToken()
		if err != nil || tok.typ != tokenLpar {
			return nil, ErrParser
		}

		// Parse <args>
		args, err := parser.argsNT()
		if err != nil {
			return nil, ErrParser
		}

		// Consume RPAR
		tok, err = parser.nextToken()
		if err != nil || tok.typ != tokenRpar {
			return nil, ErrParser
		}

		return args, nil

	// <term> ::= \epsilon
	case tokenComma, tokenRpar, tokenEOF:
		return nil, nil

	default:
		return nil, ErrParser

	}
}

// Parse the non-terminal <otherargs>
// <otherargs> ::= COMMA <args> | \epsilon
func (parser *ParserImpl) otherargsNT() (*Term, error) {
	// peek next token to determine case to do
	tok, err := parser.peekToken()
	if err != nil {
		return nil, ErrParser
	}

	// FIRST(<otherargs>) = {COMMA, \epsilon}
	// FOLLOW(<otherargs>) = {RPAR}
	switch tok.typ {
	// <otherargs> ::= COMMA <args>
	case tokenComma:
		// Consume COMMA
		tok, err := parser.nextToken()
		if err != nil || tok.typ != tokenComma {
			return nil, ErrParser
		}

		// Parse <args>
		args, err := parser.argsNT()
		if err != nil {
			return nil, ErrParser
		}

		return args, nil

	// <otherargs> ::= \epsilon
	case tokenRpar:
		return nil, nil

	default:
		return nil, ErrParser

	}
}
