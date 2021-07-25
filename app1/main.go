package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
)

func MakeRequest2(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://bapp2:10002")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
    w.Write(body)
    
}

func MakeRequest3(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://bapp3:10003")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
    w.Write(body)
    
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/get2", MakeRequest2)
    http.HandleFunc("/get3", MakeRequest3)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	log.Println("UP")
    handleRequests()
}