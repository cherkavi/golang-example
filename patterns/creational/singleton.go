package creational

import (
	"sync"
)

var singletonChecker sync.Once
var instance NameAware

func GetSingletonInstance() NameAware {
	singletonChecker.Do(func() {
		instance = animal{"Kroko"}
	})
	return instance
}
