package main

import (
	"bufio"
	"fmt"
	"github.com/willf/bloom"
	"io"
	"os"
	"os/exec"
	"strings"
)

func download(key string) error {
	// key := "mLu5xsuGQGwYso0Fa5rwakqPIJZlFq1sEWw1KrJTWau4_x"
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", key)
	cmd := exec.Command("./youtube-dl", "-o", "/tmp/%(title)s.%(ext)s", url)
	fmt.Println(url)
	//todo 标准出错，异常处理
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	cmd.Start()

	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}
	cmd.Wait()

	return nil
}

func getKeys(fname string, c chan string) {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	buf := bufio.NewReader(f)
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
		// fmt.Println(line)
		line = strings.Join(strings.Split(line, "\n"), "")
		c <- line
		break
	}

}

func main() {
	done := make(chan bool)
	cKey := make(chan string, 100)

	filter := getBloomFilter(uint(500000))

	go getKeys("ids.txt", cKey)

	go func() {
		for {
			// key := "mLu5xsuGQGwYso0Fa5rwakqPIJZlFq1sEWw1KrJTWau4_x!"
			key := <-cKey

			if filter.TestString(key) {
				fmt.Println("Already download : ", key)
			} else {
				err := download(key)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					filter.AddString(key)
				}
			}
			// break
		}
		done <- true
	}()

	<-done
}


func getBloomFilter(n uint) *bloom.BloomFilter{

	//load from file 
	// bloom.From()
	filter := bloom.New(20*n, 5) // load of 20, 5 keys
	return filter
}