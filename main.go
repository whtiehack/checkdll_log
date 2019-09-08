package main

import (
	"log"
	"syscall"
)

func main() {
	log.Print("main start")
	h, err := syscall.LoadLibrary("dll/test.dll")
	if err != nil {
		log.Panic("load test.dll error")
	}
	proc, err := syscall.GetProcAddress(h, "Showui")
	if err != nil {
		log.Panic("GetProcAddress Showui error")
	}
	_, _, err = syscall.Syscall(proc, 0, 0, 0, 0)
	if err != nil {
		log.Panic("Syscall Showui error")
	}
	log.Print("main end")
	//	print("main end")
	//	fmt.Printf("main end")
}
