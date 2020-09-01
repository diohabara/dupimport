package dupimport

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "dupimport is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "dupimport",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		usedPath := map[string]bool{}
		for _, importSpec := range file.Imports {
			path, err := strconv.Unquote(ImportSpec.Path.Value)
			if err != nil {
				return err, nil
			}
			if _, ok := usedPath[path]; ok {
				pass.Reportf(importSpec.Pos(), "dulicate import: %s", path)
			} else {
				usedPath[path] = true
			}
		}
	}

	return nil, nil
}
