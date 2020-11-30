package conf

import (
	"sync"
)

type Config struct {
	Language string
	Token    string
	Super    string
	RedisPre string
	Host     string
	Routes   []string
}

var (
	Cfg     Config
	mutex   sync.Mutex
	declare sync.Once
)

func Set(cfg Config) {
	mutex.Lock()
	Cfg.RedisPre = setDefault(cfg.RedisPre, "", "go.xiet16.pmsweb.redis")
	Cfg.Language = setDefault(cfg.Language, "", "cn")
	Cfg.Token = setDefault(cfg.Token, "", "token")
	Cfg.Super = setDefault(cfg.Super, "", "xietie")
	Cfg.Host = setDefault(cfg.Host, "", "http://localhost:8899")
	Cfg.Routes = cfg.Routes
	mutex.Unlock()
}

func setDefault(value, def, defValue string) string {
	if value == def {
		return defValue
	}

	return value
}
