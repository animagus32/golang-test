package main 

import (
	"testing"
	// "os"
	// "io"
	"net/http"
	"fmt"
	"encoding/json"
	"strings"
	"io/ioutil"
)

// func TestHttp(t *testing.T){
// 	client := &http.Client{}

// 	url := "http://www.baidu.com"

// 	req,err := http.NewRequest("GET",url,nil)

// 	if err!=nil {
// 		panic(err)
// 	}

// 	response,_ := client.Do(req)

// 	stdout := os.Stdout
// 	_, err = io.Copy(stdout, response.Body)
   
//    //返回的状态码
// 	status := response.StatusCode

// 	fmt.Println(status)
// }

func TestHttpPost(t *testing.T){
	data := map[string]interface{}{
		"msgtype":"text",
		"text" : map[string]string {
			"content" : " test ",
		},
	}
	
	dataStr,err := json.Marshal(data)
	if err!= nil {
		panic(err)
	}
	fmt.Println(string(dataStr))
	
	client := &http.Client{}
	req,err := http.NewRequest("POST","https://oapi.dingtalk.com/robot/send?access_token=c67ea86532866331a493515a5d93a51136761256f61e120512bd3867197ee1b1",strings.NewReader(string(dataStr)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type","application/json")
	resp, err := client.Do(req)
 
    defer resp.Body.Close()
 
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
 
    fmt.Println(string(body))

}