package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",Myhandler1)
	http.HandleFunc("/sifat",Myhandler2)
	http.ListenAndServe(":8000",nil)
}

func Myhandler1(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello world\n")
}

func Myhandler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello Sifat\n")
}