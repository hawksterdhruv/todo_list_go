package commons

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// -------------------- SCREEN FUNCTIONS -----------------------
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// -------------------- INPUT FUNCTIONS -----------------------
func Input(prompt string) (string) {
	// Utility function for CLI but can be kept in commons
	fmt.Printf("%s", prompt)
	in := bufio.NewReader(os.Stdin)
	result, _, _ := in.ReadLine()
	return string(result)
}