package models

type Note struct {
	Model
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Path    []string `json:"path"`
	Author  string   `json:"author"`
	Public  bool     `json:"public"`
	Type    string   `json:"type"`
	Tags    []string `json:"tags"`
}
