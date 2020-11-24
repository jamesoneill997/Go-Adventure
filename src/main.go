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

func parseStory(arcNames []string, jsonArcs map[string]interface{}) *templates.Story {
	//initialise story and arcs, max size is 25 arcs
	story := new(templates.Story)
	arcs := make([]templates.Arc, 25)

	//extract dynamic arc names from json
	for k := range jsonArcs {
		arcNames = append(arcNames, k)
	}

	//for each arc, initialise an arc type and use the map to get nested values
	for i, val := range arcNames {
		//construct a json object from map at arc name
		jsonMap, _ := json.Marshal(jsonArcs[val])

		//handle a file that is too large
		if len(arcs) > i {
			json.Unmarshal(jsonMap, &arcs[i])
		} else {
			log.Fatal("Your story is too large")
		}

		//trim array to minimum required size
		if i == len(arcNames)-1 {
			arcs = arcs[:i-1]
		}
	}

	//write list of arcs to story
	story.Arcs = arcs

	return story
}

func main() {
	//initialise data structs to assist with parsing
	arcNames := []string{}
	jsonArcs := make(map[string]interface{})

	//handle input file
	file := fileReader("./gopher.json")
	json.Unmarshal(file, &jsonArcs)

	//parse data to a story object
	story := parseStory(arcNames, jsonArcs)
	fmt.Println(story.Arcs)
}
