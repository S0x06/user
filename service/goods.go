package service

import (
	"github.com/gin-gonic/gin"
	. "github.com/official/models"
	. "github.com/official/utils"
	// "net/http"
)

//@Summary 产品服务页接口
//@Produce  json
//@Param type_id query int true "TypeId"
//@Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//@Router /api/v1/services [get]
func MultiGoods(c *gin.Context) {
	var version string = "1.0.0"
	var href string = "http://example.org/services"

	//获取类型
	var gt *GoodsType = new(GoodsType)
	goodsType, err := gt.GetAll(10, 1)
	if err != nil {
		var Error map[string]interface{} = map[string]interface{}{
			"title":   "获取产品类型",
			"code":    "40",
			"message": err,
		}
		c.JSON(200, ErrorMessage(version, href, Error))
	}

	var g *Goods = new(Goods)
	goods, err := g.GetAll(10, 1)
	if err != nil {

		var Error map[string]interface{} = map[string]interface{}{
			"title":   "获取产品",
			"code":    "40",
			"message": err,
		}

		c.JSON(200, ErrorMessage(version, href, Error))
	}

	var links map[string]interface{} = map[string]interface{}{
		"rel":  "feed",
		"href": "http://example.org/friends/rss",
	}

	var goodsTypeData map[string]interface{} = map[string]interface{}{
		"href":  "test",
		"data":  goodsType,
		"links": "links",
	}

	var goodsData map[string]interface{} = map[string]interface{}{
		"href":  "test",
		"data":  goods,
		"links": "links",
	}
	var items []interface{} = []interface{}{
		goodsTypeData,
		goodsData,
	}

	var queries map[string]interface{} = map[string]interface{}{
		"rel":    "search",
		"href":   "http://example.org/friends/search",
		"prompt": "Search",
		"data":   "data",
	}

	var template map[string]interface{} = map[string]interface{}{
		"data": "test",
	}

	var collection map[string]interface{} = map[string]interface{}{
		"version":  version,
		"href":     href,
		"links":    links,
		"items":    items,
		"queries":  queries,
		"template": template,
	}

	c.JSON(200, gin.H{
		"collection": collection,
	})

}

func OneGoods() {

}

func UpdateGoods() {

}

func DeleteGoods() {
	// return models.DeleteGoods()
}
