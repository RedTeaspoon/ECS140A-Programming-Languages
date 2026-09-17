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

//
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

// Parser is the interface for the term parser.
// Do not change the definition of this interface.
type Parser interface {
	Parse(string) (*Term, error)
}

// NewParser creates a struct of a type that satisfies the Parser interface.
func NewParser() Parser {
	return &ParserImpl{}
}

func (parser ParserImpl) Parse(input string) (*Term, error) {
	lexer := newLexer(input)

	// There are two cases, so we peek the next token.
	tok, error := parser.peekToken()
	// \eplison, so ok
	if error == nil {
		return nil, nil
	}

	// <start> -> <term>
	// Parse <bracket>.
	expr, err := parser.termNT()
	if err != nil {
		return nil, ErrParser
	}

	// FOLLOW(<start>) = $
	// Check the next token is the endmarker $, there should be nothing left
	// after parsing <start>.
	if nextTok, err := parser.nextToken(); err != nil || nextTok.typ != tokenEOF {
		return nil, ErrParser
	}

	return expr, nil
}

// Parse the non-terminal <term>.
// <term> ::= ATOM <pars> | NUM | VAR
func (p *ParserImpl) termNT() (*Term, error) {
	// There are two cases, so we peek the next token.
	tok, err := p.peekToken()
	if err != nil {
		return nil, ErrParser
	}

	switch tok.typ {

	// FIRST(<term>) = {LBRKT, NUMBER}
	// <tail> -> <term> <tail>
	case tokenAtom:
		// Parse ATOM
		term, err := p.termNT()
		if err != nil {
			return nil, ErrParser
		}

		// Parse <pars>.
		expr, err := p.parsNT()
		if err != nil {
			return nil, ErrParser
		}

		// if <pars> returns epsilon, make ATOM TermAtom, otherwise make TermCompound

		return expr, nil

	// FIRST(<term>)   =  {ATOM, NUM, VAR}
	// FOLLOW(<term>)  =  {}
	// \epsilon in FIRST(<tail>) and RBRKT is in FOLLOW(<tail>).
	// <tail> -> \epsilon
	case tokenNumber:
		return &Expr{Atom: nil, Bracket: []*Expr{}}, nil

	case tokenVariable:
		return &Expr{Atom: nil, Bracket: []*Expr{}}, nil

	default:
		return nil, ErrParser

	}
}
