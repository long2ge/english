package db_self

import (
	"github.com/gogf/gf/frame/g"
)

func Configbbbcvcc() {
	// go test -v -run TestConfigaaa config_test.go
	// INSERT INTO `user`(`name`) VALUES('john')
	_, err := g.DB().Table("test_table").Data(g.Map{"name": "john", "age": 18}).Insert()
	if err != nil {
		panic(err)
	}
}
