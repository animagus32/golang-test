package controllers


import (
	"net/http"
	"github.com/jinzhu/gorm"
	"redpacket/utils"
	"strconv"

	"encoding/json"
	"redpacket/models"
	"math"
	"fmt"
)

func Dispatch(res http.ResponseWriter,req *http.Request, db gorm.DB,tokenMapper utils.TokenMapper,packetRetriever *utils.RedpacketRetreiver){
	uid,_ := strconv.Atoi(req.FormValue("uid"))
	//fmt.Printf("..3%s3..",req.FormValue("uid"))
	if(!isAuthorized(req,tokenMapper) || !userCheck(req,uid)){
		error(res,1,"Not authorized!")
		return	;
	}
	amount,_ := strconv.ParseFloat(req.PostFormValue("amount"),64)
	num,_ := strconv.Atoi(req.PostFormValue("num"))

	randList,ok := getRandList(amount,num)
	if ok == false {
		error(res,1,"Parameter error!")
		return;
	}

	redpacketM := models.Redpacket{}
	redpacketM.Amount = math.Trunc(amount*100)/100
	redpacketM.UserId = uid
	redpacketM.Secret = GetRandomString(8)
	db.Create(&redpacketM)

	mList := make([]models.RedpacketDetail,num)


	for i,v := range randList {
		//fmt.Printf(" %v ",v)
		mList[i].UserId = 0
		mList[i].Amount = v
		mList[i].RedpacketId = redpacketM.Id
		db.Create(&mList[i])
	}

	//packetRetriever.List = append(packetRetriever.List,utils.RedpacketItem{Id:redpacketM.Id,Timeout:redpacketM.CreatedAt.Unix()})
	//fmt.Print(packetRetriever.List)
	packetRetriever.Insert(redpacketM.Id,redpacketM.CreatedAt.Unix())

	response := RedPacketDispatchResponse{}
	response.Data.List = mList
	response.Data.Secret = redpacketM.Secret
	ret,err := json.Marshal(response)
	if err!= nil {
		panic(err)
	}
	writeJson(res,ret)

}

func Grab(res http.ResponseWriter,req *http.Request, db gorm.DB,tokenMapper utils.TokenMapper)  {
	uid,_ := strconv.Atoi(req.Header.Get("uid"))
	secret := req.PostFormValue("secret")

	if(!isAuthorized(req,tokenMapper)){
		error(res,1,"Not authorized!")
		return	;
	}
	redPacketM := models.Redpacket{}
	db.Where("secret=?",secret).Where("status=?",0).First(&redPacketM)

	if redPacketM.Id == 0{
		error(res,1,"Sorry, red packet is empty")
		return ;
	}

	redpacketDetailM := models.RedpacketDetail{}

	//是否抢过
	db.Where("redpacket_id=?",redPacketM.Id).Where("user_id=?",uid).First(&redpacketDetailM)
	if redpacketDetailM.Id != 0 {
		error(res, 1, "you have alread grab this packet")
		return;
	}

	//todo 要加锁，不过没时间:)
	db.Where("redpacket_id=?",redPacketM.Id).Where("user_id=?",0).First(&redpacketDetailM)
	fmt.Printf("%v",redpacketDetailM)
	if redpacketDetailM.Id == 0{
		if redPacketM.Status == 0 {
			//fmt.Printf("....%v...",redPacketM)
			redPacketM.Status = 1
			db.Save(&redPacketM)
		}
		error(res,1,"Sorry, red packet is empty")
		return
	}
	redpacketDetailM.UserId = uint(uid)
	db.Model(&redpacketDetailM).Where("id=?",redpacketDetailM.Id).Where("user_id=?",0).Update("user_id",uid)

	//加入余额，同理，，锁
	user := models.User{}
	db.Where("id=?",uid).First(&user)
	user.Balance += redpacketDetailM.Amount
	db.Save(&user)

	//查看红包是否已抢完
	m := models.RedpacketDetail{}
	db.Where("redpacket_id=?",redPacketM.Id).Where("user_id=?",0).First(&m)
	if m.Id == 0 {
		redPacketM.Status = 1
		db.Save(&redPacketM)
	}

	response := GrabResponse{}
	response.Data = redpacketDetailM
	ret,err := json.Marshal(response)
	if err!= nil {
		panic(err)
	}
	writeJson(res,ret)
}



func Index(res http.ResponseWriter,req *http.Request, db gorm.DB,tokenMapper utils.TokenMapper)  {
	uid,_ := strconv.Atoi(req.FormValue("uid"))

	if(!isAuthorized(req,tokenMapper) || !userCheck(req,uid)){
		error(res,1,"Not authorized!")
		return	;
	}

	redpackets := []models.RedpacketDetail{}
	db.Where("user_id=?",uid).Find(&redpackets)

	response := IndexResponse{}
	response.Data = redpackets
	ret,err := json.Marshal(response)
	if err!= nil {
		panic(err)
	}
	writeJson(res,ret)
}