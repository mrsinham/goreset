package compositionexternal

import "github.com/mrsinham/goreset/test/compositionexternal/sub"

type base struct {
	field1 string
	sub.Sub
}
