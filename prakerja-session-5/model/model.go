package model



type Product struct{
	ID		int 	`json:"id"`
	NAME	string 	`json:"name"`
	PRICE   string 	`json:"price"`
}


type GetResponse struct{
	Status string 	`json:"status"`
	Products  []Product `json:"data"`
}


type CreateResponse struct{
	Status string 	`json:"status"`
	Products  Product `json:"data"`
}

type UpdateResponse struct{
	Status string 	`json:"status"`
	Products  Product `json:"data"`
}

type DeleteResponse struct{
	Status string 	`json:"status"`
	Products  any `json:"data"`
}
