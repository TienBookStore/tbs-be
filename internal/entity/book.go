package entity

type Book struct {
	ID          string     `json:"id" gorm:"type:varchar(36);primaryKey"`
	Title       string     `json:"title" gorm:"type:varchar(255);not null"`
	Quantity    int        `json:"quantity" gorm:"default:0;not null"`
	Type        string     `json:"type" gorm:"type:varchar(100)"`
	Description string     `json:"description" gorm:"type:text"`
	Supplier    string     `json:"supplier" gorm:"type:varchar(255)"`
	Price       float64    `json:"price" gorm:"not null"`
	Language    string     `json:"language" gorm:"type:varchar(100)"`
	Cover       string     `json:"cover" gorm:"type:varchar(255)"`
	Year        int        `json:"year" gorm:"type:int"`
	PageNumber  int        `json:"page_number" gorm:"type:int"`
	Categories  []Category `json:"categories,omitempty" gorm:"many2many:category_book;"`
}
