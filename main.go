package main

import (
	"fmt"
	"os"
)

const DF_NAME string = "data.op"

var logged_in bool

func main() {
	_, err := os.Open(DF_NAME)
	if err != nil {
		register()
	} else {
		o := login()
	}

	MainMenu()

}

func MainMenu() {}
