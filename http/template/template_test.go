package template

import "testing"

func TestExampleOfTemplate(t *testing.T) {
	returnValue := ExampleOfTemplate()
	expectedValue := "value from map: one"

	if returnValue != expectedValue {
		t.Errorf("expected value is: %v", expectedValue)
	}
}
