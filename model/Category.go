package model

import (
	"gofaka/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name   string `gorm:"type:varchar(20);not null" json:"name"`
	Status int    `gorm:"type:tinyint" json:"status"`
}

func CheckCategory(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ErrorCateNameExist
	}
	return errmsg.SUCCESS
}

// add
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// get list
func GetCategory(pageSize int, pageNum int) []Category {
	var categories []Category
	//Limit(x) Offset(y) skip y datas and read x datas || Limit(5) Offset((2-1)*5)
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Where("status=?", 1).Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return categories
}

// edit
func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	if len(data.Name) == 0 {
		maps["status"] = data.Status
	} else {
		maps["name"] = data.Name
		maps["status"] = data.Status
	}
	err := db.Model(&category).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//get category

// delete user
func DeleteCategory(id int) int {
	var category Category
	err = db.Where("id=?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetCategoryInfo(id int) (Category, int) {
	var category Category
	//Limit(x) Offset(y) skip y datas and read x datas || Limit(5) Offset((2-1)*5)
	err := db.Where("id=?", id).Find(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return category, errmsg.ERROR
	}
	return category, errmsg.SUCCESS
}

func GetAllCategory() interface{} {
	type Result struct {
		Category
		GoodsNum int `gorm:"" json:"goodsNum"`
	}
	var result []Result
	//Limit(x) Offset(y) skip y datas and read x datas || Limit(5) Offset((2-1)*5)
	err = db.Table("categories").Find(&result).Error
	for idx, cate := range result {
		result[idx].GoodsNum = CountGoodsNum(int(cate.ID))
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return result
}

func CountGoodsNum(id int) int {
	var count int64
	err := db.Model(&Goods{}).Where("cate_id=?", id).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0
	}
	return int(count)
}
