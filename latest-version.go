package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func getLatestVersion() int {
	files, err := ioutil.ReadDir("./flyway/")
	if err != nil {
		log.Fatal(err)
	}

	maxVersion := 0
	for _, file := range files {
		fileName := file.Name()
		version := strings.Split(fileName, "__")
		number := strings.Split(version[0], "V")

		if len(number) != 1 {
			x, err := strconv.ParseInt(number[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			maxVersion = int(math.Max(float64(int(x)), float64(maxVersion)))
		}
	}

	return maxVersion
}
