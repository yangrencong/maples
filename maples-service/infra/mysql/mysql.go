package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	MysqlClient *gorm.DB
)

func Init() error {
	username := "root"       //账号
	password := "root123123" //密码
	host := "127.0.0.1"      //数据库地址，可以是Ip或者域名
	port := 3306             //数据库端口
	Dbname := "maples"       //数据库名
	timeout := "10s"         //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Mysql connect fatal!")
		return err
	}
	MysqlClient = db
	return nil
}

func Close() {
	sqlDB, _ := MysqlClient.DB()
	err := sqlDB.Close()
	if err != nil {
		log.Println("Mysql Close Fatal")
		return
	}
	log.Println("Mysql Close Success")
}
