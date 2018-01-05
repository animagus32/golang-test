package mlog 


import (
	"log"
	"os"
)

var (
    Logger *log.Logger
    logFile *os.File
)

func init(){
    fileName := "debug.log"  
    logFile,err  := os.Create(fileName)  
    // defer logFile.Close()  
    if err != nil {  
        log.Fatalln("open file error !")  
    }  
    Logger = log.New(logFile,"[Debug]",log.Llongfile)  
    Logger.Println("A debug message here")  
    Logger.SetPrefix("[Info]")  
    Logger.Println("A Info Message here ")  
    Logger.SetFlags(Logger.Flags() | log.LstdFlags)  
    Logger.Println("A different prefix...") 
}

func Close(){
    logFile.Close()
}