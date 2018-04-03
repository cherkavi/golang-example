package yaml

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Person struct {
	ID      string   `yaml:"id,omitempty"`
	Name    string   `yaml:"name,omitempty"`
	Address *Address `yaml:"address,omitempty"`
	Note    string   `yaml:"-"` // transient field
}
type Address struct {
	City       string `yaml:"city,omitempty"`
	PostalCode string `yaml:"postal-code,omitempty"`
}

func (p Person) String() string {
	return fmt.Sprintf("ID:%v Name:%v Note(transient): %v  Address: %v %v", p.ID, p.Name, p.Note, p.Address.City, p.Address.PostalCode)
}

func Example() {
	file, _ := ioutil.TempFile("", "temp file")
	defer func() {
		_ = os.Remove(file.Name())
	}()

	person := Person{ID: "1", Name: "Jack", Note: "original person", Address: &Address{City: "London", PostalCode: "22150"}}
	fmt.Printf("> original data: %v\n", person)

	data, _ := yaml.Marshal(person)
	// json.NewEncoder
	fmt.Println("> write data into file ")
	ioutil.WriteFile(file.Name(), data, 666)

	fmt.Println("> read data from file ")
	text, _ := ioutil.ReadFile(file.Name())
	fmt.Println(string(text))

	var restoredPerson Person
	yaml.Unmarshal(text, &restoredPerson)
	fmt.Printf("> restored data: %v \n", restoredPerson)
}
