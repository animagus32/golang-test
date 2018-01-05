package main  
  
import (  
    "flag"  
    "net/http"  
    "fmt"
    "math"
)  
  
func main() {  
    host := flag.String("host", "127.0.0.1", "listen host")  
    port := flag.String("port", "8001", "listen port")  
  
    fmt.Printf("%d",int(math.Pow(2,10)))
    http.HandleFunc("/hello", Hello)  
    err := http.ListenAndServe(*host+":"+*port, nil)  
  
    if err != nil {  
        panic(err)  
    }  
}  
// hello 函数
// test 
/* Output:
daldjfl
*/
func Hello(w http.ResponseWriter, req *http.Request) {  
   w.Write([]byte("Hello World"))
}  
