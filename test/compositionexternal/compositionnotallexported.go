package compositionexternal

import "github.com/mrsinham/goreset/test/compositionexternal/sub"

type compositionNotExported struct {
	field1 string
	sub.Sub2
}

func (c *compositionNotExported) test() {
	c.Sub2 = sub.Sub2{}
}
