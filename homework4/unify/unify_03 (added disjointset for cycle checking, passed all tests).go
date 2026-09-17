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
	unifiedMap UnifyResult // used to check answers
	unified    DisjointSet
}

// returns true if termA exists in Args of termB; termB is the compound
func (unifier *UnifierImpl) cycleExists(termA, termB *term.Term) bool {
	// println("Check cycle:", termA.String(), termB.String())
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
		termARep := unifier.unified.FindSet(termA)
		termBRep := unifier.unified.FindSet(termB)

		if termARep == termBRep {
			return true
		}
	}
	return false
}

// termA is variable, termB is the compound
func (unifier *UnifierImpl) unifyCompound(termA, termB *term.Term) {
	for _, arg := range termB.Args {
		if arg.Typ == term.TermCompound {
			unifier.unifyCompound(termA, arg)
		} else if arg.Typ == term.TermVariable {
			unifier.unified.UnionSet(termA, arg)
		}
	}
}

// Unifier is the interface for the term unifier.
// Do not change the definition of this interface
type Unifier interface {
	Unify(*term.Term, *term.Term) (UnifyResult, error)
}

// NewUnifier creates a struct of a type that satisfies the Unifier interface.
func NewUnifier() Unifier {
	unifiedMap := make(map[*term.Term]*term.Term)
	unified := NewDisjointSet()
	return &UnifierImpl{unifiedMap, unified}
}

func (unifier *UnifierImpl) Unify(termA *term.Term, termB *term.Term) (UnifyResult, error) {
	// println("Unify:", termA.String(), termB.String())
	switch termA.Typ {
	case term.TermAtom:
		if termB.Typ == term.TermVariable {
			unifier.unified.UnionSet(termA, termB)
			unifier.unifiedMap[termB] = termA
			return unifier.unifiedMap, nil

		} else {
			return nil, ErrUnifier
		}

	case term.TermNumber:
		if termB.Typ == term.TermVariable {
			unifier.unified.UnionSet(termA, termB)
			unifier.unifiedMap[termB] = termA
			return unifier.unifiedMap, nil

		} else {
			return nil, ErrUnifier
		}

	case term.TermVariable:
		if termB.Typ != term.TermCompound && termB.Typ != term.TermVariable {
			unifier.unified.UnionSet(termA, termB)
			unifier.unifiedMap[termA] = termB
			return unifier.unifiedMap, nil

		} else if !unifier.cycleExists(termA, termB) {
			unifier.unified.UnionSet(termA, termB)
			unifier.unifyCompound(termA, termB)
			unifier.unifiedMap[termA] = termB
			return unifier.unifiedMap, nil

		} else {
			return nil, ErrUnifier
		}

	case term.TermCompound:
		if termB.Typ == term.TermVariable && !unifier.cycleExists(termB, termA) {
			unifier.unified.UnionSet(termA, termB)
			unifier.unifyCompound(termB, termA)
			unifier.unifiedMap[termB] = termA
			return unifier.unifiedMap, nil

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
	return unifier.unifiedMap, nil
}
