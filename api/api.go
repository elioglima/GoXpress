package main

import (
	"fmt"
	"reflect"
	"strings"
)

type klass struct{}

func (k klass) Hello(args ...string) string {
	return "Hello " + strings.Join(args, " ")
}

func main() {
	// bc := src.NewApi()
	// bc.Ini()

	v := reflect.TypeOf(reflect.ValueOf("klass")).Value
	fmt.Println(v)

}
