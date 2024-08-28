package parenlint

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "parenlint",
		Doc:  "Checks that arguments in call expressions are all on separate lines, and also not on the same lines as the parentheses",
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

			lparen := pass.Fset.Position(fc.Lparen).Line
			rparen := pass.Fset.Position(fc.Rparen).Line

			if lparen == rparen {
				return true
			}

			type report struct {
				pos token.Pos
				msg string
			}

			reports := make([]report, 0)

			// only report if we are not in a "hanging final arg" situation
			isValidHangingFinalArg := true

			var prevEnd *token.Position
			for i, e := range fc.Args {
				start := pass.Fset.Position(e.Pos())
				end := pass.Fset.Position(e.End())

				if start.Line != lparen {
					isValidHangingFinalArg = false
				}

				if i == len(fc.Args)-1 && end.Line != rparen {
					isValidHangingFinalArg = false
				}

				switch {
				case start.Line == lparen:
					reports = append(reports, report{
						pos: e.Pos(),
						msg: "Argument on same line as left paren",
					})
				case end.Line == rparen:
					reports = append(reports, report{
						pos: e.Pos(),
						msg: "Argument on same line as right paren",
					})
				case prevEnd != nil && start.Line == prevEnd.Line:
					reports = append(reports, report{
						pos: e.Pos(),
						msg: "Argument on same line as previous argument",
					})
				}

				prevEnd = &end
			}

			if isValidHangingFinalArg {
				return true
			}

			for _, r := range reports {
				pass.Reportf(r.pos, r.msg)
			}

			return true
		})
	}

	return nil, nil
}
