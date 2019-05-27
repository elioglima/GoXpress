package lib

import "fmt"

type Hello struct{}

func (y Hello) MethodTest(params interface{}) (value interface{}, err error) {
	fmt.Println("Executando MethodTest", params)
	return params, nil
}
