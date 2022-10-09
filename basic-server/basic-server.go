package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", useCors(homeHandler))
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func useCors(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)

	//---------- optioninal ---------------------
	//handling Errors
	if err != nil {
		log.Fatal(err)
	}

	//print result
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	fmt.Println(req.Header)
	fmt.Println(req.URL.Path)
	fmt.Println(req.URL.RawQuery)
	var bodyStr string
	json.NewDecoder(req.Body).Decode(&bodyStr)
	fmt.Println(bodyStr)
	json.NewEncoder(w).Encode(req.Body)
}
