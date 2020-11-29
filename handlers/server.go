package handlers

import (
	"adventure/templates"
	"fmt"
	"net/http"
)

type story templates.Story

func (s *story) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(s.Arcs)
}
