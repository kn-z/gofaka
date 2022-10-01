package model

import (
	"gofaka/utils/errmsg"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	UserId      uint   `gorm:"foreignKey:UserId" json:"userId"`
	OutTradeNo  string `gorm:"type:varchar(20);not null" json:"outTradeNo"`
	CallbackNo  string `gorm:"type:varchar(20)" json:"callbackNo"`
	TotalAmount int    `gorm:"type:uint(11);not null" json:"totalAmount"`
	Status      uint   `gorm:"type:uint(1);not null" json:"status"`
}

func CreateOrder(order *Order) int {
	err := db.Create(order).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetOrders(userId uint, pageSize int, pageNum int) []Order {
	var orders []Order
	err := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&orders).Where("user_id=?", userId).Error
	if err != nil {
		return nil
	}
	return orders
}

func EditOrders(id int, data *Order) int {
	var order Order
	var maps = make(map[string]interface{})
	maps["status"] = data.Status
	err := db.Model(&order).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
