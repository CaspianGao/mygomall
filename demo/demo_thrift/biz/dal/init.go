package dal

import (
	"github.com/CaspianGao/mygomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/CaspianGao/mygomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
