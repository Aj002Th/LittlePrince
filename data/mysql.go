package data

import (
	"fmt"

	"github.com/Aj002Th/LittlePrince/pkg/logging"
	"github.com/Aj002Th/LittlePrince/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQL *gorm.DB

func init() {
	var err error
	// dbType := setting.Database.Type
	user := setting.Database.User
	pwd := setting.Database.Password
	host := setting.Database.Host
	dbName := setting.Database.Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, dbName)

	MySQL, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		logging.Fatal(err)
	}
}

func GetMySQLConn() *gorm.DB {
	return MySQL
}
