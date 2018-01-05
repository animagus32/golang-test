package main

import (
	"fmt"
	// "labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
	"animagus/mongo/db"
	mlog "animagus/mongo/log"
	"golang.org/x/net/context"
	"reflect"
	"runtime"
	"sync"
	"time"
		"net/http"
	_ "net/http/pprof"
)

//todo 增加context，日志，运用go routine

var logger = mlog.Logger

func main() {

	defer mlog.Close()
	defer func() {
		if err := recover(); err != nil {
			// logger.Println(typeof(err))
			logger.Println(err)
			fmt.Println(err)
		}
	}()

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// runtime.GOMAXPROCS(runtime.NumCPU())
	climit := make(chan bool, 10000) //limit the max go routine
	var (
		config *Config
		err    error
	)
	var wg sync.WaitGroup
	start := time.Now()
	if config, err = ReadConfig("./conf.toml"); err != nil {
		panic(err)
	}

	fmt.Printf("host:%v\n", config.Mongo.Host)

	session := db.ConnectDB(config.Mongo.Host, config.Mongo.Port)
	defer session.Close()

	db := session.DB(config.Mongo.Db)
	db.C("movie").DropCollection()

	handler := func(ctx context.Context,info *MovieInfo) {
		climit <- true
		done := make(chan struct{})
		wg.Add(1)
		go func(info *MovieInfo) {
			// defer wg.Done()
			
			fmt.Print(".")
			err := db.C("movie").Insert(info)
			if err != nil {
				fmt.Println("error:", err.Error())
			}
			<-climit
			done <- struct{}{}
		}(info)
		
		select{
		case <- ctx.Done():
			wg.Done()
		case <- done:
			wg.Done()
		}
	}

	ctx := context.WithValue(context.Background(), "log", logger)

	ctx,_ = context.WithTimeout(ctx,time.Second*2)

	ImportVideoInfo(ctx, "filelist", handler)

	tick := time.Tick(time.Second * 2)
	go func() {
		for {
			<-tick
			fmt.Println(runtime.NumGoroutine())
		}
	}()

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf(" duartion : %v \n", duration)
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
