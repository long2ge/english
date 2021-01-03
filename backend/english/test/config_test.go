package test

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/frame/g"
)

func TestAaa(t *testing.T) {
	fmt.Println(123)
	// go test -v -run TestAaa config_test.go
}

func TestHelloWorld(t *testing.T) {
	t.Log("hello world")
	// go test -v -run TestHelloWorld config_test.go
}

// 基本配置使用
func TestConfig(t *testing.T) {
	// go test -v -run TestConfig config_test.go

	// 默认当前路径或者config路径，默认文件config.toml
	// /home/www/templates/
	fmt.Println(g.Config().Get("title"))
	fmt.Println(g.Cfg().Get("database"))
	// 127.0.0.1:6379,1
	// c := g.Cfg()
	// // 分组方式
	// fmt.Println(c.Get("redis.cache"))
	// // 数组方式：test2
	// fmt.Println(c.Get("database.default.1.name"))
}

func TestConfigaaa(t *testing.T) {
	// go test -v -run TestConfigaaa config_test.go
	// INSERT INTO `user`(`name`) VALUES('john')
	_, err := g.DB().Table("test_table").Data(g.Map{"name": "john", "age": 18}).Insert()
	if err != nil {
		panic(err)
	}
}

func TestConfigbbb() {
	// go test -v -run TestConfigaaa config_test.go
	// INSERT INTO `user`(`name`) VALUES('john')
	_, err := g.DB().Table("test_table").Data(g.Map{"name": "john", "age": 18}).Insert()
	if err != nil {
		panic(err)
	}
}
