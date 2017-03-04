package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Client   string `json:"client"`
	Server   string `json:"server"`
	Method   string `json:"method"`
	Password string `json:"password"`
	ivlen    int
}

func readConfig(path string) (configs []Config, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &configs)
	if err != nil {
		var c Config
		err = json.Unmarshal(bytes, &c)
		if err == nil {
			if len(c.Password) == 0 {
				c.Password = defaultPassword
			}
			if len(c.Method) == 0 {
				c.Method = defaultMethod
			}
			configs = append(configs, c)
		}
	}
	return
}
