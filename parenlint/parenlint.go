package parenlint

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const (
	singleLineMsg = "Single line function call with arguments on multiple lines"
	multiLineMsg  = "Multiline function call with multiple arguments on single line"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "parenlint",
		Doc:  "Checks that function call arguments are all in a single line, or all on multiple lines.",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			fc, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			lparen := pass.Fset.Position(fc.Lparen)
			rparen := pass.Fset.Position(fc.Rparen)

			isSingleLine := true
			prevEnd := lparen

			for i, e := range fc.Args {
				start := pass.Fset.Position(e.Pos())
				end := pass.Fset.Position(e.End())

				if i == 0 {
					isSingleLine = lparen.Line == start.Line
				}

				switch {
				case isSingleLine && prevEnd.Line != start.Line:
					pass.Reportf(fc.Pos(), singleLineMsg)
					return true
				case !isSingleLine && prevEnd.Line == start.Line:
					pass.Reportf(fc.Pos(), multiLineMsg)
					return true
				}

				prevEnd = end
			}

			switch {
			case isSingleLine && prevEnd.Line != rparen.Line:
				pass.Reportf(fc.Pos(), singleLineMsg)
				return true
			case !isSingleLine && prevEnd.Line == rparen.Line:
				pass.Reportf(fc.Pos(), multiLineMsg)
				return true
			}

			return true
		})
	}

	return nil, nil
}
