package main

import (
	// "github.com/labstack/echo"

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

	json_map := make(map[string]interface{})
	err_decode := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err_decode != nil {
		return c.String(http.StatusForbidden, "Invalid body of request.")
	}

	switch role[0] {
	case "client":
		// пополняем значение sum из базы
		// handle request for correct role from PostgreSQL
		// Database.Select(role, ...)
		//
	case "admin":
		// вычитаем значение sum из базы
	default:
		// pass
	}
	return c.String(http.StatusOK, "Test echo")
}

func main() {
	// лог файл
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

	e.Start(":8000")
}
