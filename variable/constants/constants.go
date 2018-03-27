package constants

var A int = 10
var B string = "value from another package"
var k string = "not visible variable"

const PI float32 = 3.1412

const (
	Message      = "this is const message"
	ErrorMessage = "this is error message "
	LongMessage  = ` this
	is multi-line
	message 
	`
)
