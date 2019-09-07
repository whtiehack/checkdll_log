package main

import "C"
import "log"

func init() {
	//	log.Println("dll init")
}

func main() {
	// Need a main function to make CGO compile package as C shared library

}

//export Showui
func Showui() int {
	// use print will fine.
	print("show uid called\n")
	// use log will panic .  If comment this line,  there is a probability of multiple runs.
	log.Println("show ui called")
	return 1
}
