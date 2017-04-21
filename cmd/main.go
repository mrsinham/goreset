package main

import (
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"os"

	"strings"

	"github.com/jawher/mow.cli"
)

func main() {

	app := cli.App("reset", "generate reset method")
	chosenPackage := app.StringArg("PKG", "", "package to walk to")
	chosenStruct := app.StringArg("STRUCTURE", "", "structure to attach to Reset() method to")

	// write
	write := app.BoolOpt("w", false, "writes the result in file")

	exitOnError := func(err error) {
		fmt.Println(err)
		os.Exit(1)
	}
	var err error
	app.Action = func() {
		err = parsePackage(chosenPackage, chosenStruct, write)
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
func parsePackage(pkg *string, structure *string, write *bool) error {

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

	fset := token.NewFileSet()
	f, err := parser.ParseDir(fset, pkgdir, nil, 0)
	if err != nil {
		return err
	}

	for i := range f {
		for j := range f[i].Files {
			if !strings.Contains(j, "_reset.go") {
				err = generate(fset, f[i].Files[j], pkgdir, i, j, *structure, writeToFile)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
