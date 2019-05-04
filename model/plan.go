package model

type Plan struct {
	Base
	TypeId    `json:"type_id"`
	Title     `json:"title"`
	Image     `json:"image"`
	MultiView `json:"multi_view"`
	Content   `json:"content"`
}

type PlanType struct {
	Base
	Name `json:"name"`
	Icon `json:"icon"`
	Desc `json:"desc"`
}
