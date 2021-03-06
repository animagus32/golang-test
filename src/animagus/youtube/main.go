package main

import (
	"animagus/youtube/utils"
	"bufio"
	"flag"
	"fmt"
	"github.com/willf/bloom"
	"io"
	"os"
	"strings"
	"time"
)

var filter *bloom.BloomFilter

func getKeys(fname string, c chan string,start string ) {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	buf := bufio.NewReader(f)
	skip := true
	if len(start) < 1 {
		skip = false
	}

	for {
		line, err := buf.ReadString('!')
		// line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err.Error())
				break
			}
		}
		line = strings.Join(strings.Split(line, "\n"), "")
		fmt.Println(line)

		if strings.TrimSpace(start) == strings.TrimSpace(line) {
			skip = false
		}

		if len(line) < 5 || skip {
			continue
		}



		c <- line
		// break
	}
	// time.Sleep(time.Second*10)
	// close(c)

}

//todo 带time的log
func main() {
	seedFile := flag.String("seed", "ids.txt", " seed file ")
	targetDir := flag.String("target-dir", "/data/youtube/video", " target dir ")

	flag.Parse()
	fmt.Println(*seedFile,*targetDir)
	// return 
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			utils.FinishNotify("youtube downloading panic")
		} else {
			utils.FinishNotify("youtube downloading success")
		}

	}()

	done := make(chan bool)
	cKey := make(chan string, 100)
	cKeyDownloaded := make(chan string, 100)
	// cLimit := make(chan bool,2)

	filter := getBloomFilter(uint(500000))
	lastDownloaded,err := findLastDownloaded()
	if err != nil {
		lastDownloaded = ""
	}
	go getKeys(*seedFile, cKey,lastDownloaded)

	go func() {
		for {
			// key := "mLu5xsuGQGwYso0Fa5rwakqPIJZlFq1sEWw1KrJTWau4_x!"
			key, ok := <-cKey
			if !ok {
				break
			}
			if filter.TestString(key) {
				fmt.Println("Already downloaded : ", key)
			} else {
				// cLimit <- true
				// go func(){
					// defer func(){
					// 	<- cLimit
					// }()
					success, str := download(key,*targetDir)
					if !success {
						fmt.Println("Download error ", str)
					} else {
						// time.Sleep(time.Second*2)
						cKeyDownloaded <- key
						filter.AddString(key)
						fmt.Println("Add key : ", key)
					}
				// }()

			}
			// break
		}
		done <- true
	}()

	go recordDownloaded(cKeyDownloaded)

	<-done
}

func getBloomFilter(n uint) *bloom.BloomFilter {
	bfFileName := "bf.data"

	bfFile, err := os.OpenFile(bfFileName, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		panic(err)
	}
	defer bfFile.Close()

	filter = bloom.New(20*n, 5) // load of 20, 5 keys
	filter.ReadFrom(bfFile)

	go func() {
		ticker := time.Tick(time.Second * 60)
		for {
			<-ticker
			f, err := os.OpenFile(bfFileName, os.O_RDWR, 0660)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			filter.WriteTo(f)

		}
	}()

	return filter
}
