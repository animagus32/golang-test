package main 
import (
	"os"
)
func isDirExists(path string) bool {
    fi, err := os.Stat(path)
 
    if err != nil {
        return os.IsExist(err)
    } else {
        return fi.IsDir()
    }
 
    panic("not reached")
}

func createDirIfNotExist(path string) error {
	exist := isDirExists(path)

	if !exist {
		return os.Mkdir(path,os.ModePerm)
	}

	return nil
}