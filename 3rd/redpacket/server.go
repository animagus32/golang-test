package main

import "github.com/go-martini/martini"
import (
	"redpacket/controllers"
	"github.com/jinzhu/gorm"
 	_ "github.com/go-sql-driver/mysql"
	"redpacket/models"
	"redpacket/utils"
	"time"
	"fmt"
)

func initDB()  *gorm.DB{
	db, err := gorm.Open("mysql","root:333222@tcp(127.0.0.1:3306)/redpacket?charset=utf8")
	if err != nil {
		//fmt.Printf("%v",err.Error())
		panic(err)
	}
	db.LogMode(true)

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Redpacket{})
	db.AutoMigrate(&models.RedpacketDetail{})

	return db
}

var tokenMapper = make(utils.TokenMapper)	//登录token，为了方便用内存
var redpacketRetreiver utils.RedpacketRetreiver	//红包列表，可以放在消息队列或redis中，为了方便用内存

var db *gorm.DB

func main(){

	db = initDB()
	defer db.Close()


	m := martini.Classic()
	m.Map(*db)
	m.Map(tokenMapper)
	m.Map(&redpacketRetreiver)
	//m.Use(func(res http.ResponseWriter, req *http.Request) {
	//	if req.Header.Get("X-API-KEY") != "test" {
	//		res.WriteHeader(http.StatusUnauthorized)
	//	}
	//})

	m.Get("/",func() string{
		return "Hello World!"
	});

	m.Group("/api/v1/user",func(r martini.Router){
		r.Post("/login",controllers.Login)
		r.Post("/register",controllers.Register)
		r.Get("/:uid/balance",controllers.UserBalance)
	});

	m.Group("/api/v1/redpacket",func(r martini.Router){
		r.Post("/dispatch",controllers.Dispatch)
		r.Post("/grab",controllers.Grab)
		r.Get("/list",controllers.Index)
	});

	retrieveRedpacketManager()
	m.Run()
}

func retrieveRedpacketManager(){

	go func(){
		for {
			if redpacketRetreiver.List.Len() == 0 {
				time.Sleep(10)
				continue
			}
			item := redpacketRetreiver.List.Front()
			redpacketRetreiver.List.Remove(item)

			pItem,ok := item.Value.(utils.RedpacketItem)
			if !ok {
				continue
			}
			span := pItem.Timeout - time.Now().Unix()
			fmt.Printf("%v",span)
			t := time.After(time.Second * time.Duration(span))
			select {
			case <-t:
				fmt.Printf(" recieved!\n ")
				go retrieveRedpacketHandler(pItem.Id)
			}
		}
	}()

}

func retrieveRedpacketHandler(id uint)  {
	redpacketM := models.Redpacket{}
	db.Model(&redpacketM).Where("id=?",id).Where("status=?",0).Update("status",2)
}
