package entity

type Product struct {
	ID          int    `gorm:"column:id;not null;primaryKey"`
	Name        string `gorm:"column:name;not null"`
	Price       int64  `gorm:"column:price;not null"`
	Description string `gorm:"column:description;not null"`
	Quantity    int    `gorm:"column:quantity;not null"`
}
