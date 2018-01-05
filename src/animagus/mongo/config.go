package main

import (
	"github.com/BurntSushi/toml"
	// "os"
	// "io/ioutil"
)

type Config struct {
	Mongo mongoConfig
}

type mongoConfig struct {
	Host string
	Port uint
	Db   string
}

func ReadConfig(fname string) (config *Config, err error) {
	var (
	// fp *os.File
	// fcontent []byte
	)
	config = new(Config)

	// if fp,err = os.Open(fname); err!=nil {
	// 	return config,err
	// }
	// if fcontent,err = ioutil.ReadAll(fp); err!=nil {
	// 	return
	// }

	// if err = toml.Unmarshal(fcontent,config);err!=nil {
	if _, err := toml.DecodeFile(fname, config); err != nil {
		return config, err
	}
	return
}
