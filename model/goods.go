package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Goods struct {
	Base

	TypeId int `json:"type_id" gorm:"index"`
	// GoodsType GoodsType `json:"goods_type"`
	Name      string `json:"name"`
	Brief     string `json:"brief"`
	Recommend int    `json:"recommend"`
	Hot       int    `json:"hot"`
}

type GoodsType struct {
	Base
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Background string `json:"background"`
	Brief      string `json:"brief"`
}

/*
func (this *Goods) GetAll(pageNum int, pageSize int) ([]Goods, error) {

	var goods []Goods

	// if pageSize > 0 && pageNum > 0 {
	// 	db = db.Offset(pageNum).Limit(pageSize)
	// }
	// err := db.Select("id, app_id, redirect_uri, create_time").Where(Authorize{AppId: AppId}).First(&authorize).Error
	// err := db.Where(maps).Find(&goods).Error
	err := db.Find(&goods).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return goods, nil

}

func (this *Goods) GetOneGoods() {

}

func (this *Goods) ModifyGoods() {

}

func (this *Goods) DeleteGoods() {

}
*/
