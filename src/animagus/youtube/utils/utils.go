package utils
import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
)


func FinishNotify(msg string) error{
	hostname,_ := os.Hostname()
	data := map[string]interface{}{
		"msgtype":"text",
		"text" : map[string]string {
			"content" : fmt.Sprintf("Server %s finish job,%s",hostname,msg),
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