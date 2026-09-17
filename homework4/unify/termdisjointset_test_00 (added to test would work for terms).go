package unify

import (
	"fmt"
	"hw4/term"
	"testing"
)

func TestDisjointSetSimple(t *testing.T) {
	func() {
		defer func() {
			if recover() != nil {
				t.Errorf("DisjointSetSimple panicked")
			}
		}()
		s := NewDisjointSet()
		parser := term.NewParser()
		term1, _ := parser.Parse("X")
		term2, _ := parser.Parse("Y")

		fmt.Println("Checking: ", s.FindSet(term1).String(), s.FindSet(term2).String())

		r := s.UnionSet(term1, term2)
		if r != term1 && r != term2 {
			t.Errorf("Union error, actual %s", r.String())
		}
		if s.FindSet(term1) != s.FindSet(term2) {
			t.Errorf("Expected true, actual false")
		}
	}()
}
