package models


type ProductListOrOne interface {
	[]Product | Product
}

type ProdResponseSuccess[T ProductListOrOne] struct {
	Status string `json:"status"`
	Data T `json:"data"`
}

type ProdResponseError struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

type ProdResponse[T ProductListOrOne] struct{
	Code int
	Success *ProdResponseSuccess[T]
	Error *ProdResponseError

}