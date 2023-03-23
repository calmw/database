package config

var Config *Conf

type Conf struct {
	Mysql MysqlConfig
}

type MysqlConfig struct {
	Address      string `toml:"address"`
	Port         string `toml:"port"`
	DbName       string `toml:"db_name"`
	UserName     string `toml:"user_name"`
	Password     string `toml:"password"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxLifeTime  int    `toml:"max_life_time"`
}
