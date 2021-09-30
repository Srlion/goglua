package glua

/*
#cgo LDFLAGS: -ldl
#include "glua.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

// free the c char when gc cycle happens
type CString struct {
	c *C.char
}

func (cstr *CString) free() {
	C.free(unsafe.Pointer(cstr.c))
}

func CStr(gs string) *CString {
	cstr := &CString{C.CString(gs)}
	runtime.SetFinalizer(cstr, (*CString).free)
	return cstr
}

type State = unsafe.Pointer

var handler unsafe.Pointer = func() unsafe.Pointer {
	var handler unsafe.Pointer = nil
	var paths []string

	if runtime.GOOS == "linux" && runtime.GOARCH == "386" {
		paths = []string{"lua_shared_srv.so", "garrysmod/bin/lua_shared_srv.so", "bin/linux32/lua_shared.so"}
	} else if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
		paths = []string{"lua_shared.so", "bin/linux64/lua_shared.so"}
	} else if runtime.GOOS == "windows" && runtime.GOARCH == "386" {
		paths = []string{"lua_shared.dll", "garrysmod/bin/lua_shared.dll", "bin/lua_shared.dll"}
	} else if runtime.GOOS == "windows" && runtime.GOARCH == "amd64" {
		paths = []string{"lua_shared.dll", "bin/win64/lua_shared.dll"}
	}

	for i := 0; i < len(paths); i++ {
		if handler == nil {
			handler = C.dlopen(CStr(paths[i]).c, C.RTLD_LAZY)
		} else {
			break
		}
	}

	if handler == nil {
		panic(fmt.Sprintf("%s", C.GoString(C.dlerror())))
	}

	return handler
}()
var luaL_newstate = C.dlsym(handler, CStr("luaL_newstate").c)
var luaL_openlibs = C.dlsym(handler, CStr("luaL_openlibs").c)
var luaL_loadstring = C.dlsym(handler, CStr("luaL_loadstring").c)
var lua_pushlstring = C.dlsym(handler, CStr("lua_pushlstring").c)
var lua_tolstring = C.dlsym(handler, CStr("lua_tolstring").c)
var lua_gettop = C.dlsym(handler, CStr("lua_gettop").c)

var lua_pushcclosure = C.dlsym(handler, CStr("lua_pushcclosure").c)
var lua_pcall = C.dlsym(handler, CStr("lua_pcall").c)
var lua_call = C.dlsym(handler, CStr("lua_call").c)
var lua_setfield = C.dlsym(handler, CStr("lua_setfield").c)

const LUA_GLOBALSINDEX = C.int(-10002)

func NewState() State {
	state := C.luaL_newstate_wrap(luaL_newstate)
	return state
}

func OpenLibs(L State) {
	C.luaL_openlibs_wrap(luaL_openlibs, L)
}

func LoadString(L State, str string) error {
	if lua_error_code := C.luaL_loadstring_wrap(luaL_loadstring, L, CStr(str).c); lua_error_code != 0 {
		return errors.New(GetErrorString(L))
	}

	return nil
}

func GetTop(L State) int {
	return int(C.lua_gettop_wrap(lua_gettop, L))
}

func PushString(L State, str string) {
	C.lua_pushlstring_wrap(lua_pushlstring, L, CStr(str).c, C.size_t(len(str)))
}

func PushFunc(L State, f unsafe.Pointer) {
	C.lua_pushcclosure_wrap(lua_pushcclosure, L, (*C.int)(f), 0)
}

func ToLString(L State, index int) string {
	return C.GoString(C.lua_tolstring_wrap(lua_tolstring, L, C.int(index), nil))
}

func PCall(L State, nargs, nresults, errfunc int) int {
	return int(C.lua_pcall_wrap(lua_pcall, L, C.int(nargs), C.int(nresults), C.int(errfunc)))
}

func Call(L State, nargs int, nresults int) {
	C.lua_call_wrap(lua_call, L, C.int(nargs), C.int(nresults))
}

func SetGlobal(L State, name string) {
	C.lua_setfield_wrap(lua_setfield, L, LUA_GLOBALSINDEX, CStr(name).c)
}

func GetErrorString(L State) string {
	return ToLString(L, -1)
}
