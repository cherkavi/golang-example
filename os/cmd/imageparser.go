package recognizer

import (
	"fmt"
	exec "os/exec"
	"strings"
)

func executeCommand(commands ...string) (string, error) {
	//if output, err := exec.Command("bash", "-c", "/usr/bin/alpr", "--help").Output(); err != nil {
	fullList := append([]string{}, "-c")
	fullList = append(fullList, commands...)

	if output, err := exec.Command("bash", fullList...).Output(); err != nil {
		return "", err
	} else {
		return string(output), nil
	}
}

func parseResult(data string) string {
	if strings.Index(data, "No license") >= 0 {
		return ""
	}
	return strings.TrimSpace(data[strings.Index(data, "-")+1 : strings.Index(data, "confidence")])
}

func RecognizeImage(path string) (string, error) {
	// if outputData, err := executeCommand(fmt.Sprintf("alpr --country eu %v ", path)); err != nil {
	// cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	// 	if outputData, err := executeCommand(fmt.Sprintf("/usr/bin/alpr --help; echo 1>&2 stderr")); err != nil {
	if outputData, err := executeCommand(fmt.Sprintf("/usr/bin/alpr --country eu %v", path)); err != nil {
		return "", err
	} else {
		return parseResult(outputData), nil
	}
	// No license plates found.
}
