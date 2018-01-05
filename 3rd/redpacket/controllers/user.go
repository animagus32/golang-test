package controllers

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"redpacket/models"
	"encoding/json"
	"redpacket/utils"
	"fmt"
	"github.com/go-martini/martini"
)

func Login(res http.ResponseWriter,req *http.Request, db gorm.DB,tokenMapper utils.TokenMapper)  {
	name := req.PostFormValue("username")
	password := req.PostFormValue("password")
	user := models.User{}

	db.Where(" username = ?",name).First(&user)
	if user.Id == 0 {
		error(res,1,"user not exist!")
		return;
	}
	if user.Password != encrypt(password+user.Salt) {
		error(res,1,"passwor is not correct!")
		return;
	}

	token := utils.Token{}
	token.Generate()

	UserResponse := UserLoginResponse{}
	UserResponse.Data = user
	UserResponse.Token = token.Token

	data,err := json.Marshal(UserResponse)
	if err!= nil {
		panic(err)
	}

	tokenMapper[user.Id] = token

	fmt.Printf("%v",token.Token)

	writeJson(res,data)
}

func Register(res http.ResponseWriter,req *http.Request, db gorm.DB)  {

	name := req.PostFormValue("username")
	password := req.PostFormValue("password")

	user := models.User{}

	db.Where(" username=?",name).First(&user)
	if user.Id != 0 {
		error(res,1,"User already exist")
		return;
	}
	user.Username = name
	user.Salt = getRand()
	user.Password = encrypt(password+user.Salt)

	db.Create(&user)

	UserResponse := UserResponse{}
	UserResponse.Data = user
	data,err := json.Marshal(UserResponse)
	if err!= nil {
		panic(err)
	}
	writeJson(res,data)
}

func UserBalance(res http.ResponseWriter,req *http.Request,params martini.Params,tokenMapper utils.TokenMapper, db gorm.DB)  {
	if(!isAuthorized(req,tokenMapper)){
		error(res,1,"Not authorized!")
		return;
	}
	uid := params["uid"]

	user := models.User{}

	db.Where(" id=?",uid).First(&user)
	if user.Id == 0 {
		error(res,1,"user not exist!")
		return;
	}

	data := UserBalanceResponse{}
	data.Data.Balance = user.Balance
	ret,err := json.Marshal(data)
	if err!= nil {
		panic(err)
	}
	writeJson(res,ret)
}