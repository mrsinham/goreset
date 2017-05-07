package simple

import "net/http"

type simple struct {
	field1  int
	field2  string
	field3  []string
	field4  [3]int
	field5  *http.Client
	field6  float32
	field7  float64
	field8  complex64
	field9  complex128
	field10 bool
	field11 byte
	field12 rune
	field13 uintptr
	field14 func()
}
