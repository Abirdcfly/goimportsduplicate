package importsduplicate

import (
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "imports_duplicate",
	Doc:      "detect duplicate package imports",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		packageAll := make(map[string]bool, 0)
		for _, i := range f.Imports {
			if i.Name.String() == "_" {
				continue
			}
			path, err := strconv.Unquote(i.Path.Value)
			if err != nil {
				pass.Reportf(i.Path.Pos(), "import not quoted")
			}
			if packageAll[path] {
				pass.Reportf(i.Path.Pos(), "duplicate import %s", path)
			}
			packageAll[path] = true
		}
	}
	return nil, nil
}
