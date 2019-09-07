package main

import (
	"log"
	"syscall"
)

func main() {
	log.Print("main start")
	h, _ := syscall.LoadLibrary("test.dll")
	proc, _ := syscall.GetProcAddress(h, "Showui")
	_, _, _ = syscall.Syscall(proc, 0, 0, 0, 0)

	log.Print("main end")
	//	print("main end")
	//	fmt.Printf("main end")
}
