package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"os"
	"time"
)

func download(key string) (bool,string) {
	// key := "mLu5xsuGQGwYso0Fa5rwakqPIJZlFq1sEWw1KrJTWau4_x"
	//文件夹名
	dirname := "/data/youtube/"+time.Now().Format("2006010215")
	err := createDirIfNotExist(dirname)
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", key)
	cmd := exec.Command("./youtube-dl", "-o", dirname+"/%(title)s.%(ext)s", url)
	fmt.Println(url)
	//todo 标准出错，异常处理
	stdout, err := cmd.StdoutPipe()
	defer stdout.Close()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	// stderr,err := cmd.Stderr
	cmd.Start()
	cmd.Stderr = os.Stdout
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}
	cmd.Wait()
	success := cmd.ProcessState.Success()
	str := cmd.ProcessState.String()
	fmt.Println(success)
	return success,str
}


func recordDownloaded(c chan string) {
	f,err := os.OpenFile("downloaded.txt",os.O_CREATE|os.O_RDWR,0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for key := range c {
		f.WriteString(key+"\n")
	}

}