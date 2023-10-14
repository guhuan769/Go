package yaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Conf struct {
	IpPort             string
	StartSendTime      string
	SendMaxCountPerDay int
	Devices            []Device
	WarnFrequency      int
	sendfrequency      int
	Http               http
}

type http struct {
	Ip   string
	Port string
}

type Device struct {
	DevId string
	Nodes []Node
}

type Node struct {
	PkId     int    `yaml:"pkid"`
	BkId     int    `yaml:"bkid"`
	Index    int    `yaml:"index"`
	Minvalue int    `yaml:"minvalue"`
	Maxvalue int    `yaml:"maxvalue"`
	DataType string `yaml:"datatype"`
}

func GetConfig() {
	data, _ := ioutil.ReadFile("configs/config.yaml")
	c := Conf{}
	yaml.Unmarshal(data, &c)
	fmt.Printf("%+v", c)
}
