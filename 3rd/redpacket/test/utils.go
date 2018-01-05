package test

import (
	"net/http"
	"io/ioutil"
	"redpacket/controllers"
	"encoding/json"
	"net/url"
)
const  HOST string = "http://127.0.0.1:3000"

func login(user string ,passwd string) ( []byte, bool) {
	params := make(url.Values)
	params["username"] = []string{user}
	params["password"] = []string{passwd}
	resp,err := http.PostForm(HOST+"/api/v1/user/login",params)
	if err != nil {
		return nil,false
	}
	d ,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,false
	}
	return d,true
}

func getLoginData(user string ,passwd string) ( *controllers.UserLoginResponse, bool) {
	d,ok := login(user ,passwd )
	if ok == false{
		return nil,false
	}
	resp := controllers.UserLoginResponse{}
	err := json.Unmarshal(d, &resp)
	if err!= nil{
		return nil,false
	}
	if resp.Code != 0 {
		return nil,false
	}

	return &resp,true
}