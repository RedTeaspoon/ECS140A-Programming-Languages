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

func (expr *SExpr) listEval() (*SExpr, error) {
	return nil, nil
}

func (expr *SExpr) consEval() (*SExpr, error) {
	return nil, nil
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

// func (expr *SExpr) innerEval() (*SExpr, error) {
// 	if expr.isNil() || expr.isAtom(){
// 		return expr.Eval()
// 	} else if expr.isConsCell() {
// 		// (1 . (2 . NIL))
// 		// start of list (NIL)
// 		// if expr.car.atom.typ == tokenSymbol && expr.car.atom.literal == "NIL" {
// 		// 	println("needed")
// 		// }

// 		if !expr.cdr.isConsCell() {
// 			return nil, ErrEval
// 		}

// 		if expr.car.isAtom() && expr.car.atom.typ == tokenSymbol && expr.car.atom.literal == "NIL"{

// 		}

// 		switch expr.car.SExprString() {
// 		case "QUOTE", "CAR", "CDR", "CONS", "CONS", "LENGTH", "ATOM", "LISTP", "ZEROP","+", "*":
// 			return expr.Eval()

// 		case :

// 		default:
// 			return nil, ErrEval
// 		}

// 		println(expr.car.SExprString())
// 		println(expr.cdr.SExprString())
// 		return nil, ErrEval

// 		// is symbol or formatted weird
// 	} else {
// 		return nil, ErrEval
// 	}
// }

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
		// (1 . (2 . NIL))
		// start of list (NIL)
		// if expr.car.atom.typ == tokenSymbol && expr.car.atom.literal == "NIL" {
		// 	println("needed")
		// }

		if !expr.cdr.isConsCell() {
			return nil, ErrEval
		}

		switch expr.car.SExprString() {
		case "QUOTE":
			if expr.cdr.cdr == nil || (expr.cdr.cdr.isNil() && expr.cdr.car != nil) || (expr.cdr.cdr.isAtom() && expr.cdr.cdr.atom.literal == "NIL") {
				return expr.cdr.car, nil
			}
			return nil, ErrEval

		case "CAR":
			if expr.cdr.isNil() || expr.cdr.cdr.cdr != nil {
				return nil, ErrEval
			}

			eval, err := (expr.cdr.car).Eval()
			if err != ErrEval {
				if eval.isNil() {
					return eval, nil
				}
				return eval.car, nil
			}
			return nil, ErrEval

		case "CDR":
			return expr.cdr.cdr, nil
		case "CONS":

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
