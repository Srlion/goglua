# goglua

# Archive Reason
Using Go with the Lua C API is really difficult because Go and Lua handle errors differently. Lua uses something called longjmp to handle errors in C, but this doesn't work well with Go and can cause serious problems. There's no easy way to stop these longjmp calls from messing with Go. So, it's better not to try combining them.

## Install

`go get -u github.com/Srlion/goglua`
