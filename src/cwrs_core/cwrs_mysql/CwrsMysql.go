package cwrs_mysql

import (
	"cwrs_go_server/src/cwrs_core/cwrs_viper"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	InitMySQL()
}

func InitMySQL() {
	var err error
	DB, err = sql.Open(
		cwrs_viper.GlobalViper.GetString("mysql.driver-name"),
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			cwrs_viper.GlobalViper.GetString("mysql.user-name"),
			cwrs_viper.GlobalViper.GetString("mysql.password"),
			cwrs_viper.GlobalViper.GetString("mysql.path"),
			cwrs_viper.GlobalViper.GetString("mysql.port"),
			cwrs_viper.GlobalViper.GetString("mysql.db-name"),
			cwrs_viper.GlobalViper.GetString("mysql.config")))
	if err != nil {
		panic(err.Error())
	}
	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("MySQL Initialize OK !")
}
