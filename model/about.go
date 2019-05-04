package model

type About struct {
	Base

	Title   `json:"title"`
	Content `json: "content"`
	Image   `json:"image"`
}
