package main

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -L. -lunhook_ntdll
#include <stdlib.h>
extern int unhook_ntdll();
*/
import "C"

import (
    "fmt"
    "os"
)

func main() {
    ret := C.unhook_ntdll()
    if ret != 0 {
        fmt.Println("Failed to unhook ntdll.dll")
        os.Exit(1)
    }
    fmt.Println("ntdll.dll unhooked successfully")
}
