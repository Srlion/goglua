package main

// usage: go build -buildmode=c-shared

/*
#cgo LDFLAGS: "-Wl,--version-script=script.exp"

int gmod13_open(void* L);
int gmod13_close(void* L);

int test(void* L);
*/
import "C"

import (
	"fmt"
	// "github.com/Srlion/goglua"
	"unsafe"
)

// go complains if we use glua.State because it can't export it for w/e reason
type State = unsafe.Pointer

//export test
func test(L State) C.int {
	fmt.Printf("Current stack size: %v\n", glua.GetTop(L))
	return 0
}

//export gmod13_open
func gmod13_open(L State) C.int {
	glua.PushFunc(L, C.test)
	glua.SetGlobal(L, "test")
	return 0
}

//export gmod13_close
func gmod13_close(L State) C.int {
	return 0
}

// required by go for w/e reason
func main() {}
