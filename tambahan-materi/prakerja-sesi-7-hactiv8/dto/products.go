package dto

type CreateProductRequestDto struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
