package requests

import "regexp"

var idURLMap = map[string]string{
	"pout":  "/v1/payouts/",
}

// pout_GLBIjRm3dN3i4Y

var idRegex = regexp.MustCompile("^([a-z]{2,5})_[a-zA-Z0-9]{14}$")
