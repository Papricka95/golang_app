package testing

import (
	"banks/Config"
	"banks/Database"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/labstack/gommon/log"
)


func TestConnectDb(t *testing.T) {
	cfg_file := "../default.yml"
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
	if err != nil {
		t.Errorf("Connect to database is incorrect")
	}
}