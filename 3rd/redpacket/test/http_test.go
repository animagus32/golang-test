package test

import (
	"testing"
	"net/http"
	"io/ioutil"
	"redpacket/controllers"
	"encoding/json"
	"strconv"
	"net/url"
	"strings"
)

var secret string
var token string
var uid uint

func TestLogin(t *testing.T)  {
	d,ok := login("test7","123456")
	if ok == false{
		t.Error("http failed!")
		return
	}
	resp := controllers.UserLoginResponse{}
	err := json.Unmarshal(d, &resp)
	if err!= nil{
		t.Error(err.Error())
		return
	}
	if resp.Code != 0 {
		t.Error("login failed!")
		return
	}

}

func TestLoginWrongPasswd(t *testing.T)  {
	d,ok := login("test7","wrong pass")
	if ok == false{
		t.Error("http failed!")
		return
	}
	resp := controllers.UserLoginResponse{}
	err := json.Unmarshal(d, &resp)
	if err!= nil{
		t.Error(err.Error())
		return

	}
	if resp.Code != 1 {
		t.Error("login validation failed!")
	}
}

func TestDispatch(t *testing.T)  {
	loginData,ok := getLoginData("test7","123456")
	if ok == false{
		t.Error("http failed!")
		return
	}

	token = loginData.Token
	uid = loginData.Data.Id

	v := url.Values{}
	v.Set("uid", strconv.Itoa(int(loginData.Data.Id)))
	v.Set("amount","50")
	v.Set("num","4")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码

	client := &http.Client{}
	req, _ := http.NewRequest("POST" ,HOST+"/api/v1/redpacket/dispatch",body)
	req.Header.Set("uid",strconv.Itoa(int(loginData.Data.Id)))
	req.Header.Set("token",loginData.Token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	response,err := client.Do(req)
	defer response.Body.Close()
	if err != nil {
		t.Error(err.Error())
		return
	}

	rbody, _ := ioutil.ReadAll(response.Body)
	resp := controllers.RedPacketDispatchResponse{}
	err = json.Unmarshal(rbody, &resp)
	if err!= nil{
		t.Error(err.Error())
		return
	}
	if resp.Code != 0 {
		t.Error("code!=0")
		return
	}

	
	secret = resp.Data.Secret
}

func TestGrab(t *testing.T)  {

	v := url.Values{}
	v.Set("secret",secret)
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码

	client := &http.Client{}
	req, _ := http.NewRequest("POST" ,HOST+"/api/v1/redpacket/grab",body)
	req.Header.Set("uid",strconv.Itoa(int(uid)))
	req.Header.Set("token",token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	response,err := client.Do(req)
	defer response.Body.Close()
	if err != nil {
		t.Error(err.Error())
		return
	}

	rbody, _ := ioutil.ReadAll(response.Body)
	//fmt.Print(string(rbody))
	resp := controllers.GrabResponse{}
	err = json.Unmarshal(rbody, &resp)
	if err!= nil{
		t.Error(err.Error())
		return
	}
	if resp.Code != 0 {
		t.Error("code!=0")
		return
	}


}

func TestGrabFail(t *testing.T)  {
	v := url.Values{}
	v.Set("secret","wrong secret")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码

	client := &http.Client{}
	req, _ := http.NewRequest("POST" ,HOST+"/api/v1/redpacket/grab",body)
	req.Header.Set("uid",strconv.Itoa(int(uid)))
	req.Header.Set("token",token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	response,err := client.Do(req)
	defer response.Body.Close()
	if err != nil {
		t.Error(err.Error())
		return
	}

	rbody, _ := ioutil.ReadAll(response.Body)
	//fmt.Print(string(rbody))
	resp := controllers.GrabResponse{}
	err = json.Unmarshal(rbody, &resp)
	if err!= nil{
		t.Error(err.Error())
		return
	}
	if resp.Code == 0 {
		t.Error("code!=0")
		return
	}

}
//
//func TestRedpacketTimeout(t *testing.T)  {
//
//}
//

