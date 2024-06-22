package main

import (
	"fmt"
	"net/http"
)

func ZoneKisumu(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Recode the world")
}

func LakeHUb(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "LakeHUb KIsumu")
}

func main(){
	mux := http.NewServeMux()// this line is responsible for routing between pages
	mux.HandleFunc("/lakehub", LakeHUb) //  calling the handlers at specific pages/ routes
	mux.HandleFunc("/zone01", ZoneKisumu)


	//http.HandleFunc("/", ZoneKisumu)
	// http.ListenAndServe(":8008", nil)
	fmt.Println("Running server at localhost:8008")
	// http.ListenAndServe(":8008", mux)

	err := http.ListenAndServe(":8008", mux)
	if err != nil{
		fmt.Println("Error starting server:", err)
	}
	
}