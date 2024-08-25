package parenlint

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestParenLint(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer())
}
