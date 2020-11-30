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

	}
	arcs = arcs[:len(arcNames)]
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
	//Parse path
	path := r.URL.Path
	//if path not set, start at intro
	if path == "" || path == "/" {
		path = "/intro"
	}
	//remove leading '/'
	path = path[1:]
	tplate := template.Must(template.New("").Parse(htmlTemplate))
	fmt.Println(h.story.Arcs)
	if arc, ok := h.story.Arcs[path]; ok {
		err := tplate.Execute(w, arc)

		if err != nil {
			//error for me
			log.Printf("%v", err)

			//error for user
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

		}

		return
	}

	http.Error(w, "Chapter not found", http.StatusNotFound)

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
