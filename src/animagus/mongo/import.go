package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type MovieInfo struct {
	Title  string
	Domain string
	Ext    string
	Path   string
	HdNum  int
	//todo timestamp
}

// var VIDEO_SUFFIX = []string{"mp4","avi","rmvb","mkv"}

func ImportVideoInfo(ctx context.Context, dir string, handler func(context.Context,*MovieInfo)) error {
	var wg sync.WaitGroup
	var count int = 0
	logger, ok := ctx.Value("log").(*log.Logger)
	if !ok {
		panic(" log converstion error!")
	}
	logger.Println(" log test ! ")
	dir_list, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	reg := regexp.MustCompile(`movie_(\d+).txt`)

	suffixReg := regexp.MustCompile(`.*\.(mp4|mkv|rmvb|avi)$`)

	for i, v := range dir_list {
		wg.Add(1)
		go func(i int, v os.FileInfo) error {
			defer wg.Done()
			fmt.Println(i, "=", v.Name())
			match := reg.FindStringSubmatch(v.Name())
			if len(match) > 1 {
				hdNum, _ := strconv.Atoi(match[1])

				f, err := os.Open(path.Join(dir, v.Name()))
				if err != nil {
					return err
				}
				buf := bufio.NewReader(f)
				for {
					line, err := buf.ReadString('\n')
					line = strings.TrimSpace(line)
					if err != nil {
						if err == io.EOF {
							break
						} else {
							fmt.Println(err.Error())
							break
						}
					}
					match := suffixReg.FindStringSubmatch(line)
					if len(match) > 1 {
						sp := strings.Split(line, "/")
						domain := sp[0]
						title := sp[len(sp)-1]
						// fmt.Println(domain," : ",title)

						info := MovieInfo{
							Title:  title,
							Domain: domain,
							Path:   line,
							Ext:    match[1],
							HdNum:  hdNum,
						}
						handler(ctx,&info)

						count++
					}
				}
			}
			fmt.Println(" ")
			return nil
		}(i, v)

	}
	// break
	wg.Wait()
	fmt.Println("count: ", count)
	return nil
}
