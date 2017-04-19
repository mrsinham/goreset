package rst

import (
	"io"

	"github.com/mrsinham/rst/other"
	"github.com/mrsinham/rst/othertoo"
)

type blu struct {
	tutu int
}

type customFiles int

type test struct {
	//monAnge                         [5]int
	diane [4]customFiles
	//moineau                         [5]other.MyOtherType
	belleDiane                      map[customFiles]other.MyOtherType `reset:"nonil" tutu:"test"`
	blublu                          []int
	blubluNoNil                     []other.MyOtherType       `reset:"nonil"`
	bluWithNoNil                    *othertoo.Diane           `reset:"nonil" json:"none"`
	channelDeMoineau                chan chan *othertoo.Diane `reset:"nonil"`
	blu                             *blu
	fieldWithCustomTypeFromOtherPkg other.MyOtherType
	fieldWithCustomType             customFiles
	teststr                         string
	testMap                         map[string]int
	// testint is here
	testint int
	io.Writer
}
