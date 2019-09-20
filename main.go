package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("not ok")
	}
	fmt.Println(bi.Main.Version)
}