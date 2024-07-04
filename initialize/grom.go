package initialize

import (
	"GoLangStudy/config"
	"GoLangStudy/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始grom框架
func InitGorm() {
	// 拼接链接
	var mysqlConfig config.ServerConfig
	fmt.Println("mysql配置信息：", mysqlConfig.Mysql.User)
	// 拼接mysql连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)%s?charset=%s&parseTime=%s&loc=%s",
		mysqlConfig.Mysql.User, mysqlConfig.Mysql.Password, mysqlConfig.Mysql.Host, mysqlConfig.Mysql.Port, mysqlConfig.Mysql.Database, mysqlConfig.Mysql.Charset,
		mysqlConfig.Mysql.ParseTime, mysqlConfig.Mysql.TimeZone)
	fmt.Println("mysql连接信息：", dsn)

	// 设置mysql连接
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: mysqlConfig.Mysql.Gorm.SkipDefaultTx, // 跳过事务
		PrepareStmt:            mysqlConfig.Mysql.Gorm.PreparedStmt,
	}

	// 打开连接 得到连接
	clinet, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         mysqlConfig.Mysql.DefaultStringSize,
		DisableDatetimePrecision:  mysqlConfig.Mysql.DisableDatetimePrecision,
		SkipInitializeWithVersion: mysqlConfig.Mysql.SkipInitializeWithVersion,
	}), gormConfig)

	if err != nil {
		panic(fmt.Sprintf("创建mysql客户端失败：%s", err))
	}
	global.GvaMysqlClient = clinet
}
