package rst

import (
	"io"

	"github.com/mrsinham/rst/other"
)

type blu struct {
	tutu int
}

type customFiles string

type test struct {
	monAnge                         [5]int
	diane                           [4]customFiles
	moineau                         [5]other.MyOtherType
	blublu                          []int
	bluWithNoNil                    *blu `reset:"nonil" json:"none"`
	blu                             *blu
	fieldWithCustomTypeFromOtherPkg other.MyOtherType
	fieldWithCustomType             customFiles
	teststr                         string
	testMap                         map[string]int
	// testint is here
	testint int
	io.Writer
}
