package main

import "github.com/cateiru/cateiru.com/src"

func init() {
	src.Init()
}

func main() {
	// Run backend server
	src.Server()
}
