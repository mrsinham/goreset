package external

type ExternalReset struct {
	field1 int
}

func (e *ExternalReset) Reset() {
	e.field1 = 0
}
