package main 

import (
	"net/http"
)

func main(){
	handler := http.FileServer(http.Dir("../../"))
	http.ListenAndServe("127.0.0.1:8801",handler)
}