package models

// Product - specifies product type
type Product struct {
	ID           uint `gorm:"primarykey"`
	CategoryName string
	Category     Category `gorm:"foreignKey:CategoryName"`
	Name         string   `gorm:"unique;not null"`
	Price        float64
}

// Category - specifies category type
type Category struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"unique;not null"`
}

// Cart - describes cart type userID is assummed to be cart ID
type Cart struct {
	ID           uint `gorm:"primarykey"`
	UserID       uint
	User         User `gorm:"foreignKey:UserID"`
	ProductID    uint
	CartProducts []Product `gorm:"foreignKey:ID;references:ProductID"`
}

// Invoice - describes bill
type Invoice struct {
	ID          uint `gorm:"primarykey"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	TotalAmount float64
	Discount    float64
	FinalAmount float64
}

// User - decribes user
type User struct {
	ID      uint   `gorm:"primarykey"`
	Name    string `gorm:"unique;not null" json:"name"`
	IsAdmin bool   `json:"isAdmin"`
}
