package model

import (
	"gofaka/utils/errmsg"
	"gorm.io/gorm"
)

type Notice struct {
	gorm.Model
	Title   string `gorm:"type:varchar(20);not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	Status  int    `gorm:"type:tinyint" json:"status"`
	Image   string `gorm:"type:varchar(100)" json:"image"`
}

func CreateNotice(data *Notice) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func EditNotice(id int, data *Notice) int {
	var notice Notice
	var maps = make(map[string]interface{})
	if len(data.Title) == 0 && len(data.Content) == 0 && len(data.Image) == 0 {
		maps["status"] = data.Status
	} else {
		maps["title"] = data.Title
		maps["status"] = data.Status
		maps["Content"] = data.Content
		maps["image"] = data.Image
	}
	err := db.Model(&notice).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetNoticeList(pageSize int, pageNum int) []Notice {
	var notices []Notice
	err = db.Order("updated_at desc").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("status=?", 1).Find(&notices).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return notices
}

func GetAllNotice(pageSize int, pageNum int) []Notice {
	var notices []Notice
	//Limit(x) Offset(y) skip y datas and read x datas || Limit(5) Offset((2-1)*5)
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&notices).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return notices
}

func GetNoticeByID(id int) (Notice, int) {
	var notice Notice
	rows, err := db.Where("id=?", id).Find(&notice).Rows()
	if err == gorm.ErrRecordNotFound {
		return notice, errmsg.ErrorNoticeNotExist
	}
	if err != nil {
		return notice, errmsg.ERROR
	}
	if !rows.Next() {
		return notice, errmsg.ErrorNoticeNotExist
	}
	return notice, errmsg.SUCCESS
}

func DeleteNotice(id int) int {
	var notice Notice
	err = db.Where("id=?", id).Delete(&notice).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
