package main
import (
	"fmt"
	"os"
)

type I interface {
	Test() int 
}
type P int 

type sp *S
type S struct {
	
	P1 int 
	P2 int 
}


func (P) Test() int {
	return 111
}

func (S) Test() int {
	return 222
}


func main(){

	{
		filename := "text.txt"
		s,err := os.Stat(filename)
		if err!=nil {
			return 
		}
		fmt.Printf("%d",s.Size())
		f , err1 := os.OpenFile(filename,os.O_APPEND|os.O_RDWR,0600)
		if err1!= nil {
			return 
		}
		defer f.Close()
		f.Write([]byte("hasjdfkdkkdkk\n"))
		f.Sync()
	
		fmt.Print("end")
	}

	{
		s1 := S{P1:2}
		s := sp(&s1)
		fmt.Print(s.P1)
	}

}