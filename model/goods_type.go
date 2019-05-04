package model

/*
import (
	"github.com/jinzhu/gorm"
	"time"
)
*/
/**
func (this *GoodsType) GetAll(pageNum int, pageSize int) ([]GoodsType, error) {

	var goodType []GoodsType

	// if pageSize > 0 && pageNum > 0 {
	// 	db = db.Offset(pageNum).Limit(pageSize)
	// }
	// err := db.Select("id, app_id, redirect_uri, create_time").Where(Authorize{AppId: AppId}).First(&authorize).Error
	// err := db.Where(maps).Find(&goods).Error
	err := db.Find(&goodType).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return goodType, nil

}

*/
