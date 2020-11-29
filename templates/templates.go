package templates

//Story struct contains list of arcs
type Story struct {
	Arcs map[string]Arc `json:"arcs"`
}

//An Arc type may store a chapter or a character
type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

//Option struct details options and lists next arc
type Option struct {
	Text    string `json:"text"`
	NextArc string `json:"arc"`
}
