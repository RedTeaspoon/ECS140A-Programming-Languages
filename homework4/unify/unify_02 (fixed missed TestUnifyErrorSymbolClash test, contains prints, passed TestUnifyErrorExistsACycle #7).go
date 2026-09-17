package unify

import (
	"errors"
	// "hw4/disjointset"
	"hw4/term"
)

// ErrUnifier is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrUnifier = errors.New("unifier error")

// UnifyResult is the result of unification. For example, for a variable term
// `s`, `UnifyResult[s]` is the term which `s` is unified with.
type UnifyResult map[*term.Term]*term.Term

type UnifierImpl struct {
	unified UnifyResult
}

// returns true if termA exists in Args of termB; termB is the compound
func (unifier *UnifierImpl) cycleExists(termA, termB *term.Term) bool {
	println("Check cycle:", termA.String(), termB.String())
	if termA == termB {
		return true
	} else if termB.Typ == term.TermCompound {
		for _, arg := range termB.Args {
			if termA == arg || (arg.Typ == term.TermCompound && unifier.cycleExists(termA, arg)) || unifier.cycleExists(termA, arg) {
				return true
			}
		}

		// need to check that neither termA maps to compound with termB or termB maps to compound with termA
	} else if termB.Typ == term.TermVariable {
		// _, exists := unifier.unified[termB]
		if mapTerm, exists := unifier.unified[termB]; exists {
			// println("Mapterm:", mapTerm.String())
			// if mapTerm.String() == "A"{
			// 	println()
			// }

			if mapTerm == termA {
				return true
			} else if mapTerm.Typ == term.TermCompound {
				for _, arg := range mapTerm.Args {
					if termA == arg || (arg.Typ == term.TermCompound && unifier.cycleExists(termA, arg)) || unifier.cycleExists(termA, arg) {
						return true
					}
				}
			} else if mapTerm.Typ == term.TermVariable && unifier.cycleExists(termA, mapTerm) {
				// println("Returning true:", termA.String(), mapTerm.String())
				return true
			}
		}

		// if termB.String() == "A" {
		// 	println("B")
		// 	return true
		// }

		if mapTerm, exists := unifier.unified[termA]; exists {
			// if termB.String() == "A" {
			// 	println("A")
			// 	return true
			// }
			println("Mapterm2:", mapTerm.String())
			if mapTerm == termB {
				return true
			} else if mapTerm.Typ == term.TermCompound {
				for _, arg := range mapTerm.Args {
					if termB == arg || (arg.Typ == term.TermCompound && unifier.cycleExists(termB, arg)) || unifier.cycleExists(termA, arg) {
						return true
					}
				}
			} else if mapTerm.Typ == term.TermVariable && unifier.cycleExists(termB, mapTerm) {
				return true
			}
		}
	}

	return false
}

// Unifier is the interface for the term unifier.
// Do not change the definition of this interface
type Unifier interface {
	Unify(*term.Term, *term.Term) (UnifyResult, error)
}

// NewUnifier creates a struct of a type that satisfies the Unifier interface.
func NewUnifier() Unifier {
	unified := make(map[*term.Term]*term.Term)
	return &UnifierImpl{unified}
}

func (unifier *UnifierImpl) Unify(termA *term.Term, termB *term.Term) (UnifyResult, error) {
	println("Unify:", termA.String(), termB.String())
	switch termA.Typ {
	case term.TermAtom:
		if termB.Typ == term.TermVariable {
			unifier.unified[termB] = termA
			return unifier.unified, nil

		} else {
			return nil, ErrUnifier
		}

	case term.TermNumber:
		if termB.Typ == term.TermVariable {
			unifier.unified[termB] = termA
			return unifier.unified, nil

		} else {
			return nil, ErrUnifier
		}

	case term.TermVariable:
		if termB.Typ != term.TermCompound && termB.Typ != term.TermVariable {
			unifier.unified[termA] = termB
			return unifier.unified, nil

		} else if !unifier.cycleExists(termA, termB) {
			unifier.unified[termA] = termB
			return unifier.unified, nil

		} else {
			return nil, ErrUnifier
		}

	case term.TermCompound:
		if termB.Typ == term.TermVariable && !unifier.cycleExists(termB, termA) {
			unifier.unified[termB] = termA
			return unifier.unified, nil

		} else if termB.Typ != term.TermCompound || termA.Functor != termB.Functor || len(termA.Args) != len(termB.Args) {
			return nil, ErrUnifier
		}

		for index := range termA.Args {
			if termA.Args[index].Typ == termB.Args[index].Typ && (termA.Args[index].Typ == term.TermAtom || termA.Args[index].Typ == term.TermNumber) {
				continue
			}
			_, err := unifier.Unify(termA.Args[index], termB.Args[index])
			if err != nil {
				return nil, ErrUnifier
			}
		}
	}
	// println(unifier.unified[&term.Term{Typ: term.TermVariable, Literal: "A"}].String())
	return unifier.unified, nil
}
