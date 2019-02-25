package TransHttp

import (
	"time"
)

type RetornoHttpAPI struct {
	DataProcessameto string
	Token            []byte
	HttpStatus       int
	HttpResponse     string
}

func NewRetornoHttpAPI() *RetornoHttpAPI {

	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	t := time.Now()

	rth := &RetornoHttpAPI{}
	rth.DataProcessameto = t.UTC().Format(layout)
	return rth
}
