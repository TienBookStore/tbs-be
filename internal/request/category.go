package request

type ReqCreateCategory struct {
	Name string `json:"name" binding:"required,min=3,max=50"`
}

type ReqUpdateCategory struct {
	ID   string `json:"id" binding:"required,uuid"`
	Name string `json:"name" binding:"required,min=3,max=50"`
}
