package copy

import (
	"reflect"
	"fmt"
	"github.com/jinzhu/copier" // just execute "go get <url>"
)


type Movie struct{
	Title string `custom_name:"value1"  another_tag:"value number 2"`
	LengthInMinutes int
}


func ReflectCopy(){
	var one = []string{"z","b","k"}
	var two = []string{"a","b","c","d", "e", "f"}
	fmt.Printf("copy %#v into %#v\n", one, two)
	reflect.Copy(reflect.ValueOf(two), reflect.ValueOf(one))
	fmt.Printf("copy result %#v \n", two)

	var movie1 = Movie{"Alien", 109}
	var movie2 Movie
	copier.Copy(&movie2, &movie1)
	fmt.Printf("first: %#v    second after copy: %#v \n", movie1, movie2)

	fmt.Println()
}