package model

import (
	"gofaka/utils/errmsg"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserId       uint   `gorm:"" json:"userId"`
	Email        string `gorm:"type:varchar(64);not null" json:"email"`
	SearchPasswd string `gorm:"type:varchar(64)" json:"searchPasswd"`
	GoodsId      int    `gorm:"" json:"goodsId"`
	BuyAmount    int    `gorm:"type:integer(11);not null" json:"buyAmount"`
	OutTradeNo   string `gorm:"type:varchar(20);not null" json:"outTradeNo"`
	CallbackNo   string `gorm:"type:varchar(20)" json:"callbackNo"`
	TotalAmount  int    `gorm:"type:integer(11);not null" json:"totalAmount"`
	Status       int    `gorm:"type:tinyint(1);not null;comment:'0:待付 1:完成 2:取消'" json:"status"`
}

func CreateOrder(order *Order) int {
	if order.BuyAmount < 1 {
		return errmsg.ErrorInvalidQuantity
	}
	now := time.Now()
	outTradeNo := now.Format("20060102150405.999")
	order.OutTradeNo = outTradeNo[:14] + outTradeNo[15:]
	order.CreatedAt = now
	goods, errCode := GetGoods(order.GoodsId)
	if errCode != errmsg.SUCCESS {
		return errCode
	}
	if goods.Stock < order.BuyAmount {
		return errmsg.ErrorInsufficientStock
	}
	goods.Stock -= order.BuyAmount
	errCode = EditGoods(int(goods.ID), &goods)
	if errCode != errmsg.SUCCESS {
		return errCode
	}
	order.TotalAmount = goods.Price * order.BuyAmount
	order.Status = 0
	err := db.Create(&order).Error
	if err != nil {
		return errmsg.ERROR
	}
	errCode = LockItems2Order(order)
	if errCode != errmsg.SUCCESS {
		return errCode
	}
	return errmsg.SUCCESS
}

func CancelOrder(order *Order) int {
	Order, errCode := GetOrder(order.OutTradeNo)
	if errCode != errmsg.SUCCESS {
		return errCode
	}
	if Order.Status != 0 {
		return errmsg.ErrorOrderInvalid
	}
	Goods, errCode := GetGoods(Order.GoodsId)
	if errCode != errmsg.SUCCESS {
		return errCode
	}
	Order.Status = 2
	errCode = EditOrder(&Order)
	if errCode != errmsg.SUCCESS {
		return errCode
	}
	Goods.Stock += Order.BuyAmount
	errCode = EditGoods(int(Goods.ID), &Goods)
	if errCode != errmsg.SUCCESS {
		return errCode
	}
	errCode = UnlockItems2Order(order)
	if errCode != errmsg.SUCCESS {
		return errCode
	}
	return errmsg.SUCCESS
}

func PayOrder(order *Order) (Order, int) {
	Order, errCode := GetOrder(order.OutTradeNo)
	if errCode != errmsg.SUCCESS {
		return Order, errCode
	}
	return Order, errmsg.SUCCESS
}

func GetAllOrder(pageSize int, pageNum int, sortType string, sortKey string) ([]Order, int) {
	var count int64
	var orders []Order
	order := sortKey + " " + sortType
	if sortType == "" || sortKey == "" {
		order = "id desc"
	}
	err := db.Order(order).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&orders).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0
	}
	return orders, int(count)
}

func GetUserOrderByToken(email string, status int) interface{} {
	type Result struct {
		Order
		Title string `json:"title"`
	}
	var result []Result
	err := db.Table("orders").Select("orders.id,orders.created_at,orders.email,orders.buy_amount,orders.out_trade_no,orders.total_amount,orders.status", "goods.title").Joins("join goods on orders.goods_id = goods.id").Where(&Order{Email: email, Status: status}).Scan(&result).Error
	if err != nil {
		return nil
	}
	return result
	//var orders []Order
	//err := db.Where("email=?", email).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&orders).Error
	//if err != nil {
	//	return nil
	//}
	//return orders
}

func GetOrderByEmail(email string, status int) []Order {
	var orders []Order
	err := db.Table("orders").Select("orders.id,orders.created_at,orders.email,orders.buy_amount,orders.out_trade_no,orders.total_amount,orders.status", "goods.title").Joins("join goods on orders.goods_id = goods.id").Where(&Order{Email: email, Status: status}).Find(&orders).Error
	if err != nil {
		return nil
	}
	return orders
}

func EditOrder(data *Order) int {
	var order Order
	var maps = make(map[string]interface{})
	maps["status"] = data.Status
	err := db.Model(&order).Where("out_trade_no=?", data.OutTradeNo).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetOrder(OutTradeNo string) (Order, int) {
	var Order Order
	err := db.Where("out_trade_no=?", OutTradeNo).Find(&Order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Order, errmsg.ERROR
	}
	return Order, errmsg.SUCCESS
}

func GetOrderInfo(OutTradeNo string) (interface{}, int) {
	type Result struct {
		Order
		Title string `json:"title"`
		Price string `json:"price"`
	}
	var result Result
	err := db.Table("orders").Select("orders.id,orders.created_at,orders.buy_amount,orders.out_trade_no,orders.total_amount,orders.status,goods.title,goods.price").Joins("join goods on orders.goods_id = goods.id").Where("out_trade_no = ?", OutTradeNo).Scan(&result).Error
	if err == gorm.ErrRecordNotFound {
		return result, errmsg.ERROR
	}
	if err != nil {
		return result, errmsg.ERROR
	}
	return result, errmsg.SUCCESS
}

func CallBackOrder(OutTradeNo string) int {
	Order, errCode := GetOrder(OutTradeNo)
	if errCode != errmsg.SUCCESS {
		return errCode
	}
	if Order.Status != 0 {
		return errmsg.ErrorOrderInvalid
	}
	Order.Status = 1
	errCode = EditOrder(&Order)
	if errCode != errmsg.SUCCESS {
		return errCode
	}
	errCode = BindItems2Order(OutTradeNo)
	if errCode != errmsg.SUCCESS {
		return errCode
	}
	return errmsg.SUCCESS
}

func CancelExpiredOrder() {
	var orders []Order
	_ = db.Where("status=?", 0).Find(&orders).Error
	for _, order := range orders {
		_ = CancelOrder(&order)
	}
}
