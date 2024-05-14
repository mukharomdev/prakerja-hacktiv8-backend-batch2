package model



type Product struct{
	ID		int 	`json:"id"`
	NAME	string 	`json:"name"`
	PRICE   string 	`json:"price"`
}


type Response struct{
	Status string 	`json:"status"`
	Products  []Product `json:"data"`
}