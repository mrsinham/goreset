package nonilexternalreset

import "github.com/mrsinham/goreset/test/nonil/external"

type nonilExternalWithReset struct {
	Field1 int
	external.ExternalReset
	interfaceReset
}

type interfaceReset interface {
	Reset()
}
