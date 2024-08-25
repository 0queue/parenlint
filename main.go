package main

import (
	"github.com/0queue/parenlint/parenlint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(parenlint.Analyzer())
}
