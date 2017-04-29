package composition

type composition struct {
	subfield1 int
	subfield2 string
}

type composited struct {
	composition
	fieldsimple string
}
