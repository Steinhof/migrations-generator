package main

import "fmt"

func main() {
	version := getLatestVersion()
	createFile(version)
	fmt.Print(version)
}
