package entity

// Cart struct
type Cart struct {
	ID              int     `gorm:"column:id;primary_key;auto_increment;not null"`
	TransactionCode string  `gorm:"column:transaction_code;not null"`
	TotalPrice      float64 `gorm:"column:total_price;not null"`
	CustomerID      int     `gorm:"column:customer_id;not null"`
	Customer        Users   `gorm:"foreignKey:CustomerID"`
	BaseModelWithDeleteAt
}

func (c *Cart) TableName() string {
	return "brands"
}
