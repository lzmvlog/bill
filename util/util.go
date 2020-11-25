package util

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine *xorm.Engine // 创建mysql连接

var Pool *redis.Pool //创建redis连接池

func Init() {
	var err error
	//Engine, err = xorm.NewEngine("mysql", "root:123456@tcp(192.168.1.105:3306)/bill?charset=utf8")
	Engine, err = xorm.NewEngine("mysql", "root:Root5683@@tcp(127.0.0.1:3306)/bill?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	// 控制台打印出生成的SQL语句
	Engine.ShowSQL(true)

	// 打印调试信息
	//engine.Logger().SetLevel(core.LOG_DEBUG)
	// 设置连接池的空闲数大小
	Engine.SetMaxIdleConns(10)
	// 设置最大打开连接数
	Engine.SetMaxOpenConns(100)
	//engine.SetMapper(core.SameMapper)
	// 关闭数据库连接
	//defer database.Close() // 注意这行代码要写在上面err判断的下面

	Pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按   需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}
