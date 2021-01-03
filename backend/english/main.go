package main

import (
	"english/db_self"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()

	// db := g.DB("default").Table("user").Data(g.Map{"name": "john"}).Insert()

	s.BindHandler("/:name", func(r *ghttp.Request) {

		// _, err := g.DB().Table("test_table").Data(g.Map{"name": "john", "age": 18}).Insert()
		// if err != nil {
		// 	panic(err)
		// }

		// r.Response.Writef("name: %v, pass: %v", r.Get("name"), r.Get("pass"))
		db_self.Configbbb()
		db_self.Configbbbcvcc()

		r.Response.WriteJson(g.Map{
			"name": r.Get("name"),
			"pass": r.Get("pass"),
		})

		// INSERT INTO `user`(`name`) VALUES('john')
		// r, err := g.DB().Table("user").Data(g.Map{"name": "john"}).Insert()

		// r, err := db.Insert("user", gdb.Map {
		// 	"name": "john",
		// })

		// r, err := g.DB().Table("user").Data(g.List{
		// 	{"name": "john_1"},
		// 	{"name": "john_2"},
		// 	{"name": "john_3"},
		// }).Insert()

	})

	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("哈喽世界！")
	})
	s.Run()
}
