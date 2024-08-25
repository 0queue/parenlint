package parenlint

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "parenlint",
	Doc:  "Checks that arguments in call expressions are all on separate lines, and also not on the same lines as the parentheses",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			fc, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			lparen := pass.Fset.Position(fc.Lparen).Line
			rparen := pass.Fset.Position(fc.Rparen).Line

			if lparen == rparen {
				return true
			}

			var prevEnd *token.Position
			for _, e := range fc.Args {
				start := pass.Fset.Position(e.Pos())
				end := pass.Fset.Position(e.End())

				switch {
				case start.Line == lparen:
					pass.Reportf(e.Pos(), "Argument on same line as left paren")
				case end.Line == rparen:
					pass.Reportf(e.Pos(), "Argument on same line as right paren")
				case prevEnd != nil && start.Line == prevEnd.Line:
					pass.Reportf(e.Pos(), "Argument on same line as previous argument")
				}

				prevEnd = &end
			}

			return true
		})
	}

	return nil, nil
}
