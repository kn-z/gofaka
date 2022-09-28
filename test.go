package main

import (
	"fmt"
	"gofaka/model"
	"gofaka/utils/errmsg"
	"gorm.io/gorm"
)

var db = model.GetDb()

type Item struct {
	gorm.Model
	Title   string `gorm:"type:varchar(20);not null" json:"title"`
	Price   int    `gorm:"type:integer(11);not null" json:"price"`
	Content string `gorm:"type:text" json:"content"`
	Stock   int    `gorm:"type:integer(6);not null" json:"stock"`
	Image   string `gorm:"type:varchar(100)" json:"image"`
}

func CreateItem(data *Item) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetItemList(pageSize int, pageNum int) []Item {
	var items []Item
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&items).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return items
}

func EditItem(id int, data *Item) int {
	var item Item
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["price"] = data.Price
	maps["content"] = data.Content
	maps["stock"] = data.Stock
	maps["image"] = data.Image
	err := db.Model(&item).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteItem(id int) int {
	var item Item
	fmt.Println(&item.Price)
	err := db.Where("id=?", id).Delete(&item).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func main() {
	DeleteItem(3)
}
