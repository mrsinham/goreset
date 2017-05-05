package structnoname

import "net/http"

type structnoname struct {
	field1 int
	field2 struct {
		tutu1, tutu2 *http.Request
		tutu3        struct {
			file string
		}
	}
}
