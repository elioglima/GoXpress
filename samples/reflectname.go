package main

import (
	"GoXpress/samples/lib"
	"errors"
	"fmt"
	"reflect"
)

type RegistryData struct {
	Name     string
	Path     string
	FileName string
	Type     reflect.Type
}

var typeRegistry = make(map[int]RegistryData)
var ItemCount int = 0

func registerType(typedNil interface{}) {

	t := reflect.TypeOf(typedNil).Elem()
	data := RegistryData{}
	data.Name = t.Name()
	data.Path = t.PkgPath()
	data.FileName = t.PkgPath() + "/" + t.Name()
	data.Type = t

	ItemCount++
	typeRegistry[ItemCount] = data
	fmt.Println(data)
}

func init() {
	registerType((*lib.Hello)(nil))
}

func makeInstance(name string) (interface{}, error) {
	for _, data := range typeRegistry {
		if data.Name == name {
			return reflect.New(data.Type).Elem().Interface(), nil
		}
	}

	return nil, errors.New("Struct not found")
}

func invoke(any interface{}, name string, args ...interface{}) (reflect.Value, error) {
	method := reflect.ValueOf(any).MethodByName(name)
	methodType := method.Type()
	numIn := methodType.NumIn()
	if numIn > len(args) {
		return reflect.ValueOf(nil), fmt.Errorf("Method %s must have minimum %d params. Have %d", name, numIn, len(args))
	}
	if numIn != len(args) && !methodType.IsVariadic() {
		return reflect.ValueOf(nil), fmt.Errorf("Method %s must have %d params. Have %d", name, numIn, len(args))
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		var inType reflect.Type
		if methodType.IsVariadic() && i >= numIn-1 {
			inType = methodType.In(numIn - 1).Elem()
		} else {
			inType = methodType.In(i)
		}
		argValue := reflect.ValueOf(args[i])
		if !argValue.IsValid() {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argValue.String())
		}
		argType := argValue.Type()
		if argType.ConvertibleTo(inType) {
			in[i] = argValue.Convert(inType)
		} else {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argType)
		}
	}
	return method.Call(in)[0], nil
}

func main() {
	s, err := makeInstance("Hello")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	value, err := invoke(s, "MethodTest", "parametros")
	fmt.Println(value, err)
}
