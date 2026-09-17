package sexpr

import (
	"errors"
	"math/big" // You will need to use this package in your implementation.
)

// ErrEval is the error value returned by the Evaluator if the contains
// an invalid token.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrEval = errors.New("eval error")

func (expr *SExpr) quoteEval() (*SExpr, error) {
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
}

func (expr *SExpr) carEval() (*SExpr, error) {
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
}

func (expr *SExpr) cdrEval() (*SExpr, error) {
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
}

func (expr *SExpr) consEval() (*SExpr, error) {
	arg1 := expr.cdr
	arg2 := arg1.cdr

	// check that only given one list & that current cons cell is a list
	if arg1.isNil() || arg2.isNil() || arg2.cdr.cdr != nil || (arg2.cdr.isAtom() && !arg2.cdr.isNil() && arg2.cdr.atom.literal != "NIL") {
		return nil, ErrEval
	}

	eval1, err1 := (arg1.car).Eval()
	eval2, err2 := (arg2.car).Eval()
	if err1 != ErrEval && err2 != ErrEval {
		return mkConsCell(eval1, eval2), nil
	}
	return nil, ErrEval
}

func (expr *SExpr) lengthEval() (*SExpr, error) {
	arg1 := expr.cdr
	// check that only given one list & that current cons cell is a list
	if arg1.isNil() || arg1.cdr.cdr != nil || (arg1.cdr.isAtom() && !arg1.cdr.isNil() && arg1.cdr.atom.literal != "NIL") {
		return nil, ErrEval
	}

	eval, err := (arg1.car).Eval()
	if err != ErrEval && eval.isConsCell() {
		// length of nil is 0
		if eval.isNil() {
			return mkNumber(big.NewInt(0)), nil
		}

		var length int64 = 0
		for i := eval; i.cdr != nil; i = i.cdr {
			length++
		}
		return mkNumber(big.NewInt(length)), nil
	}
	return nil, ErrEval
}

func (expr *SExpr) sumEval() (*SExpr, error) {
	arg1 := expr.cdr
	// println(arg1.SExprString())

	// return 0 if no args
	if arg1.isNil() {
		return mkNumber(big.NewInt(0)), nil
	}

	// nil has a sum of 0
	eval, _ := (arg1.car).Eval()
	if !arg1.isConsCell() && eval.isNil() {
		return mkNumber(big.NewInt(0)), nil
	}

	total := big.NewInt(0)
	for i := arg1; i.car != nil; i = i.cdr {
		eval, err := (i.car).Eval()
		if err == ErrEval || !eval.isNumber() {
			return nil, ErrEval
		}
		total.Add(total, eval.atom.num)
	}
	return mkNumber(total), nil
}

func (expr *SExpr) productEval() (*SExpr, error) {
	arg1 := expr.cdr

	// return 1 if no args
	if arg1.isNil() {
		return mkNumber(big.NewInt(1)), nil
	}

	eval, _ := (arg1.car).Eval()
	// nil has a product of 1
	if !arg1.isConsCell() && eval.isNil() {
		return mkNumber(big.NewInt(1)), nil
	}

	total := big.NewInt(1)
	for i := arg1; i.car != nil; i = i.cdr {
		// println("i:", i.SExprString())
		eval, err := (i.car).Eval()
		if err == ErrEval || !eval.isNumber() {
			return nil, ErrEval
		}
		total.Mul(total, eval.atom.num)
		// println("Total: ", total)
	}
	return mkNumber(total), nil
}

func (expr *SExpr) atomEval() (*SExpr, error) {
	arg1 := expr.cdr
	// check that only given one list & that current cons cell is a list
	if arg1.isNil() || arg1.cdr.cdr != nil || (arg1.cdr.isAtom() && !arg1.cdr.isNil() && arg1.cdr.atom.literal != "NIL") {
		return nil, ErrEval
	}

	eval, err := (arg1.car).Eval()
	if err != ErrEval {
		if eval.isAtom() {
			return mkSymbolTrue(), nil
		}
		return mkNil(), nil
	}
	return nil, ErrEval
}

func (expr *SExpr) zeropEval() (*SExpr, error) {
	arg1 := expr.cdr
	// check that only given one list & that current cons cell is a list
	if arg1.isNil() || arg1.cdr.cdr != nil || (arg1.cdr.isAtom() && !arg1.cdr.isNil() && arg1.cdr.atom.literal != "NIL") {
		return nil, ErrEval
	}

	eval, err := (arg1.car).Eval()
	if err != ErrEval {
		if eval.atom.num.Cmp(big.NewInt(0)) == 0 {
			return mkSymbolTrue(), nil
		}
		return mkNil(), nil
	}
	return nil, ErrEval
}

func (expr *SExpr) listpEval() (*SExpr, error) {
	arg1 := expr.cdr
	// check that only given one list & that current cons cell is a list
	if arg1.isNil() || arg1.cdr.cdr != nil || (arg1.cdr.isAtom() && !arg1.cdr.isNil() && arg1.cdr.atom.literal != "NIL") {
		return nil, ErrEval
	}

	eval, err := (arg1.car).Eval()
	if err != ErrEval {
		// return error if eval isn't a number
		if !eval.isNumber() {
			return nil, ErrEval
		} else if eval.atom.num.Cmp(big.NewInt(0)) == 0 {
			return mkSymbolTrue(), nil
		}
		return mkNil(), nil
	}
	return nil, ErrEval
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
			println("default error")
			return nil, ErrEval
		}

		switch expr.car.SExprString() {
		case "QUOTE":
			return expr.quoteEval()

		case "CAR":
			return expr.carEval()

		case "CDR":
			return expr.cdrEval()

		case "CONS":
			return expr.consEval()

		case "LENGTH":
			return expr.lengthEval()

		case "+":
			return expr.sumEval()

		case "*":
			return expr.productEval()

		case "ATOM":
			return expr.atomEval()

		case "ZEROP":
			return expr.zeropEval()

		case "LISTP":
			return expr.listpEval()

		default:
			return nil, ErrEval
		}

		// is symbol or formatted weird
	} else {
		return nil, ErrEval
	}
}
