package main

import (
	"encoding/json"
	"fmt"
    "log"
    "net/http"
)

type JSONResponse struct {
	Value1 string `json:"key1"`
	Value2 string `json:"key2"`
}

func process() []byte {

	jsonResponse := JSONResponse{
		Value1: "Test value app3",
		Value2: "Test value app3",
	}

	fmt.Printf("The struct returned before marshalling\n\n")
	fmt.Printf("%+v\n\n\n\n", jsonResponse)

	// The MarshalIndent function only serves to pretty print, json.Marshal() is what would normally be used
	byteArray, err := json.MarshalIndent(jsonResponse, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The JSON response returned when the struct is marshalled\n\n")
	fmt.Println(string(byteArray))

	return byteArray
}

func homePage(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(process())
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10003", nil))
}

func main(){
	log.Println("UP")
	handleRequests()
}