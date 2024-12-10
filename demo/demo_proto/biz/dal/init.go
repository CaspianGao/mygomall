package dal

import (
	"github.com/CaspianGao/mygomall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
