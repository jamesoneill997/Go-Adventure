package main

import (
	"adventure/templates"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var htmlTemplate string

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
	story.Arcs = make(map[string]templates.Arc)
	for i := range arcs {
		story.Arcs[arcNames[i]] = arcs[i]
	}

	return story
}

//HTTP hanlder setup
//handler type
type handler struct {
	story *templates.Story
}

//returns httphandler
func myHandler(st *templates.Story) http.Handler {
	return handler{story: st}
}

//ServeHTTP function
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tplate := template.Must(template.New("").Parse(htmlTemplate))

	err := tplate.Execute(w, h.story.Arcs["intro"])

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func main() {
	port := 8080
	pageTemplate, err := ioutil.ReadFile("./ui/template.html")
	if err != nil {
		fmt.Println("Error reading html template")
		log.Fatal(err)
	}
	htmlTemplate = string(pageTemplate)
	//initialise data structs to assist with parsing
	arcNames := []string{}
	jsonArcs := make(map[string]interface{})

	//handle input file
	file := fileReader("./gopher.json")
	json.Unmarshal(file, &jsonArcs)

	//parse data to a story object
	storyPtr := parseStory(arcNames, jsonArcs)

	h := myHandler(storyPtr)
	fmt.Printf("Serving CYOA on %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), h))
}
