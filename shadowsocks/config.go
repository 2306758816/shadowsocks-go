package shadowsocks

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Type         string    `json:"type"`
	Localaddr    string    `json:"localaddr"`
	Remoteaddr   string    `json:"remoteaddr"`
	Method       string    `json:"method"`
	Password     string    `json:"password"`
	Nonop        bool      `json:"nonop"`
	UdpRelay     bool `json:"udprelay"`
	UdpOverTCP bool `json:"udpovertcp"`
	Backend      *Config   `json:"backend"`
	Backends     []*Config `json:"backends"`
	Ivlen        int
}

func ReadConfig(path string) (configs []*Config, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &configs)
	if err != nil {
		var c Config
		err = json.Unmarshal(bytes, &c)
		if err == nil {
			configs = append(configs, &c)
		}
	}
	for _, c := range configs {
		CheckConfig(c)
	}
	return
}

func CheckConfig(c *Config) {
	if len(c.Password) == 0 {
		c.Password = defaultPassword
	}
	if len(c.Method) == 0 {
		c.Method = defaultMethod
	}
	if c.Ivlen == 0 {
		c.Ivlen = GetIvLen(c.Method)
	}
	if c.Backend != nil {
		c.Backends = append(c.Backends, c.Backend)
	}
	if len(c.Type) == 0 {
		if len(c.Localaddr) != 0 && len(c.Remoteaddr) != 0 {
			c.Type = "local"
		} else if len(c.Localaddr) != 0 {
			c.Type = "server"
		}
	}
	if c.UdpRelay && c.Type != "server" && c.Type != "local" {
		c.UdpRelay = false
	}
	if c.UdpOverTCP && c.Type != "server" && c.Type != "local" {
		c.UdpOverTCP = false
	}
	for _, v := range c.Backends {
		CheckConfig(v)
	}
}