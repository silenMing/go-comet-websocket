package main

import(
	"net/http/cgi"
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		handler := new(cgi.Handler);
	})
		
	
}

