#include <stdlib.h>
#include <stdio.h>

#ifdef _WIN32
#include <Windows.h>
#define dlsym GetProcAddress
#else
#include <dlfcn.h>
#include <link.h>
#endif

typedef void *(*luaL_newstate)();
void *luaL_newstate_wrap(void *f)
{
    return ((luaL_newstate)f)();
}

typedef void (*luaL_openlibs)(void *L);
void luaL_openlibs_wrap(void *f, void *L)
{
    ((luaL_openlibs)f)(L);
}

typedef int (*luaL_loadstring)(void *L, const char *s);
int luaL_loadstring_wrap(void *f, void *L, const char *s)
{
    return ((luaL_loadstring)f)(L, s);
}

typedef int (*luaL_loadbuffer)(void *L, const char *buff, size_t sz, const char *name);
int luaL_loadbuffer_wrap(void *f, void *L, const char *buff, size_t sz, const char *name)
{
    return ((luaL_loadbuffer)f)(L, buff, sz, name);
}

typedef void (*lua_pushlstring)(void *L, const char *s, size_t len);
void lua_pushlstring_wrap(void *f, void *L, const char *s, size_t len)
{
    ((lua_pushlstring)f)(L, s, len);
}

typedef const char *(*lua_tolstring)(void *L, int index, size_t *len);
const char *lua_tolstring_wrap(void *f, void *L, int index, size_t *len)
{
    return ((lua_tolstring)f)(L, index, len);
}

typedef double (*lua_tonumber)(void *L, int index);
double lua_tonumber_wrap(void *f, void *L, int index)
{
    return ((lua_tonumber)f)(L, index);
}

typedef int (*lua_toboolean)(void *L, int index);
int lua_toboolean_wrap(void *f, void *L, int index)
{
    return ((lua_toboolean)f)(L, index);
}

typedef int (*lua_gettop)(void *L);
int lua_gettop_wrap(void *f, void *L)
{
    return ((lua_gettop)f)(L);
}

typedef int (*lua_type)(void *L, int index);
int lua_type_wrap(void *f, void *L, int index)
{
    return ((lua_type)f)(L, index);
}

typedef const char *(*lua_typename)(void *L, int tp);
const char *lua_typename_wrap(void *f, void *L, int tp)
{
    return ((lua_typename)f)(L, tp);
}

typedef void (*lua_pushcclosure)(void *L, int *fn, int n);
void lua_pushcclosure_wrap(void *f, void *L, int *fn, int n)
{
    ((lua_pushcclosure)f)(L, fn, n);
}

typedef int (*lua_pcall)(void *L, int nargs, int nresults, int errfunc);
int lua_pcall_wrap(void *f, void *L, int nargs, int nresults, int errfunc)
{
    return ((lua_pcall)f)(L, nargs, nresults, errfunc);
}

typedef void (*lua_setfield)(void *L, int index, const char *k);
void lua_setfield_wrap(void *f, void *L, int index, const char *k)
{
    ((lua_setfield)f)(L, index, k);
}

typedef void (*lua_call)(void *L, int nargs, int nresults);
void lua_call_wrap(void *f, void *L, int nargs, int nresults)
{
    ((lua_call)f)(L, nargs, nresults);
}