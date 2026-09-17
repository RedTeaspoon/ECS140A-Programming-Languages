package branch

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// go test -coverprofile=temp.cov
// go tool cover -html=temp

func nestedBranchCount(checkingStmt *ast.BlockStmt) uint {
	var count uint = 0

	for _, stmt := range checkingStmt.List {
		if nestedStmt, ok := stmt.(*ast.IfStmt); ok {
			println("added to nested count")
			count++
			count += nestedBranchCount(nestedStmt.Body)
		}
		//  else if nestedStmt, ok := stmt.(*ast.SwitchStmt); ok {
		// 	count++
		// 	count += nestedBranchCount(nestedStmt.Body)
		// } else if nestedStmt, ok := stmt.(*ast.TypeSwitchStmt); ok {
		// 	count++
		// 	count += nestedBranchCount(nestedStmt.Body)
		// } else if nestedStmt, ok := stmt.(*ast.ForStmt); ok {
		// 	count++
		// 	count += nestedBranchCount(nestedStmt.Body)
		// } else if nestedStmt, ok := stmt.(*ast.RangeStmt); ok {
		// 	count++
		// 	count += nestedBranchCount(nestedStmt.Body)
		// }
	}
	return count
}

func branchCount(fn *ast.FuncDecl) uint {
	var count uint = 0

	for _, stmt := range fn.Body.List {
		if nestedStmt, ok := stmt.(*ast.IfStmt); ok {
			count++
			println(count)
			count += nestedBranchCount(nestedStmt.Body)
		} else if nestedStmt, ok := stmt.(*ast.SwitchStmt); ok {
			count++
			count += nestedBranchCount(nestedStmt.Body)
		} else if nestedStmt, ok := stmt.(*ast.TypeSwitchStmt); ok {
			count++
			count += nestedBranchCount(nestedStmt.Body)
		} else if nestedStmt, ok := stmt.(*ast.ForStmt); ok {
			count++
			count += nestedBranchCount(nestedStmt.Body)
		} else if nestedStmt, ok := stmt.(*ast.RangeStmt); ok {
			count++
			count += nestedBranchCount(nestedStmt.Body)
		}
	}
	return count
}

// ComputeBranchFactors returns a map from the name of the function in the given
// Go code to the number of branching statements it contains.
func ComputeBranchFactors(src string) map[string]uint {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	m := make(map[string]uint)
	for _, decl := range f.Decls {
		switch fn := decl.(type) {
		case *ast.FuncDecl:
			m[fn.Name.Name] = branchCount(fn)
		}
	}

	return m
}
