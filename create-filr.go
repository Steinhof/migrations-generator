package main

import (
	"fmt"
	"os"
	"strconv"
)

func createFile(fileVersion int) {
	version := strconv.Itoa(fileVersion)
	err := os.WriteFile("V"+version+".sql", []byte("Hello"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}
