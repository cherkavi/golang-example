package main

import (
	"fmt"

	"./behavior"
	"./creational"
)

func main() {
	fmt.Println(" Creational: Factory method ")
	creational.FactoryMethod()

	val := creational.GetSingletonInstance()
	fmt.Printf("Singleton: %#v     %p\n", creational.GetSingletonInstance(), &val)
	val = creational.GetSingletonInstance()
	fmt.Printf("Singleton: %#v     %p\n", creational.GetSingletonInstance(), &val)
	val = creational.GetSingletonInstance()
	fmt.Printf("Singleton: %#v     %p\n", creational.GetSingletonInstance(), &val)
	val = creational.GetSingletonInstance()

	fmt.Println("Creational: Builder")
	creational.Builder()

	fmt.Println("Behavior: ChainOfResponsibilities")
	behavior.ChaineOfResponsibilities()
}
