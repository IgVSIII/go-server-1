package readconfig

import (
	"encoding/json"
	"loglib"
	"os"
)

type Config struct {
	Address string `json: "address"`
	Key     string `json: "key"`
	DBname  string `json: "dbname"`
}

func GetConfig(pathconf string) Config {

	file, err := os.Open(pathconf)

	loglib.CheckFatall(err, "Config file error open")

	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Config{}
	errd := decoder.Decode(&configuration)

	loglib.CheckFatall(errd, "Config file parse json erro")

	return configuration

}
