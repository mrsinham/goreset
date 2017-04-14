package reset

import "go/ast"

type Field struct {
	Name string
	Type ast.StructType
}

type structure struct {
	fieldList []*Field
	Name string
}


