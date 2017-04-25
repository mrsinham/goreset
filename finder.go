package main

import (
	"go/ast"
	"go/token"

	"strings"

	"io"
	"os"
)

type structFinder struct {
	name     string
	filename string
	found    []*ast.TypeSpec
}

func newStructFinder(name string, filename string) *structFinder {
	return &structFinder{name: name, filename: filename}
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
	allFiles []*ast.File,
	dirname string,
	pkgName string,
	fileName string,
	structToFind string,
	write bool,
) error {
	sf := newStructFinder(structToFind, fileName)
	ast.Inspect(currentFile, sf.find)

	// structure not found
	if len(sf.matches()) == 0 {
		return nil
	}

	var writer io.Writer
	if !write {
		writer = os.Stdout
	} else {
		resetFile := strings.Replace(fileName, ".go", "_reset.go", 1)
		// delete if needed
		_ = os.Remove(resetFile)

		// write to a file
		var err error
		writer, err = os.OpenFile(resetFile, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			return err
		}
	}

	g := newGenerator(sf.matches(), dirname, allFiles, set, pkgName, writer)
	return g.do()
}
