package main

import "C"
import (
	"log"
	"syscall"
	"unsafe"
)

var count = 0

var user32 syscall.Handle
var dialogBoxParam uintptr
var endDialog uintptr

func init() {
	count = 10
	log.Println("dll init", count)
	user32, _ = syscall.LoadLibrary("user32.dll")
	dialogBoxParam, _ = syscall.GetProcAddress(user32, "DialogBoxParamW")
	endDialog, _ = syscall.GetProcAddress(user32, "EndDialog")

}

func main() {
	// Need a main function to make CGO compile package as C shared library

}

//export Showui
func Showui() int {
	log.Println("show uid called")

	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	getModuleHandle, _ := syscall.GetProcAddress(kernel32, "GetModuleHandleW")
	dllModule, _, _ := syscall.Syscall(getModuleHandle, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("test.dll"))), 0, 0)
	dlg := &Dialog{id: 100}
	syscall.Syscall6(dialogBoxParam, 5, dllModule, 100, 0, syscall.NewCallback(dlg.dialogWndProc), 0, 0)
	log.Println("show ui finish")
	return 1
}

type Dialog struct {
	id int
}

func (dlg *Dialog) dialogWndProc(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	log.Println("dlg??", dlg, dlg.id)
	switch msg {
	// WM_INITDIALOG             = 272
	case 272:
		log.Println("dialog init")
		return 1
	// WM_CLOSE                  = 16
	case 16:
		log.Println("dialog close")
		// enddialog
		syscall.Syscall(endDialog, 2, hwnd, 0, 0)
		return 0
	}
	return 0
}
