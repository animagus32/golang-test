package main


import (
	"fmt"
	"io"
	"os"
	"flag"
	"time"
	"math"
	"animagus/quploader/utils"
	"github.com/willf/bloom"
	. "animagus/youtube/utils"
	"path"
	"bufio"

)

func UploadDir(root string,filename string){
	filelist := utils.GetFileList(root)

	utils.OutputFileList(filelist,filename)

	filter := genBloomFilterFromText(1000000,root)

	defer utils.FinishNotify("upload youtube done")

	logFile,err := os.OpenFile(fmt.Sprintf("log/time_spent_%d.txt",time.Now().Unix()),os.O_CREATE|os.O_APPEND|os.O_RDWR,0666)
	if err != nil {
		return 
	}
	defer logFile.Close()

	tStart := time.Now().Unix()

	for _,v := range(filelist) {
		fname := path.Base(v.Path)
		if filter.TestString(fname) {
			//skip
			fmt.Printf("skip %s ,already downloaded\n",v.Path)
			continue
		}

		if fname[0] == '.' {
			// fmt.Printf("skip %\n",v.Path)
			continue
		}

		fmt.Printf("%s\n",v.Path)		
		start := time.Now().Unix()
		utils.StartUpload(root,v.Path)
		duration := time.Now().Unix()-start
		
		io.WriteString(logFile,fmt.Sprintf("%s : %d : %d\n",v.Path,v.Size/int64(math.Pow(2,20)),duration))
	}

	tDuation := time.Now().Unix()-tStart
	io.WriteString(logFile,fmt.Sprintf("\n\n\ntotal time spend: %d\n",tDuation))
}

func main(){	
	flag.Parse()
	root := flag.Arg(0)
	filename := flag.Arg(1)
	
	for {
		UploadDir(root,filename)
		tick := time.Tick(time.Hour)
		<- tick
	}
}

func genBloomFilterFromText(n uint,root string) (filter *bloom.BloomFilter){
	defaultFile := path.Join(root,".success")
	exist := IsFileExists(defaultFile)
	
	filter = bloom.New(20*n, 5) // load of 20, 5 keys
	if exist {
		f,err := os.Open(defaultFile)
		if err != nil {
			fmt.Println("open .success file error")
			return 
		}

		br := bufio.NewReader(f) 
		defer f.Close()
		for {
			line,_,err := br.ReadLine()
			if err == io.EOF {
				break
			}
			filter.Add(line)
		}
		return 
	}
	return 
}
