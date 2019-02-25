package TransHttp

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

var (
	TokenHttp     tokenHttp
	privado_token []byte
)

func Token() []byte {
	return privado_token
}

type tokenHttp struct {
	Usuario          string
	Senha            string
	DataProcessameto string
	DataAutenticacao time.Time
}

func NewTokenHttp() *tokenHttp {

	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	t := time.Now()

	rth := &tokenHttp{}
	rth.DataProcessameto = t.UTC().Format(layout)
	rth.DataAutenticacao = t
	return rth

}

func (tk *tokenHttp) Gerar() error {

	sjson, err := json.Marshal(tk)
	if err != nil {
		privado_token = []byte("")
		return err
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(sjson))

	/*

		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			fmt.Println("decode error:", err)
			return
		}

	*/

	privado_token = []byte(encoded)
	return nil

}
