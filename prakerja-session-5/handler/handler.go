package handler

import (
	"net/http"
	"encoding/json"
	//"io"
	//"fmt"
	//"strings"
	//"log"
	"strconv"
	//. "prakerja-session-5/database"
	. "prakerja-session-5/model"

)
var Products []Product = []Product{}

func GenerateProductId() int {
	if len(Products) > 0 {
		productId := Products[len(Products)-1].ID

		//Id, _ := strconv.Atoi(productId)

		return productId + 1
	}

	return 1
}


func HandlerProducts(w http.ResponseWriter,r *http.Request){
	w.Header().Add("Content-Type","application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

     switch r.Method {

	 case "GET":
			GetProducts(w,r)
	 case "POST":
			CreateProducts(w,r)
	 case "PUT":
			UpdateProducts(w,r)
	 case "DELETE":
			DeleteProducts(w,r)
	 default:
			http.Error(w ," invalid method ",http.StatusBadRequest)

      }
  }

func GetProducts(w http.ResponseWriter,r *http.Request){


		var getResponse = GetResponse{
			Status:"success",
			Products:Products,
		}

	    w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&getResponse)
        return
}

func CreateProducts(w http.ResponseWriter,r *http.Request){
		 product := Product{
         	ID: GenerateProductId(),
         }

		 json.NewDecoder(r.Body).Decode(&product)

		 Products = append(Products,product)

         var createResponse = CreateResponse{
         	Status:"success",
         	Products:product,

         }

         w.WriteHeader(201)
		 json.NewEncoder(w).Encode(&createResponse)
		return
}


func UpdateProducts(w http.ResponseWriter,r *http.Request){
	query := r.URL.Query()
    id, _ := strconv.Atoi(query.Get("id"))

	var product Product

	product = Product{
 			ID:id,
 	}

    json.NewDecoder(r.Body).Decode(&product)

    limiter := 0

	for index,value := range Products {

				if value.ID == id {
					limiter++
					Products[index]  = product
		 			Products[index].NAME	= product.NAME
		 			Products[index].PRICE	= product.PRICE
					break
				}

				if limiter >= 1 {
			    	http.Error(w, "not allowed ", 422)
			    	return
			    }

			    if len(Products) == 0 {
			    		http.Error(w, "Data not found", 404)
				    return
			   	}
	}

	if len(Products) == 0{
		http.Error(w,"Tak ada data",404)
		return
	}


	var updateResponse = UpdateResponse{
		Status:"success",
		Products:product,
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(updateResponse)
	return
}

func DeleteProducts(w http.ResponseWriter,r *http.Request){
	query := r.URL.Query()
    id, _ := strconv.Atoi(query.Get("id"))

	// product := Product{
 // 			ID:id,
 // 	}

    //json.NewDecoder(r.Body).Decode(&product)

    index := 0
	for i,v := range Products{

		if v.ID == id {
		   index = i
		   break
	    }
	    if index == 0 {
			http.Error(w, "Data not found", 404)
			return
		}
    }

	if len(Products) == 0{
		http.Error(w,"Tak ada data",404)
		return
	}



    Products = append(Products[:index],Products[index+1:]...)


	respons := DeleteResponse{
    	Status:"success",
       	Products  : nil,
    }

    //log.Println(index)
    w.WriteHeader(200)
  	json.NewEncoder(w).Encode(respons)
	return
}