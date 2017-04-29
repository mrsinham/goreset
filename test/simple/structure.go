package simple

import "net/http"

type simple struct {
	field1 int
	field2 string
	field3 []string
	field4 [3]int
	field5 *http.Client
}
