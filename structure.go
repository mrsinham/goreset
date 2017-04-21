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

type customFunc func() error

type test struct {
	//monAnge                         [5]int
	diane [4]customFiles
	//moineau                         [5]other.MyOtherType
	belleDiane                      map[customFiles][]other.MyOtherType `reset:"nonil" tutu:"test"`
	blublu                          []int
	blubluNoNil                     []other.MyOtherType       `reset:"nonil"`
	bluWithNoNil                    *othertoo.Diane           `reset:"nonil" json:"none"`
	channelDeMoineau                chan chan *othertoo.Diane `reset:"nonil"`
	blu                             *blu
	fieldWithCustomTypeFromOtherPkg other.MyOtherType
	testFunction                    []func(i int) error `reset:"nonil"`
	fieldWithCustomType             customFiles
	teststr                         string
	testMap                         map[string]int
	// testint is here
	testint int
	tutu    io.Writer
	io.Writer
}
