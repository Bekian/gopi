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
	envMap := LoadEnv()
	fmt.Printf("appID = %s\n", envMap["APP_ID"])
	fmt.Printf("token = %s\n", envMap["DISCORD_TOKEN"])
	fmt.Printf("pub key = %s\n", envMap["PUBLIC_KEY"])
}
