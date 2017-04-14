package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/jawher/mow.cli"
)

func main() {

	app := cli.App("reset", "generate reset method")
	chosenDir := app.StringArg("DIR", "", "package to walk to")
	chosenStruct := app.StringArg("STRUCTURE", "", "structure to attach to Reset() method to")

	exitOnError := func(err error) {
		fmt.Println(err)
		os.Exit(1)
	}
	var err error
	app.Action = func() {
		err = parsePackage(chosenDir, chosenStruct)
		if err != nil {
			exitOnError(err)
		}
	}
	err = app.Run(os.Args)
	if err != nil {
		exitOnError(err)
	}

	spew.Dump(chosenDir, chosenStruct)

}

// parsePackage launchs the generation
func parsePackage(dir *string, structure *string) error {

	if dir == nil {
		return errors.New("no directory submitted")
	}

	if strings.TrimSpace(*dir) == "" {
		return errors.New("directory empty submitted")
	}

	fset := token.NewFileSet()
	f, err := parser.ParseDir(fset, *dir, nil, 0)
	if err != nil {
		return err
	}

	ast.Print(fset, f)

	return nil
}
