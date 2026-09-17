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

func (expr *SExpr) Eval() (*SExpr, error) {
	if expr.isAtom() {
		return expr, nil
	} else if expr.isConsCell() {
		// if expr.car.isConsCell() {
		// 	return expr.car.Eval()
		// }

		switch expr.car.SExprString() {
		case "QUOTE":
			if expr.cdr.isConsCell() {
				return expr.cdr.car, nil
			}
			return nil, ErrEval

		case "CAR":
			return expr.cdr.car, nil
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
