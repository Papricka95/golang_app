package Config

import "github.com/jinzhu/configor"

type Cfg struct {
	DB struct {
		Host                   string `default:"localhost"`
		Port                   uint   `default:"3306"`
		User                   string `required:"true"`
		Pwd                    string `required:"true"`
		Database               string `required:"true"`
		Max_Connections        int    `required:"true"`
		Connection_LifeTime_Ms int    `default:"3600"` // 1 min
	}
	Tables struct {
		Transact string `required:"true"`
	}
}

var Config = Cfg{}

func LoadConfig(path string) error {
	return configor.Load(&Config, path)
}
