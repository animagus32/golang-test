package main 

import (
	"fmt"
	// "reflect"
	"time"
	"animagus/routinepool"
	"golang.org/x/net/context"
)

func main(){
	// for i:=1 ;i<1000;i++ {
	// 	go func(){
	// 		fmt.Println(i)
	// 	}()
	// }
	// var a string 
	// v := reflect.ValueOf(a)
	// fmt.Println(v.Kind())
	// v.Call()
	ctx,cancel := context.WithCancel(context.Background())
	pool := routinepool.GetPool(ctx,10)

	for i:=0 ;i<100; i++ {
		j:=i
		f := func(){
			fmt.Println(j)
			time.Sleep(time.Second*2)	
		}

		go func(){
			pool.Add(f)
		}()
		
	}

	go func(){
		time.Sleep(time.Second*10)
		cancel()
	}()

	pool.Run()
	
}