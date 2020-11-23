package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func fileReader(path string) []byte {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func jsonParser(data []byte)

func main() {
	fmt.Println(fileReader("gopher.json"))
}
