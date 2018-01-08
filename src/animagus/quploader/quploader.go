package main 

import (
	"fmt"
	"io"
	"os"
	"flag"
	"time"
	"math"
	"animagus/quploader/utils"

)

// var root string 
// var filename string 
// func init(){
// 	flag.StringVar(&string,"")
// }

func main(){	
	flag.Parse()
	root := flag.Arg(0)
	filename := flag.Arg(1)
	
	filelist := utils.GetFileList(root)

	utils.OutputFileList(filelist,filename)

	defer utils.FinishNotify("upload done")

	logFile,err := os.OpenFile(fmt.Sprintf("log/time_spent_%d.txt",time.Now().Unix()),os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
	if err != nil {
		return 
	}
	defer logFile.Close()

	tStart := time.Now().Unix()

	for _,v := range(filelist) {
		fmt.Printf("%v\n",v)
		start := time.Now().Unix()
		utils.StartUpload(root,v.Path)
		duration := time.Now().Unix()-start
		
		io.WriteString(logFile,fmt.Sprintf("%s : %d : %d\n",v.Path,v.Size/int64(math.Pow(2,20)),duration))
	}

	tDuation := time.Now().Unix()-tStart
	io.WriteString(logFile,fmt.Sprintf("\n\n\ntotal time spend: %d\n",tDuation))

}