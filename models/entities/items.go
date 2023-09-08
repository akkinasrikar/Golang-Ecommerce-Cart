package entities


type Item struct {
	ItemID          int     `gorm:"column:item_id;primaryKey"`
	ItemTitle       string  `gorm:"column:item_title;not null"`
	ItemPrice       float64 `gorm:"column:item_price;not null"`
	ItemDescription string  `gorm:"column:item_description;not null"`
	ItemCategory    string  `gorm:"column:item_category;not null"`
	ItemImage       string  `gorm:"column:item_image;not null"`
	ItemRating      float64 `gorm:"column:item_rating;not null"`
	ItemCount       int     `gorm:"column:item_count;not null"`
}
