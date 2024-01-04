package domain

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}
