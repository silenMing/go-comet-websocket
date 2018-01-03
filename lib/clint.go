package lib

import(
	"net/http"
	"log"
)

func clint(){
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}