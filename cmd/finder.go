package main

import (
	"go/ast"
	"go/token"

	"strings"

	"io"
	"os"
)

type structFinder struct {
	name  string
	found []*ast.TypeSpec
}

func newStructFinder(name string) *structFinder {
	return &structFinder{name: name}
}

func (s *structFinder) find(n ast.Node) bool {

	var ts *ast.TypeSpec
	var ok bool

	if ts, ok = n.(*ast.TypeSpec); !ok {
		return true
	}

	if ts.Name == nil {
		return true
	}

	if !strings.Contains(ts.Name.Name, s.name) {
		return false
	}

	s.found = append(s.found, ts)

	return false
}

func (s *structFinder) matches() []*ast.TypeSpec {
	return s.found
}

func generate(
	set *token.FileSet,
	currentFile *ast.File,
	dirname string,
	pkgName string,
	fileName string,
	structToFind string,
	write bool,
) error {
	sf := newStructFinder(structToFind)
	ast.Inspect(currentFile, sf.find)

	var writer io.Writer
	if !write {
		writer = os.Stdout
	} else {
		// write to a file
		var err error
		writer, err = os.OpenFile(strings.Replace(fileName, ".go", "_reset.go", 1), os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			return err
		}
	}

	g := newGenerator(sf.matches(), dirname, []*ast.File{currentFile}, set, pkgName, writer)
	return g.do()
}
