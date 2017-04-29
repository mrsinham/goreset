package compositionexternal

import "github.com/mrsinham/zerogen/test/compositionexternal/sub"

type base struct {
	field1 string
	sub.Sub
}
