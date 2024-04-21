package Database

import (
	"banks/Config"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

var Db *sqlx.DB

func InitDatabase(conf Config.Cfg) error {
	var my_str string = fmt.Sprintf("postgres://%s:%s@%s:%d/%s", conf.DB.User, conf.DB.Pwd, conf.DB.Host, conf.DB.Port, conf.DB.Database)
	var err error
	Db, err = sqlx.Connect("pgx", my_str)
	Db.SetMaxOpenConns(Config.Config.DB.Max_Connections)
	Db.SetMaxIdleConns(Config.Config.DB.Max_Connections)
	Db.SetConnMaxLifetime(time.Millisecond * time.Duration(Config.Config.DB.Connection_LifeTime_Ms))
	return err
}


func Query(query string, args ...interface{}) (*sqlx.Rows, error) {
	// log.V(5).Info("Query", query)
	log.Info("query %s", query)
	err := Db.Ping()
	if err != nil {
		return nil, err
	}
	return Db.Queryx(query, args...)
}

func Select(dst interface{}, query string, args ...interface{}) error {
	// log.Info("query %s", query)
	err := Db.Ping()
	if err != nil {
		return err
	}
	return Db.Select(dst, query, args...)
}

func SelectOne(dst interface{}, query string, args ...interface{}) error {
	err := Db.Ping()
	if err != nil {
		return err
	}

	if args != nil {
		for i := 0; i < len(args); i++ {
			arg_str := fmt.Sprintf("'%v'", args[i])
			query = strings.Replace(query, "?", arg_str, 1)
		}
	}

	// result := Db.Get(dst, query)
	// fmt.Println(result)
	// fmt.Println(dst)

	return Db.Get(dst, query)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	log.Info("query %s", query)
	err := Db.Ping()
	if err != nil {
		return nil, err
	}
	return Db.Exec(query, args...)
}