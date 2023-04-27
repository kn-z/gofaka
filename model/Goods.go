package model

import (
	"gofaka/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

type Goods struct {
	gorm.Model
	Title   string `gorm:"type:varchar(20);not null" json:"title"`
	Price   int    `gorm:"type:integer(11);not null" json:"price"`
	Content string `gorm:"type:text" json:"content"`
	Stock   int    `gorm:"type:integer(6);not null" json:"stock"`
	Image   string `gorm:"type:text" json:"image"`
	CateId  int    `gorm:"" json:"cateId"`
	Sort    uint   `gorm:"type:int;autoIncrement" json:"sort"`
	Status  int    `gorm:"type:int" json:"status"`
}

func CreateGoods(data *Goods) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//func GetGoodsList(pageSize int, pageNum int) []Goods {
//	var Goods []Goods
//	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&Goods).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nils
//	}
//	return Goods
//}

func GetGoodsList(pageSize int, pageNum int) []Goods {
	var goodsList []Goods
	err := db.Order("sort").Limit(pageSize).Offset((pageNum-1)*pageSize).Joins("left join categories on goods.cate_id = categories.id").Where("goods.status = ? AND categories.status = ?", 1, 1).Find(&goodsList).Error
	//	err := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Where("status=?", 1).Find(&Goods).Error
	for idx, goods := range goodsList {
		goodsList[idx].Stock = CountItemNum(int(goods.ID))
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return goodsList
}

func GetAllGoods(pageSize int, pageNum int) interface{} {
	type Result struct {
		Goods
		Name string `gorm:"" json:"name"`
	}
	var result []Result
	err := db.Order("sort desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Table("goods").Select("goods.id,goods.created_at,goods.title,goods.status,goods.cate_id,goods.price,goods.image,goods.content,categories.name").Joins("left join categories on goods.cate_id = categories.id").Find(&result).Error
	//err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&Goods).Error
	for idx, goods := range result {
		result[idx].Stock = CountItemNum(int(goods.ID))
		if result[idx].Name == "" {
			result[idx].Name = "无"
		}
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return result
}

func EditGoods(id int, data *Goods) int {
	var goods Goods
	var maps = make(map[string]interface{})
	if len(data.Title) == 0 && data.Price == 0 && data.CateId == 0 {
		maps["status"] = data.Status
	} else {
		maps["title"] = data.Title
		maps["price"] = data.Price
		maps["content"] = data.Content
		maps["stock"] = CountItemNum(id)
		maps["image"] = data.Image
		maps["cate_id"] = data.CateId
		maps["status"] = data.Status
	}
	err := db.Model(&goods).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteGoods(id int) int {
	var Good Goods
	err := db.Where("id=?", id).Delete(&Good).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetGoods(id int) (Goods, int) {
	var Goods Goods
	err := db.Where("id=?", id).Find(&Goods).Error
	Goods.Stock = CountItemNum(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return Goods, errmsg.ERROR
	}
	return Goods, errmsg.SUCCESS
}

func CheckGoodsStock() int {
	GoodsList := GetGoodsList(0, 0)
	for _, goods := range GoodsList {
		goods.Stock = CountItemNum(int(goods.ID))
		err := EditGoods(int(goods.ID), &goods)
		if err != errmsg.SUCCESS {
			return err
		}
	}
	return errmsg.SUCCESS
}

func UpdateSort(id int, sort uint) int {
	var goods Goods
	var maps = make(map[string]interface{})
	maps["sort"] = sort
	err := db.Model(&goods).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func SortGoods(orders map[string]int) int {
	for id, sort := range orders {
		log.Printf("%s\t%d\n", id, sort)
	}
	return errmsg.SUCCESS
}
