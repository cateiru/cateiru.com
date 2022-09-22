package main

import (
	"github.com/cateiru/cateiru.com/src"
)

// Set this variable at build time.
//
// Example:
//
//	go build  -ldflags="-X main.mode=prod"
var mode string = ""

func init() {
	src.Init(mode)
}

func main() {
	// Run backend server
	src.Server()
}
