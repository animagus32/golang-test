package main

func main(){

	c := make(chan struct{})

	go func(){
		c <- struct{}{}
	}()

}