package gosum

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"path"
	"testing"
)

func TestGoSumPermissions(t *testing.T) {
	var sourceDir = "."

	// load the offending source file (gosum.go)
	var files []*ast.File
	var fileset = token.NewFileSet()
	if f, err := parser.ParseFile(fileset, path.Join(sourceDir, "gosum.go"), nil, 0); err != nil {
		t.Error(err)
		return
	} else {
		files = append(files, f)
	}

	// execute the type-checker on the file
	var info = &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	var conf = types.Config{
		Importer: importer.For("source", nil),
	}

	_, err := conf.Check(sourceDir, fileset, files, info)
	if err != nil {
		t.Error(err)
		return
	}
}
