package main 


import (
	"net/http"
	"log"
	"io"
) 


func main() {


	http.HandleFunc("/hello",Hello)
	err := http.ListenAndServe("127.0.0.1:12345",nil)
	if err != nil {
		log.Fatal("error")
		
	}

}

func Hello(w http.ResponseWriter , r *http.Request){
	
	w.Write([]byte("hellloooasdjfj"))
	io.WriteString(w, "hello, world!\n")
}