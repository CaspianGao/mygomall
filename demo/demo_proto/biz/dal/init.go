package dal

import (
	"github.com/CaspianGao/mygomall/demo/demo_proto/biz/dal/mysql"
	"github.com/CaspianGao/mygomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
