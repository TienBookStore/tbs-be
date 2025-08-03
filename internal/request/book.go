package request

type ReqCreateBook struct {
	Title       string   `json:"title" binding:"required"`
	Quantity    int      `json:"quantity"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Supplier    string   `json:"supplier"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Language    string   `json:"language"`
	Cover       string   `json:"cover"`
	Year        int      `json:"year"`
	PageNumber  int      `json:"page_number"`
	CategoryIDs []string `json:"category_ids"` // danh sách ID category (UUID hoặc string)
}
