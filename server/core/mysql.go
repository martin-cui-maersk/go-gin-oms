package core

import (
	"fmt"
	"github.com/martin-cui-maersk/go-gin-oms/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var DB *gorm.DB

// DB 连接数据库
func DB() *gorm.DB {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	global.AppLogger.Fatal("Error loading .env file. " + err.Error())
	//	log.Fatalf("Error loading .env file. %v\n", err)
	//	return nil
	//}
	//
	//DbHost := os.Getenv("DB_HOST")
	//DbPort := os.Getenv("DB_PORT")
	//DbUser := os.Getenv("DB_USER")
	//DbPass := os.Getenv("DB_PASS")
	//DbName := os.Getenv("DB_NAME")

	DbHost := global.Server.MySQL.Host
	DbPort := global.Server.MySQL.Port
	DbUser := global.Server.MySQL.User
	DbPass := global.Server.MySQL.Password
	DbName := global.Server.MySQL.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPass, DbHost, DbPort, DbName)
	// DB, err := gorm.Open... := 对全局变量赋值会是nil，外部调用 ConnectDB 的时候，使用:=，DB是局部变量
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		global.AppLogger.Fatal(fmt.Sprintf("error connecting to database. %s", err.Error()))
		panic(fmt.Errorf("error connecting to database, %w", err))
	}
	// DB.AutoMigrate(&User{})
	return DB
}
