package main 

import (
	"fmt"
)

func main(){
	for i:=1 ;i<1000;i++ {
		go func(){
			fmt.Println(i)
		}()
	}
}