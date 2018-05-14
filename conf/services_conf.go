package conf

import (
	"flag"
	"os"
	"io/ioutil"
	"github.com/naoina/toml"
)

type ServiceConf struct {
	Name string
	Server *ServerConf
	Wallet *WalletConf
	Redis *RedisConf
	Mysql *MysqlConf
}

type ServerConf struct {
	Host         string
	Port         int
	Endpoints []string
	Timeout int
}

type MysqlConf struct {
	User string
	Password string
	Host string
	Port int
}

type RedisConf struct {
	Host string
	Password string
	Port int
	Timeout int
}


type WalletConf struct {
	CaCertPath string
	SnsCertPath string
	SnsKeyPath string
}

var (
	rootConfigPath = flag.String("c", "../../conf", "service root config path")
)

func LoadServiceConf(name string) (conf *ServiceConf, err error) {
	flag.Parse()
	path := *rootConfigPath + "/" + name + ".conf"
	return conf.loadconfig(path, name)
}

func (conf *ServiceConf) loadconfig(path, name string) (*ServiceConf, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var buf []byte
	if buf, err = ioutil.ReadAll(file); err != nil {
		return nil, err
	}

	conf = &ServiceConf{}
	if err = toml.Unmarshal(buf, conf); err != nil {
		return nil, err
	}

	return conf, err
}