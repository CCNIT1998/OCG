package model

type Product struct {
	Id         int64   `json:"id"`
	CategoryId int64  `json:"categoryId"`
	Image      []string  `json:"image"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	IsSale     bool    `json:"isSale"`
	CreatedAt  string  `json:"createdAt"`
	ModifiedAt string  `json:"modifiedAt"`
}
