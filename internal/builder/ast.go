package builder

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func (a *App) ReadAst() error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, a.Source, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	ast.Inspect(f, func(n ast.Node) bool {
		// Called recursively.
		ast.Print(fset, n)
		return true
	})

	return nil
}
