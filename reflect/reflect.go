package main

import (
	"fmt"
	"reflect"
)

type Conf struct {
	Id int `gconf:"service:id"`
	Desc string  `gconf:"service:desc"`
	Address []string `gconf:"service:address"`
	HashMap  map[string]string `gconf:"service:hashmap"`
}

func main() {
	var f = Conf{}
	var rv  = reflect.ValueOf(&f).Elem()
	var rt  = rv.Type()
	var n = rv.NumField()
	for i := 0; i <  n ;i++ {
		var field = rv.Field(i)
		var fieldtype = rt.Field(i)
		fmt.Println(fieldtype.Tag.Get("gconf"))
		switch field.Kind() {
		    case reflect.Int:
				field.SetInt(10)
			case reflect.String:
				field.SetString("hello")
		    case reflect.Slice:
                var elemType = field.Type()
                var sl = reflect.MakeSlice(elemType, 0 ,2)
                sl = reflect.Append(sl, reflect.ValueOf("this is "),reflect.ValueOf("fun"))
                field.Set(sl)
		    case reflect.Map:
				var keyType = field.Type().Key()
				var valueType = field.Type().Elem()
				fmt.Println(keyType.String(), valueType.String())
				var sm = reflect.MakeMap(field.Type())
				sm.SetMapIndex(reflect.ValueOf("123"), reflect.ValueOf("456"))
				sm.SetMapIndex(reflect.ValueOf("key"), reflect.ValueOf("key2"))
				field.Set(sm)
		}
	}
	fmt.Println(f)
}