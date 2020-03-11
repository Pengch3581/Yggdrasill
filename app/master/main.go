package main

import (
	"log"
	"master/cmd"
)


func main() {

	version := cmd.GetVersion()
	log.Printf("Received:", version)
}
