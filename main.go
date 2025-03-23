package main

import (
	"fmt"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("starting")
	// load the env map variables
	envMap := LoadEnv()

}
