package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	// Writeはbyte[]を受け取るので、Stringを変換して渡す
	file.Write([]byte("os.File example\n"))
	file.Close()
}
