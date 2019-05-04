package model

type News struct {
	Base
	TypeId  `json:"type_id"`
	Title   `json:"title"`
	Content `json:"content"`
	Read    `json:"read"`
	Image   `json:"image"`
}

type NewsType struct {
	Base
	Title `json:"title"`
	Icon  `json:"icon"`
	Desc  `json:"desc"`
}
