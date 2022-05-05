package model

import (
	"gofaka/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"username"`
}

func CheckCategory(name string) (code int) {
	var category Category
	db.Select("id").Where("id = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //1001
	}
	return errmsg.SUCCESS
}

//add
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//get list
func GetCategory(pageSize int, pageNum int) []Category {
	var categories []Category
	//Limit(x) Offset(y) skip y datas and read x datas || Limit(5) Offset((2-1)*5)
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return categories
}

//edit user
func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&category).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//get category article

//delete user
func DeleteCategory(id int) int {
	var category Category
	err = db.Where("id=?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
