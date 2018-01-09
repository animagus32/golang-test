package utils


import (
	"os/exec"
	"fmt"
	"bufio"
	"io"
	"os"
	"path/filepath"
	"net/http"
	"encoding/json"
	"strings"

)

func StartUpload(root string,filepath string){
	var err error
	cmd := exec.Command("bash","exec.sh",root,filepath)

	stdout,err := cmd.StdoutPipe()
	if err!=nil {
		fmt.Println(err.Error())
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
}

type FileInfo struct{
	Path string
	Size int64
}

func GetFileList(path string) (output []FileInfo){
	err := filepath.Walk(path,func(p string,f os.FileInfo,err error) error {
		if err != nil || f == nil {
		// if (  ) {
		// 	// fmt.Printf("%v",err)
			return filepath.SkipDir
		}
		if f.IsDir() {return nil}
		// fmt.Println(f.Size())
		output = append(output,FileInfo{Path:p[len(path)+1:],Size:f.Size()})
		return nil
	})

	if err != nil {
                fmt.Printf("filepath.Walk() returned %v\n", err)
    }
	return 
}


func OutputFileList(filelist []FileInfo,filename string){
	f,err := os.OpenFile(fmt.Sprintf("filelist/%s.txt",filename),os.O_CREATE|os.O_RDWR,0666)
	if err != nil {
		return 
	}
	defer f.Close()

	for _,v := range(filelist) {	
		fmt.Println(v.Path)	
		// s := strings.Split(v.Path,".")
		// fmt.Println(s[len(s)-1])
		io.WriteString(f,v.Path+"\n")
	}

}

func FinishNotify(msg string) error{
	hostname,_ := os.Hostname()
	data := map[string]interface{}{
		"msgtype":"text",
		"text" : map[string]string {
			"content" : fmt.Sprintf("Server %s finish job, %s",hostname,msg),
		},
	}
	
	dataStr,err := json.Marshal(data)
	if err!= nil {
		return err 
	}
	fmt.Println(string(dataStr))
	
	client := &http.Client{}
	req,err := http.NewRequest("POST","https://oapi.dingtalk.com/robot/send?access_token=c67ea86532866331a493515a5d93a51136761256f61e120512bd3867197ee1b1",strings.NewReader(string(dataStr)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type","application/json")
	client.Do(req)

	return nil
}