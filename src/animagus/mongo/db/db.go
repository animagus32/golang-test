package db 
  import (
	"fmt"
	// "labix.org/v2/mgo"
	"gopkg.in/mgo.v2"
	// "labix.org/v2/mgo/bson"
)

func ConnectDB(host string,port uint) *mgo.Session{
	session, err := mgo.Dial(fmt.Sprintf("%s:%d",host,port))  //连接数据库
	if err != nil {
		panic(err)
	}
	// defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	return session
}
