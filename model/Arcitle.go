package model

import (
	"fmt"
	"gofaka/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

//add
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//todo get article in a category

//get list
func GetArticle(pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	//Limit(x) Offset(y) skip y datas and read x datas || Limit(5) Offset((2-1)*5)
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articles, errmsg.SUCCESS
}

//get one
func GetArtInfo(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		fmt.Println(err)
		return article, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

//query article in a category
func GetCateArt(cid int, pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", cid).Find(&articles).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}
	return articles, errmsg.SUCCESS
}

//edit
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&article).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//delete
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id=?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
