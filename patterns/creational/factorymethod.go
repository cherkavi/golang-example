package creational

import "fmt"

type NameAware interface {
	GetName() string
}

type human struct {
	fullname string
}

func (h human) GetName() string {
	return h.fullname
}

type animal struct {
	shortname string
}

func (a animal) GetName() string {
	return a.shortname
}

type subject struct {
	material string
}

func (s subject) GetName() string {
	return s.material
}

func createObject(choice int, value string) NameAware {
	switch choice {
	case 1:
		{
			return human{value}
		}
	case 2:
		{
			return animal{value}
		}
	default:
		{
			return subject{value}
		}
	}
}

func FactoryMethod() {
	fmt.Printf("%#v \n", createObject(1, "Petya"))
	fmt.Printf("%#v \n", createObject(2, "Murka"))
	fmt.Printf("%#v \n", createObject(3, "Water"))
}
