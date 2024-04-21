package Database

import (
	"banks/Config"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

func InitDatabase(conf Config.Cfg) error {
	var my_str string = fmt.Sprintf("postgres://%s:%s@%s:%d/%s", conf.DB.User, conf.DB.Pwd, conf.DB.Host, conf.DB.Port, conf.DB.Database)
	// db, err = sqlx.Connect("postgres", mysql_conf.FormatDSN())
	Db, err := sqlx.Connect("pgx", my_str)
	Db.SetMaxOpenConns(Config.Config.DB.Max_Connections)
	Db.SetMaxIdleConns(Config.Config.DB.Max_Connections)
	Db.SetConnMaxLifetime(time.Millisecond * time.Duration(Config.Config.DB.Connection_LifeTime_Ms))
	return err
}
