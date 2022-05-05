package models

type Note struct {
	Model
	Title   string
	Content string
	Path    []string
	Author  string
	Public  bool
	Type    string
	Tags    []string
}
