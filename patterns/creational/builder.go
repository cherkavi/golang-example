package creational

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Surname string
	Age     int
	Male    bool
}

func (p Person) String() string {
	return fmt.Sprintf("name: %v, surname: %v, age: %v, male: %v", p.Name, p.Surname, p.Age, p.Male)
}

type PersonBuilder struct {
	data map[string]interface{}
}

func (p *PersonBuilder) New() *PersonBuilder {
	p.data = make(map[string]interface{})
	return p
}
func (p *PersonBuilder) Set(fieldName string, value interface{}) *PersonBuilder {
	p.data[fieldName] = value
	return p
}

func (p *PersonBuilder) Build() Person {
	returnValue := Person{}
	valueOf := reflect.ValueOf(&returnValue)
	for key, value := range p.data {
		field := valueOf.Elem().FieldByName(key)
		// fmt.Printf("setValue: %#v to field: %#v  ( with name: %v )\n", value, field, key)
		field.Set(reflect.ValueOf(value))
	}
	return returnValue
}

func Builder() {
	builder := (&PersonBuilder{}).New()
	builder.Set("Name", "Stepan")
	builder.Set("Surname", "Razin")
	builder.Set("Age", 33)
	builder.Set("Male", true)
	fmt.Println(builder.Build())
}
