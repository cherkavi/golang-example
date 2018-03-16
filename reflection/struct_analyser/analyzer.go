package struct_analyser

import (
	"reflect"
	"fmt"
)

func ReflectInfo(target interface{}){
	value := reflect.ValueOf(target).Elem()

	for i := 0; i < value.NumField(); i++ {
		valueField := value.Field(i)
		typeField := value.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Whole tag: %#v \t Tag Value: %v\t Package path: %v   \n",
			typeField.Name, valueField.Interface(), tag, tag.Get("custom_name"), typeField.Type)
	}
}
