# Go nil map access panic

This repository demonstrates a common error in Go: accessing an uninitialized map.  Attempting to read or write to a nil map will cause a runtime panic.

The `bug.go` file contains code that reproduces the issue. The `bugSolution.go` demonstrates how to safely handle nil maps. 

## Solution

The best way to prevent a panic is to always check if the map is initialized before accessing it. You can use the `make()` function to initialize the map to an empty map if needed.