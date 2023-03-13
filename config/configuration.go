package config

import (
	"fmt"
	"generate_btc/util"
	"log"
	"os"
)

type Configuration struct {
	Server Server   `yaml:"server"`
	BTC    BTCParam `yaml:"btc"`
}

type Server struct {
	Port   uint   `yaml:"port"`
	Active string `yaml:"active"`
}

type BTCParam struct {
	ENV string `yaml:"env"`
}

func covertConfiguration() *Configuration {
	c := new(Configuration)
	err := util.ReadYamlConfig("conf/app.yaml", c)
	if err != nil {
		log.Panic(err.Error())
	}

	if c.Server.Active != "" {
		cc := convertActiveConfiguration(c.Server.Active)
		if cc != nil {
			//覆盖配置
			err := util.StructCopy(c, cc)
			if err != nil {
				log.Panic(err.Error())
			}
		}
	}
	return c
}

func convertActiveConfiguration(env string) *Configuration {
	path := fmt.Sprintf("conf/app-%s.yaml", env)
	if !fileExist(path) {
		return nil
	}
	c := new(Configuration)
	err := util.ReadYamlConfig(path, c)
	if err != nil {
		log.Panic(err.Error())
	}
	return c
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

var C = covertConfiguration()
