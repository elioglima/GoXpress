package libs

import (
	"reflect"
	"strings"
)

type ReflectsReturn struct {
	dados reflect.Value
}

func (RRi *ReflectsReturn) Load(input reflect.Value) {
	RRi.dados = input
}

func (RRi *ReflectsReturn) ToStr() string {
	return RRi.dados.String()
}

type Reflects struct {
}

func (R *Reflects) F(Name string) ReflectsReturn {
	RRi := ReflectsReturn{}
	s := reflect.ValueOf(R).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		// typeOfT.Field(i).Name, f.Type(), f.Interface()
		if typeOfT.Field(i).Name == Name {
			RRi.Load(f)
		}
	}

	return RRi
}

func InterfaceToName(i interface{}) string {
	val := reflect.Indirect(reflect.ValueOf(i))
	return strings.ToLower(val.Type().Name())
}
