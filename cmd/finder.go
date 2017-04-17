package main

import (
	"go/ast"
	"go/token"

	"strings"

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

func findStructures(set *token.FileSet, currentFile *ast.File, dirname string, pkgName string, fileName string, structToFind string) error {
	//spew.Dump(currentFile)
	sf := newStructFinder(structToFind)
	ast.Inspect(currentFile, sf.find)

	g := newGenerator(sf.matches(), dirname, []*ast.File{currentFile}, set, pkgName, os.Stdout)
	return g.do()
}
