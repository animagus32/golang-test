package main
import "testing"
import "fmt" 

func Test_001(t *testing.T){
	c := make(chan int)
	defer close(c)
	c<-1 
	a:= <-c 
	fmt.Printf("%d ",a) 

}