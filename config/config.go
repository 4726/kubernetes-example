package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB   Redis
	Addr string
}

type Redis struct {
	Addr, Password string
	DB             int
}

//FromFile returns a new Config object from path
func FromFile(path string) (Config, error) {
	var conf Config
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, err
	}
	err = yaml.Unmarshal(content, &conf)
	return conf, err
}
