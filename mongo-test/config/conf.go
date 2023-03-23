package config

var Config *Conf

type Conf struct {
	Mongo MongoConfig
}

type MongoConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"scheme"`
	Verbose  bool   `yaml:"verbose"`
	Pem      string `yaml:"pem"`
}
