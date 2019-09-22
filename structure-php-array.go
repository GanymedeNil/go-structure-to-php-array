package go_structure_to_php_array

import (
	"fmt"
	"reflect"
)

const (
	arrayPrefix = "["
	arraySuffix = "]"
	arrayLink   = "=>"
	arrayEnding = ","
)

var deep = 0

func tab() (str string) {
	str = ""
	for i := 0; i < deep; i++ {
		str += "   "
	}
	return

}

func StructTOPhpArray(v interface{}) string {
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	switch t.Kind() {
	case reflect.Map:
		return isMap(value)
	case reflect.Array, reflect.Slice:
		return isSlice(value)
	case reflect.Struct:
		return isStruct(t, value)
	case reflect.String:
		return isString(value)
	default:
		return isOther(value)
	}
}

func isMap(v reflect.Value) string {
	deep++
	str := arrayPrefix + "\n"
	keys := v.MapKeys()
	for _, key := range keys {
		value := v.MapIndex(key)
		str += fmt.Sprintf("%s%v %s %v%s\n",
			tab(), StructTOPhpArray(key.Interface()), arrayLink, StructTOPhpArray(value.Interface()), arrayEnding)
	}
	deep--
	str += tab() + arraySuffix
	return str
}

func isSlice(v reflect.Value) string {
	deep++
	str := arrayPrefix + "\n"
	for i := 0; i < v.Len(); i++ {
		str += fmt.Sprintf("%s%v%s\n",
			tab(), StructTOPhpArray(v.Index(i).Interface()), arrayEnding)
	}
	deep--
	str += tab() + arraySuffix
	return str
}

func isStruct(t reflect.Type, v reflect.Value) string {
	deep++
	str := arrayPrefix + "\n"
	for i := 0; i < v.NumField(); i++ {
		key := t.Field(i).Tag.Get("php")
		str += fmt.Sprintf("%s'%s' %s %v%s\n",
			tab(), key, arrayLink, StructTOPhpArray(v.Field(i).Interface()), arrayEnding)
	}
	deep--
	str += tab() + arraySuffix
	return str
}

func isString(v reflect.Value) string {
	return fmt.Sprintf("'%s'", v.String())
}

func isOther(v reflect.Value) string {
	return fmt.Sprintf("%v", v.Interface())
}
