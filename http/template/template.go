package template

import (
	"strings"
	"text/template" // html/template
)

func ExampleOfTemplate() string {
	var data = map[string]string{"key1": "one", "key2": "two"}
	t1 := template.New("nameOfTemplate")
	t1.Parse("value from map: {{.key1}}")

	writer := &strings.Builder{}
	t1.Execute(writer, data) // os.Stdout
	return writer.String()
}
