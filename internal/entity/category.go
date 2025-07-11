package entity

type Category struct {
	ID    string `json:"id" gorm:"type:varchar(36);primaryKey"`
	Name  string `gorm:"type:varchar(255);not null" json:"name"`
	Slug  string `gorm:"type:varchar(255);unique;not null" json:"slug"`
	Books []Book `gorm:"many2many:category_book;" json:"books,omitempty"`
}
