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
	fmt.Println(story)
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

	for _, st := range h.story.Arcs {
		fmt.Println(st)
	}

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

/*
Serving CYOA on 8080
2020/11/29 22:10:47 http: panic serving [::1]:51692: template: :10: unexpected "}" in operand
goroutine 6 [running]:
net/http.(*conn).serve.func1(0xc00010efa0)
        /usr/local/go/src/net/http/server.go:1801 +0x147
panic(0x74ea80, 0xc000013a30)
        /usr/local/go/src/runtime/panic.go:975 +0x47a
html/template.Must(...)
        /usr/local/go/src/html/template/template.go:372
main.handler.ServeHTTP(0xc00000e620, 0x81da60, 0xc00017e0e0, 0xc000196000)
        /home/james/Desktop/github/go-adventure/src/main.go:72 +0x1f9
net/http.serverHandler.ServeHTTP(0xc00017e000, 0x81da60, 0xc00017e0e0, 0xc000196000)
        /usr/local/go/src/net/http/server.go:2843 +0xa3
net/http.(*conn).serve(0xc00010efa0, 0x81e3a0, 0xc000066740)
        /usr/local/go/src/net/http/server.go:1925 +0x8ad
created by net/http.(*Server).Serve
        /usr/local/go/src/net/http/server.go:2969 +0x36c
2020/11/29 22:10:47 http: panic serving [::1]:51694: template: :10: unexpected "}" in operand
goroutine 7 [running]:
net/http.(*conn).serve.func1(0xc00010f040)
        /usr/local/go/src/net/http/server.go:1801 +0x147
panic(0x74ea80, 0xc000013b10)
        /usr/local/go/src/runtime/panic.go:975 +0x47a
html/template.Must(...)
        /usr/local/go/src/html/template/template.go:372
main.handler.ServeHTTP(0xc00000e620, 0x81da60, 0xc00017e1c0, 0xc000196100)
        /home/james/Desktop/github/go-adventure/src/main.go:72 +0x1f9
net/http.serverHandler.ServeHTTP(0xc00017e000, 0x81da60, 0xc00017e1c0, 0xc000196100)
        /usr/local/go/src/net/http/server.go:2843 +0xa3
net/http.(*conn).serve(0xc00010f040, 0x81e3a0, 0xc000066840)
        /usr/local/go/src/net/http/server.go:1925 +0x8ad
created by net/http.(*Server).Serve
        /usr/local/go/src/net/http/server.go:2969 +0x36c
2020/11/29 22:10:47 http: panic serving [::1]:51696: template: :10: unexpected "}" in operand
goroutine 12 [running]:
net/http.(*conn).serve.func1(0xc00010f0e0)
        /usr/local/go/src/net/http/server.go:1801 +0x147
panic(0x74ea80, 0xc0000940d0)
        /usr/local/go/src/runtime/panic.go:975 +0x47a
html/template.Must(...)
        /usr/local/go/src/html/template/template.go:372
main.handler.ServeHTTP(0xc00000e620, 0x81da60, 0xc0000a8000, 0xc000096000)
        /home/james/Desktop/github/go-adventure/src/main.go:72 +0x1f9
net/http.serverHandler.ServeHTTP(0xc00017e000, 0x81da60, 0xc0000a8000, 0xc000096000)
        /usr/local/go/src/net/http/server.go:2843 +0xa3
net/http.(*conn).serve(0xc00010f0e0, 0x81e3a0, 0xc000092000)
        /usr/local/go/src/net/http/server.go:1925 +0x8ad
created by net/http.(*Server).Serve
        /usr/local/go/src/net/http/server.go:2969 +0x36c

*/
