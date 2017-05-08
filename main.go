package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"io"

	"github.com/jawher/mow.cli"
)

func main() {

	app := cli.App("goreset", "generate reset method")
	chosenPackage := app.StringArg("PKG", "", "package to walk to")
	chosenStruct := app.StringArg("STRUCTURE", "", "structure to attach the Reset() method to")

	// writeType
	write := app.BoolOpt("w write", false, "writes the result in file")

	exitOnError := func(err error) {
		fmt.Println(err)
		os.Exit(1)
	}
	var err error
	app.Action = func() {
		err = parsePackage(chosenPackage, chosenStruct, write, nil)
		if err != nil {
			exitOnError(err)
		}
	}
	err = app.Run(os.Args)
	if err != nil {
		exitOnError(err)
	}

}

// parsePackage launchs the generation
func parsePackage(pkg *string, structure *string, write *bool, customWriter io.Writer) error {

	if pkg == nil {
		return errors.New("no directory submitted")
	}

	if strings.TrimSpace(*pkg) == "" {
		return errors.New("directory empty submitted")
	}

	var writeToFile bool
	if write != nil && *write {
		writeToFile = true
	}

	// get the path of the package
	pkgdir := os.Getenv("GOPATH") + "/src/" + *pkg

	// reinstall package to be sure that we are uptodate

	c := exec.Command(runtime.GOROOT()+"/bin/go", []string{"install", *pkg}...)
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	var f map[string]*ast.Package
	f, err = parser.ParseDir(fset, pkgdir, nil, 0)
	if err != nil {
		return err
	}

	for i := range f {
		var files []*ast.File
		for j := range f[i].Files {
			files = append(files, f[i].Files[j])

		}

		for j := range f[i].Files {
			if !strings.Contains(j, "_reset.go") {
				err = generate(fset, f[i].Files[j], files, pkgdir, i, j, *structure, writeToFile, customWriter)
				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}
