package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("no argument")
		return
	}

	zipName := "master.zip"

	err := getProjectTemplate(zipName)
	check(err)

	dest := os.Args[1]
	err = extract(zipName, dest)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(nil)
	}
}
