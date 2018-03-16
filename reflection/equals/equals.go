package equals

import (
	"fmt"
	"reflect"
)

type Movie struct{
	Title string
	LengthInMinutes int
}

func (movie Movie) String() string{
	return fmt.Sprintf("['%v' %vmin]", movie.Title, movie.LengthInMinutes)
}


func Equals(){
	var firstMovie = Movie{"Shine", 95}
	var secondMovie = Movie{"Shine", 95}
	// anonymous struct
	var thirdMovie = struct{Title string;LengthInMinutes int}{"Shine", 95}
	fmt.Printf("movie %v is equals to %v ? %v \n", firstMovie, secondMovie, reflect.DeepEqual(firstMovie, secondMovie))
	fmt.Printf("movie %v is equals to %v ? %v \n", firstMovie, thirdMovie, reflect.DeepEqual(firstMovie, thirdMovie))
}