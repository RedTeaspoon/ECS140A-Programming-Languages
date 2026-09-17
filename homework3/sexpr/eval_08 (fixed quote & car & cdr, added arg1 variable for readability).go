package sexpr

import (
	"errors"
	//"math/big" // You will need to use this package in your implementation.
)

// ErrEval is the error value returned by the Evaluator if the contains
// an invalid token.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrEval = errors.New("eval error")

func (expr *SExpr) consEval() (*SExpr, error) {
	arg1 := expr.cdr
	arg2 := arg1.cdr
	// check that only given one list & that current cons cell is a list
	if arg1.isNil() || (arg1.cdr.isAtom() && arg1.cdr.atom.literal != "NIL") {
		return nil, ErrEval
	}
	println(arg2)

	eval, err := (expr.cdr.car).Eval()
	if err != ErrEval {
		if eval.isNil() {
			return eval, nil
		}
		return eval.cdr, nil
	}
	return nil, ErrEval
}

func (expr *SExpr) lengthEval() (*SExpr, error) {
	return nil, nil
}

func (expr *SExpr) atomEval() (*SExpr, error) {
	return nil, nil
}

func (expr *SExpr) listpEval() (*SExpr, error) {
	return nil, nil
}

func (expr *SExpr) zeropEval() (*SExpr, error) {
	return nil, nil
}

func (expr *SExpr) addEval() (*SExpr, error) {
	return nil, nil
}

func (expr *SExpr) multiplyEval() (*SExpr, error) {
	return nil, nil
}

func (expr *SExpr) Eval() (*SExpr, error) {
	if expr.isNil() {
		return expr, nil
	} else if expr.isAtom() {
		if expr.atom.typ == tokenNumber {
			return expr, nil
		} else if expr.atom.typ == tokenSymbol && expr.atom.literal == "NIL" {
			return mkNil(), nil
		}
		return nil, ErrEval

	} else if expr.isConsCell() {

		if !expr.cdr.isConsCell() {
			return nil, ErrEval
		}

		switch expr.car.SExprString() {
		case "QUOTE":
			arg1 := expr.cdr
			// check that at least 1 arg is given
			if arg1.cdr == nil {
				return nil, ErrEval
			}

			// check that only given one list & that current cons cell is a list
			if (arg1.cdr.isNil() && arg1.car != nil) || (arg1.cdr.isAtom() && arg1.cdr.atom.literal == "NIL") {
				return arg1.car, nil
			}
			return nil, ErrEval

		case "CAR":
			arg1 := expr.cdr
			// check that only given one list & that current cons cell is a list
			if arg1.isNil() || arg1.cdr.cdr != nil || (arg1.cdr.isAtom() && !arg1.cdr.isNil() && arg1.cdr.atom.literal != "NIL") {
				return nil, ErrEval
			}

			eval, err := (arg1.car).Eval()
			if err != ErrEval {
				if eval.isNil() {
					return eval, nil
				}
				return eval.car, nil
			}
			return nil, ErrEval

		case "CDR":
			arg1 := expr.cdr
			// check that only given one list & that current cons cell is a list
			if arg1.isNil() || arg1.cdr.cdr != nil || (arg1.cdr.isAtom() && !arg1.cdr.isNil() && arg1.cdr.atom.literal != "NIL") {
				return nil, ErrEval
			}

			eval, err := (arg1.car).Eval()
			if err != ErrEval {
				if eval.isNil() {
					return eval, nil
				}
				return eval.cdr, nil
			}
			return nil, ErrEval
		case "CONS":
			return expr.consEval()
		case "LENGTH":

		case "ATOM":

		case "LISTP":

		case "ZEROP":

		case "+":

		case "*":

		default:
			return nil, ErrEval
		}

		println(expr.car.SExprString())
		println(expr.cdr.SExprString())
		return nil, ErrEval

		// is symbol or formatted weird
	} else {
		return nil, ErrEval
	}
}
