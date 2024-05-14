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
var Data Response

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
	    w.WriteHeader(http.StatusOK)
	 	Data.Status = "success"

		json.NewEncoder(w).Encode(&Data)

		}else{
		http.Error(w ," invalid method ",http.StatusBadRequest)
		}
        return
}

func CreateProducts(w http.ResponseWriter,r *http.Request){
		if r.Method == "POST"{
		 w.WriteHeader(http.StatusOK)
      	 Data.Status ="success"
         product := Product{
         	ID: len(Data.Products)+1,
         }
		 json.NewDecoder(r.Body).Decode(&product)
		 Data.Products = append(Data.Products,product)
         type data struct{
         	Status string `json:"status"`
            Datas Product  `json:"data"`
         }
        var  sendata data
         sendata.Status = Data.Status
         sendata.Datas = product




		 json.NewEncoder(w).Encode(sendata)


		}else{

		http.Error(w ," invalid method ",http.StatusBadRequest)
		}

		return



}


func UpdateProducts(w http.ResponseWriter,r *http.Request){
	query := r.URL.Query()
    id, _ := strconv.Atoi(query.Get("id"))
    var product []Product
    product = Data.Products
	if r.Method == "PUT" {
		for index,value := range product{
			json.NewDecoder(r.Body).Decode(&product)
			if value.ID == id{
				value.NAME = product[index].NAME
				value.PRICE = product[index].PRICE
				w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"message": "Success to update product"}`))


			}
		}
	}else{
	http.Error(w ," invalid method ",http.StatusBadRequest)
	}
	return
}

func DeleteProducts(w http.ResponseWriter,r *http.Request){
	query := r.URL.Query()
    id, _ := strconv.Atoi(query.Get("id"))
	if r.Method == "DELETE" {
		for index,value := range Data.Products{
			json.NewDecoder(r.Body).Decode(&Data)
			if value.ID == id{
				Data.Products = append(Data.Products[:index],Data.Products[index+1:]...)
				w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"message": "Success to delete product"}`))
			}
		}
	}else{
	http.Error(w ," invalid method ",http.StatusBadRequest)
	}
	return
}