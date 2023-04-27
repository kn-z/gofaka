package model

import (
	"gofaka/utils/errmsg"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);not null" json:"name"`
	Image        string `gorm:"type:text" json:"image"`
	Profile      string `gorm:"type:text" json:"profile"`
	NotifyDomain string `gorm:"type:varchar(200)" json:"notify_domain"`
	Status       int    `gorm:"type:int(1);not null" json:"status"`
	Type         int    `gorm:"type:int(1);not null;comment:'0:alipay 1:wechat pay'" json:"type"`
}

func CreatePayment(payment *Payment) int {
	err := db.Create(&payment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetPayment(id int) (Payment, int) {
	var payment Payment
	err = db.Where("payments.id = ?", id).Find(&payment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return payment, errmsg.ERROR
	}
	return payment, errmsg.SUCCESS
}

func EditPayment(id int, data *Payment) int {
	var payment Payment
	var maps = make(map[string]interface{})
	if len(data.Name) == 0 && len(data.Image) == 0 && len(data.Profile) == 0 && len(data.NotifyDomain) == 0 && data.Type == 0 {
		maps["status"] = data.Status
	} else {
		maps["name"] = data.Name
		maps["image"] = data.Image
		maps["profile"] = data.Profile
		maps["notify_domain"] = data.NotifyDomain
		maps["type"] = data.Type
	}
	err = db.Model(&payment).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetPaymentList(pageSize int, pageNum int) ([]Payment, int) {
	var payments []Payment
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Select("payments.id, payments.name, payments.image").Where("payments.status = ?", 1).Find(&payments).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return payments, errmsg.ERROR
	}
	return payments, errmsg.SUCCESS
}

func GetAllPayment(pageSize int, pageNum int) ([]Payment, int) {
	var payments []Payment
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&payments).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return payments, errmsg.ERROR
	}
	return payments, errmsg.SUCCESS
}
