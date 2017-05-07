package structnonamepointer

type structnonamepointer struct {
	field1 int
	field2 *struct {
		field1 int
		field2 *struct {
			test int
		}
		field3 struct {
			test string
		}
	} `reset:"nonil"`
}
