package main

import (
	"banks/Config"
	"banks/Database"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func handle_get_sum(c echo.Context) error {
	header := c.Request().Header
	role, err := header["User-Role"]
	if !err {
		// fmt.Println(role) //
		return c.String(http.StatusForbidden, "access denied.")
	}

	data_request := make(map[string]int)
	err_decode := json.NewDecoder(c.Request().Body).Decode(&data_request)
	if err_decode != nil {
		return c.String(http.StatusForbidden, "Invalid body of request.")
	}

	GET_QUERY := fmt.Sprintf(`SELECT id, sum FROM %s`, Config.Config.Tables.Transact)
	type SumData struct {
		Id int
		Sum int
	}
	


	sum_request, err_sum := data_request["sum"]
	if !err_sum {
		return c.String(http.StatusForbidden, "Request hasn't the sum.")
	}
	info := []SumData{}
	err_select := Database.Select(&info, GET_QUERY)

	if err_select != nil {
		log.Errorf("Error select: %s", GET_QUERY)
		return err_select
	}
	
	sum_base := 0
	switch role[0] {
	case "client":
		UPDATE_QUERY := ""
		if len(info) == 0 {
			UPDATE_QUERY = fmt.Sprintf("INSERT INTO %s (sum) VALUES (%d)", Config.Config.Tables.Transact, sum_request)
		}else{
			sum_base = info[0].Sum // текущая сумма из базы
			UPDATE_QUERY = fmt.Sprintf("UPDATE %s SET sum=%d", Config.Config.Tables.Transact, sum_base + sum_request)
		}
		Database.Exec(UPDATE_QUERY)
		return c.String(http.StatusOK, "Sum was been changed.")
	case "admin":
		if len(info) == 0 {
			return c.String(http.StatusForbidden, "Amount not paid")
		}
		sum_base = info[0].Sum // текущая сумма из базы
		
		if sum_request > sum_base {
			return c.String(http.StatusForbidden, "Insufficient funds")
		}
		UPDATE_QUERY := fmt.Sprintf("UPDATE %s SET sum=%d", Config.Config.Tables.Transact, sum_base - sum_request)
		Database.Query(UPDATE_QUERY)
		return c.String(http.StatusOK, "Sum was been changed.")
	default:
		return c.String(http.StatusForbidden, "access denied.")
	}
}

func main() {
	// лог файл
	fmt.Println("Start app")
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)

	// connect with database
	cfg_file := "./default.yml"
	cfg_path, err := filepath.Abs(cfg_file)
	if err != nil {
		log.Error(fmt.Sprintf("Couldn't get absolute path to config file: %s", err))
		return
	}
	err = Config.LoadConfig(cfg_path)
	if err != nil {
		log.Error(fmt.Sprintf("Couldn't load config: %s", err))
		return
	}

	err = Database.InitDatabase(Config.Config)
	if err != nil {
		log.Error(err, "Couldn't init database", "config", Config.Config)
	}

	// старт Echo servers and handle requests
	e := echo.New()
	e.GET("/bank", handle_get_sum)

	e.Start("0.0.0.0:3000")
}
