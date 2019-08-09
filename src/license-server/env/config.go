package env

import (
	"encoding/json"
	"license-server/utils/logger"
	"os"
	"path/filepath"
	"runtime"
)

// Config contians variables for DB and Server
var Config = struct {
	Database struct {
		Host     string `json:"db_host"`
		Port     string `json:"db_port"`
		User     string `json:"db_user"`
		Password string `json:"db_password"`
		Name     string `json:"db_name"`
		Secure   bool   `json:"db_secure"`
	}
	Server struct {
		URL string `json:"server_url"`
	}
}{}

func getConfigDir() string {
	_, cwd, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	parentDir := filepath.Dir(cwd)
	return filepath.Join(parentDir, "config.json")
}

func initializeConfig() {

	configFile, err := os.Open(getConfigDir())
	if err != nil {
		logger.Error.Println("Cannot locate config.json; no configuration is loaded")
	} else {
		defer configFile.Close()
		for _, ptr := range []interface{}{&Config.Database, &Config.Server} {
			decoder := json.NewDecoder(configFile)
			if err := decoder.Decode(ptr); err != nil {
				panic(err)
			} else if _, err = configFile.Seek(0, 0); err != nil {
				panic(err)
			}
		}
	}
	return
}

func init() {
	initializeConfig()
}
