package constants

var A int = 10
var B string = "value from another package"
var k string = "not visible variable"

const PI float32 = 3.1412

const (
	Message        = "this is const message"
	EmptyConstant1 // example of the empty constant - value will be picked up from value above
	EmptyConstant2 // example of the empty constant - value will be picked up from value above
	ErrorMessage   = "this is error message "
	LongMessage    = ` this
	is multi-line
	message 
	`
)

const (
	// example of the sequence of constants
	Zero = 1 * iota
	One
	Two // some value can be ignored using underscore
	Three
	Four
	Five
	Six
	Seven
)
