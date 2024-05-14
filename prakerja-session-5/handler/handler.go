package handler

import (
	"net/http"
	"encoding/json"
	//"io"
	//"fmt"
	//"strings"
	"strconv"
	//. "prakerja-session-5/database"
	. "prakerja-session-5/model"

)
var Products []Product = []Product{}

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

		if r.Method == "GET"{
		var getResponse = GetResponse{
			Status:"success",
			Products:Products,
		}
	    w.WriteHeader(http.StatusOK)


		json.NewEncoder(w).Encode(&getResponse)

        }
        return
}

func CreateProducts(w http.ResponseWriter,r *http.Request){
		if r.Method == "POST"{
		 w.WriteHeader(201)

         product := Product{
         	ID: len(Products)+1,
         }
		 json.NewDecoder(r.Body).Decode(&product)
		 Products = append(Products,product)

         var createResponse = CreateResponse{
         	Status:"success",
         	Products:product,

         }




		 json.NewEncoder(w).Encode(&createResponse)


		}


		return



}


func UpdateProducts(w http.ResponseWriter,r *http.Request){
	query := r.URL.Query()
    id, _ := strconv.Atoi(query.Get("id"))

	if r.Method == "PUT" {
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

					var updateResponse = UpdateResponse{
							Status:"success",
							Products:product,
	    					}
	    			w.WriteHeader(201)

					json.NewEncoder(w).Encode(updateResponse)

					break

				}

				if limiter >= 1 {
    				http.Error(w, "not allowed ", 422)
    			}



    }

		if len(Products) == 0 {
    		http.Error(w, "Data not found", 404)
		    return
    	}


	}
	return
}

func DeleteProducts(w http.ResponseWriter,r *http.Request){
	query := r.URL.Query()
    id, _ := strconv.Atoi(query.Get("id"))

	if r.Method == "DELETE" {



	var product Product

    product = Product{
    	ID:id,
    }

    Exist := 0

		for index,value := range Products{
			if value.ID == id {
				Exist = index
				Products[index] = product
                Products = append(Products[:index],Products[index+1:]...)
                var deleteResponse = DeleteResponse{
									Status:"success",
									Products: nil,
									}

    			json.NewEncoder(w).Encode(deleteResponse)
    			break


            }

            if Exist == 0 {
    				http.Error(w, "Data not found", 404)
					return
    		}
    	}

    	if len(Products) == 0 {
    		http.Error(w, "Data not found", 404)
		    return
    	}

}
	return
}