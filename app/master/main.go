package main

import (
	"log"
	"github.com/Yggdrasill/app/master/cmd"
)


func main() {

	version := cmd.GetVersion()
	log.Printf("Received:", version)
}
