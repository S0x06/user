package model

type Culture struct {
	Base
	Title     `json:"title"`
	Brief     `json:"brief"`
	Image     `json:"image"`
	MultiView `json:"multi_view"`
	Link      `json:"link"`
	Icon      `json:"icon"`
}
