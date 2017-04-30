package nonil

import "net/http"

type nonil struct {
	field1 [4]int
	field2 *http.Request `zerogen:"nonil"`
}
