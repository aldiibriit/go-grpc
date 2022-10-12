package database

import (
	"encoding/json"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	DBType   string `json:"DBType"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Protocol string `json:"Protocol"`
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	Database string `json:"Database"`
}

func Connect(database string) (string, string) {
	var config DBConfig
	var err error

	switch database {
	case "peduliATM":
		err = json.Unmarshal(PeduliATM, &config)
	}

	if err != nil {
		panic(err)
	}

	var dbname string
	switch config.DBType {
	case "mysql":
		dbname = config.Username + ":" + config.Password + "@" + config.Protocol + "(" + config.Host + ":" + config.Port + ")/" + config.Database
	case "mssql":
		dbname = "server=" + config.Host + ";user id=" + config.Username + ";database=" + config.Database + ";password=" + config.Password
	}

	return config.DBType, dbname
}
