package entity

type Book struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"type:varchar(255);not null"`
	Quantity    int    `gorm:"default:0;not null"`
	Type        string `gorm:"type:varchar(100)"`
	Description string
	Supplier    string  `gorm:"type:varchar(255)"`
	Price       float64 `gorm:"not null"`
	Language    string  `gorm:"type:varchar(100)"`
	Cover       string  `gorm:"type:varchar(255)"`
	Year        int
	PageNumber  int
	Categories  []Category `gorm:"many2many:category_book;"`
}
