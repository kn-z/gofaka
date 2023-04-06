package model

import (
	"encoding/base64"
	"gofaka/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(64);primaryKey;not null" json:"email"`
	Password string `gorm:"type:varchar(64);not null" json:"password"`
	Status   uint   `gorm:"type:tinyint;default:0" json:"status"`
	Role     int    `gorm:"type:int" json:"role"`
	Content  string `gorm:"type:text" json:"content"`
}

func CheckEmail(email string) (code int) {
	var count int64
	db.Where("email = ?", email).Find(&[]User{}).Count(&count)
	if count > 0 {
		return errmsg.ErrorEmailExist //1012
	}
	return errmsg.SUCCESS
}

func CheckUser(user User) (code int) {
	var count int64
	db.Where("id <> ? AND email = ?", user.ID, user.Email).Find(&[]User{}).Count(&count)
	if count > 0 {
		return errmsg.ErrorEmailUsed //1012
	}
	return errmsg.SUCCESS
}

// add user
func CreateUser(data *User) int {
	if len(data.Password) < 8 && len(data.Password) > 0 {
		return errmsg.ErrorPasswordLess8
	}
	if len(data.Password) == 0 {
		data.Password = data.Email
	}
	data.Password = ScryptPwd(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetUserByID(id int) (User, int) {
	var user User
	rows, err := db.Where("id=?", id).Find(&user).Rows()
	if err == gorm.ErrRecordNotFound {
		return user, errmsg.ErrorEmailNotExist
	}
	if err != nil {
		return user, errmsg.ERROR
	}
	if !rows.Next() {
		return user, errmsg.ErrorUserNotExist
	}
	return user, errmsg.SUCCESS
}

func GetUser(email string) (User, int) {
	var user User
	err := db.Where("email=?", email).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, errmsg.ErrorEmailNotExist
	}
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// get users list
func GetUsers(pageSize int, pageNum int, sortType string, sortKey string) ([]User, int) {
	var users []User
	var count int64
	order := sortKey + " " + sortType
	if sortType == "" || sortKey == "" {
		order = "id desc"
	}
	//Limit(x) Offset(y) skip y datas and read x datas || Limit(5) Offset((2-1)*5)
	err = db.Order(order).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, int(count)
}

// edit user
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["email"] = data.Email
	if len(data.Password) < 8 && len(data.Password) > 0 {
		return errmsg.ErrorPasswordLess8
	}
	if len(data.Password) != 0 {
		maps["password"] = ScryptPwd(data.Password)
	}
	maps["status"] = data.Status
	maps["role"] = data.Role
	maps["content"] = data.Content
	err := db.Model(&user).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// delete user
func DeleteUser(id int) int {
	var user User
	err = db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//passwd
//func (data *User) BeforeSave() {
//	data.Password = ScryptPwd(data.Password)
//}

func ScryptPwd(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 78, 45, 12, 36, 4}
	HashPwd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPwd)
	return fpw
}

// check login
func CheckLogin(email string, password string) (int, int) {
	var user User
	db.Where("email = ?", email).Find(&user)
	if user.ID == 0 {
		return -1, errmsg.ErrorEmailNotExist
	}
	if user.Password != ScryptPwd(password) {
		return -1, errmsg.ErrorPasswordWrong
	}
	if user.Status == 1 {
		return -1, errmsg.ErrorUserBanned
	}
	return user.Role, errmsg.SUCCESS
}
