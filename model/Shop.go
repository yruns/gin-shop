package model

type Shop struct {
	Id            int64  `gorm:"column:id;primary_key" json:"id"`
	Name          string `gorm:"column:name" json:"name"`
	PromotionInfo string `gorm:"column:promotion_info" json:"promotion_info"`
	Address       string `gorm:"column:address" json:"address"`
	Phone         string `gorm:"column:phone" json:"phone"`
	Status        int    `gorm:"column:status" json:"status"`

	Longitude float64 `gorm:"column:longitude" json:"longitude"`
	Latitude  float64 `gorm:"column:latitude" json:"latitude"`
	ImagePath string  `gorm:"column:image_path" json:"image_path"`

	IsNew     bool `gorm:"column:is_new" json:"is_new"`
	IsPremium bool `gorm:"column:is_premium" json:"is_premium"`

	Rating            float64 `gorm:"column:rating" json:"rating"`
	RatingCount       int     `gorm:"column:rating_count" json:"rating_count"`
	RecentOrderNum    int     `gorm:"column:recent_order_num" json:"recent_order_num"`
	MinimumOrderAmout int     `gorm:"column:minimum_order_amount" json:"minimum_order_amount"`
	DeliveryFee       int     `gorm:"column:delivery_fee" json:"delivery_fee"`

	OpeningHours string `gorm:"column:opening_hours" json:"opening_hours"`
}
