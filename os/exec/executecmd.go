package exec

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	// import "github.com/toqueteos/webbrowser"
)

// webbrowser.Open("http://golang.org")
func Openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// execute "command" with "parameters"
// output to current output
func ExecuteCommandAndOutput(command string, parameters ...string) {
	cmd := exec.Command(command, parameters...)
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println()
}
