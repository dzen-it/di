package configs

import "github.com/kelseyhightower/envconfig"

// Config is default config
var Config = Variables{
	Mode:    "cache",
	HTTPS:   false,
	Address: ":8080",
	Level:   "debug",
	Redis:   "localhost:6379",
	RedisDB: 0,
}

type Variables struct {
	Mode     string
	Address  string
	Level    string
	HTTPS    bool
	Certfile string
	Keyfile  string
	Redis    string
	RedisDB  int
}

func init() {
	if err := envconfig.Process("di", &Config); err != nil {
		panic(err)
	}
}
