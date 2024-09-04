package parenlint

import (
	"cmp"
	"go/ast"
	"go/token"

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
		if ast.IsGenerated(file) {
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			fc, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			lparen := pass.Fset.Position(fc.Lparen)
			rparen := pass.Fset.Position(fc.Rparen)

			isSingleLine := true
			prevEnd := lparen

			edits := make([]analysis.TextEdit, 0)
			var errorMsg *string

			for i, e := range fc.Args {
				start := pass.Fset.Position(e.Pos())
				end := pass.Fset.Position(e.End())

				if i == 0 {
					isSingleLine = lparen.Line == start.Line
				}

				switch {
				case isSingleLine && prevEnd.Line != start.Line:
					errorMsg = cmp.Or(errorMsg, ptr(singleLineMsg))
				case !isSingleLine && prevEnd.Line == start.Line:
					errorMsg = cmp.Or(errorMsg, ptr(multiLineMsg))
				}

				// fixes will always turn into multiline
				if prevEnd.Line == start.Line {
					edits = append(edits, analysis.TextEdit{
						Pos:     e.Pos(),
						End:     token.NoPos,
						NewText: []byte("\n"),
					})
				}

				prevEnd = end
			}

			switch {
			case isSingleLine && prevEnd.Line != rparen.Line:
				errorMsg = cmp.Or(errorMsg, ptr(singleLineMsg))
			case !isSingleLine && prevEnd.Line == rparen.Line:
				errorMsg = cmp.Or(errorMsg, ptr(multiLineMsg))
			}

			if prevEnd.Line == rparen.Line {
				edits = append(edits, analysis.TextEdit{
					Pos:     fc.Rparen,
					End:     token.NoPos,
					NewText: []byte(",\n"),
				})
			}

			if errorMsg != nil {
				pass.Report(analysis.Diagnostic{
					Pos:     fc.Pos(),
					End:     fc.End(),
					Message: *errorMsg,
					URL:     "",
					SuggestedFixes: []analysis.SuggestedFix{
						{
							Message:   "Make function call multiline",
							TextEdits: edits,
						},
					},
					Related: []analysis.RelatedInformation{},
				})
			}

			return true
		})
	}

	return nil, nil
}

func ptr[T any](t T) *T {
	return &t
}
