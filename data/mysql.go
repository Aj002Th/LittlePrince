package data

import (
	"fmt"

	"github.com/Aj002Th/LittlePrince/pkg/logging"
	"github.com/Aj002Th/LittlePrince/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var MySQL *sqlx.DB

func init() {
	var err error
	dbType := setting.DatabaseSetting.Type
	user := setting.DatabaseSetting.User
	pwd := setting.DatabaseSetting.Password
	host := setting.DatabaseSetting.Host
	port := setting.DatabaseSetting.Port
	dbName := setting.DatabaseSetting.Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pwd, host, port, dbName)

	MySQL, err = sqlx.Connect(dbType, dsn)
	if err != nil {
		logging.Fatal(err)
	}
}

func GetMySQLConn() *sqlx.DB {
	return MySQL
}
