package TransHttp

import (
	"time"
)

type recebeHttp struct {
	Usuario          string
	Senha            string
	DataProcessameto string
	Token            []byte
}

func NewRecebeHttp() *recebeHttp {

	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	t := time.Now()

	rth := &recebeHttp{}
	rth.DataProcessameto = t.UTC().Format(layout)
	return rth

}
