package controllers

import (
	"net/http"
	"time"
	"strconv"
	"math/rand"
	"crypto/sha1"
	"io"
	"fmt"
	"redpacket/models"
	"encoding/json"
	"redpacket/utils"
	"math"
)

type CommonResponse struct {
	Code int 	`json:"code"`
	Msg string	`json:"msg"`
}

type UserLoginResponse struct {
	CommonResponse
	Token string `json:"token"`
	Data 	models.User `json:"data"`
}

type UserResponse struct {
	CommonResponse
	Data 	models.User `json:"data"`
}

type UserBalanceResponse struct {
	CommonResponse
	Data struct{
		Balance float64 `json:"balance"`
	     }
}

type RedPacketDispatchResponse struct {
	CommonResponse
	Data struct{
		List     []models.RedpacketDetail `json:"list"`
		Secret	string `json:"secret"`
	     }
}

type GrabResponse struct {
	CommonResponse
	Data models.RedpacketDetail `json:"data"`
}

type IndexResponse struct {
	CommonResponse
	Data []models.RedpacketDetail `json:"data"`
}
//func success(res http.ResponseWriter,data interface{})  {
//	data.Code = 0
//	data.Msg = "ok"
//	data,err := json.Marshal(data)
//	if()
//	writeJson(res,data)
//}
func error(res http.ResponseWriter,code int,msg string)  {
	ret := CommonResponse{Code:code,Msg:msg}
	data,err := json.Marshal(ret)
	if err!= nil {
		panic(err)
	}
	writeJson(res,data)
}

func writeJson(res http.ResponseWriter,data []byte){
	res.Header().Set("Content-Type","application/json")
	res.Write(data);
}

func getRand() string {
	rand.Seed(time.Now().Unix())
	str := strconv.Itoa(rand.Int())
	return str
}

func encrypt(data string) string {
	t := sha1.New();
	io.WriteString(t,data);
	return fmt.Sprintf("%x",t.Sum(nil));
}

func checkPassword(target string,password string,salt string) bool {
	if target != encrypt(password+salt) {
		return false
	}else{
		return true
	}
}

func isAuthorized(req *http.Request,tokenMapper utils.TokenMapper) bool {
	token := req.Header.Get("token")
	uid,_ := strconv.Atoi(req.Header.Get("uid"))

	if t,ok := tokenMapper[uint(uid)];ok==false || t.Validate(token) == false{
		return false;
	}
	return true;
}

func userCheck(req *http.Request,uid int) bool{
	if id,_ := strconv.Atoi(req.Header.Get("uid")); id == uid {
		return true;
	}else{
		return false;
	}

}

func getRandList(amount float64,num int) ([]float64,bool) {
	amount = math.Trunc(amount*100)

	if(amount <= 0 || num < 1 || int(amount)< num ){
		return nil,false;
	}

	randList := make([]float64,num)

	left  := int(amount)
	i := 0

	thres := int(amount/3)

	for i < num-1 {
		r := rand.Intn(left)
		if r > thres {
			r = rand.Intn(thres)
		}
		//fmt.Printf(" %v ",r)
		if r!= 0 {
			if (left - r) >= num-i-1 {
				randList[i] = float64(r)/100
				i++
				left = left - r

				if left == i {
					for i < num-1 {
						randList[i] = 0.01
						i++
					}
					break
				}
			}
		}

	}

	randList[i] = float64(left)/100

	return randList,true
}

func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}