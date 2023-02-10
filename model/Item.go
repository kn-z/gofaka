package model

import (
	"gofaka/utils/errmsg"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Status     int    `gorm:"type:tinyint(1);not null;comment:'0:空闲 1:已售 2:待售'" json:"status"`
	GoodsId    int    `gorm:"" json:"goodsId"`
	OutTradeNo string `gorm:"type:varchar(20);not null" json:"outTradeNo"`
	IsLoop     int    `gorm:"type:tinyint(1);not null" json:"isLoop"`
	Cami       string `gorm:"type:text" json:"cami"`
}

func CreateItem(data *Item) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetAllItem(pageSize int, pageNum int, sortType string, sortKey string) (interface{}, int) {
	type Result struct {
		Item
		Title string `gorm:"" json:"title"`
	}
	var result []Result
	var count int64
	var order string
	order = sortKey + " " + sortType
	if sortType == "" || sortKey == "" {
		order = "id desc"
	}
	err := db.Table("items").Order(order).Select("items.id,items.created_at,items.out_trade_no,items.goods_id,items.cami,items.status,items.out_trade_no,goods.title").Limit(pageSize).Offset((pageNum - 1) * pageSize).Joins("left join goods on items.goods_id = goods.id").Scan(&result).Limit(-1).Offset(-1).Count(&count).Error
	//err := db.Order(order).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&items).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return result, int(count)
}

func GetItemListByOrder(outTradeNo string) interface{} {
	type Result struct {
		Item
		Title string `gorm:"" json:"title"`
	}
	var result []Result
	err := db.Table("items").Select("items.id,items.created_at,items.out_trade_no,items.cami,goods.title").Joins("join goods on items.goods_id = goods.id").Where("out_trade_no=? and items.status=?", outTradeNo, 1).Scan(&result).Error
	//err := db.Where("out_trade_no=? and status=?", outTradeNo, 1).Find(&Item).Error
	//for idx,items := range result {
	//	//result[idx].Stock = CountItemNum(int(items.ID))
	//	if result[idx].Name == "" {
	//		result[idx].Name = "无"
	//	}
	//}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return result
}

func EditItem(id int, data *Item) int {
	var maps = make(map[string]interface{})
	maps["status"] = data.Status
	maps["goods_id"] = data.GoodsId
	maps["is_loop"] = data.IsLoop
	maps["cami"] = data.Cami
	err := db.Model(&Item{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteItem(id int) int {
	var Item Item
	err := db.Where("id=?", id).Delete(&Item).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetItem(id int) (Item, int) {
	var Item Item
	err := db.Where("id=?", id).Find(&Item).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Item, errmsg.ERROR
	}
	return Item, errmsg.SUCCESS
}

func CountItemNum(goodsId int) int {
	var count int64
	err := db.Model(&Item{}).Where("goods_id=? and status=?", goodsId, 0).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0
	}
	return int(count)
}

func LockItems2Order(order *Order) int {
	var maps = make(map[string]interface{})
	maps["out_trade_no"] = order.OutTradeNo
	maps["status"] = 2

	err := db.Model(&Item{}).Limit(order.BuyAmount).Where("goods_id=? and status=?", order.GoodsId, 0).Updates(maps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func UnlockItems2Order(order *Order) int {
	var maps = make(map[string]interface{})
	maps["out_trade_no"] = ""
	maps["status"] = 0
	err := db.Model(&Item{}).Where("out_trade_no=?", order.OutTradeNo).Updates(maps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func BindItems2Order(OutTradeNo string) int {
	var maps = make(map[string]interface{})
	maps["status"] = 1
	err := db.Model(&Item{}).Where("out_trade_no=?", OutTradeNo).Updates(maps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
