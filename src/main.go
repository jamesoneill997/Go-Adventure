package main

import (
	"adventure/templates"
	"encoding/json"
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

func main() {
	arcNames := []string{}

	jsonArcs := make(map[string]interface{})

	arcs := make([]templates.Arc, 25)

	file := fileReader("./gopher.json")

	json.Unmarshal(file, &jsonArcs)

	for k := range jsonArcs {
		arcNames = append(arcNames, k)
	}

	for i, val := range arcNames {
		jsonMap, _ := json.Marshal(jsonArcs[val])
		if len(arcs) > i {
			json.Unmarshal(jsonMap, &arcs[i])
		} else {
			log.Fatal("Your story is too large")
		}

		//trim array to minimum required size
		if i == len(arcNames)-1 {
			arcs = arcs[:i]
		}
	}

	fmt.Println(arcs)
}
