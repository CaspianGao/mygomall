/* 
该文件用于测试数据库的操作 
*/

package main

import (
	"github.com/CaspianGao/mygomall/demo/demo_proto/biz/dal"
	"github.com/CaspianGao/mygomall/demo/demo_proto/biz/dal/mysql"
	"github.com/CaspianGao/mygomall/demo/demo_proto/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()

	// CURD:

	// Create
	mysql.DB.Create(&model.User{Email: "123@example.com", Password: "examplepassword"})

	// Update
	// mysql.DB.Model(&model.User{}).Where("email = ?", "123@example.com").Update("password", "1234567")

	// Read
	//var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "123@example.com").First(&row)
	// fmt.Println(row)

	// Delete
	// mysql.DB.Unscoped().Where(" = ?", "123@example.com").Delete(&model.User{})

}
