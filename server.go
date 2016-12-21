//@see <https://www.dougcodes.com/go-lang/building-a-web-application-with-martini-and-gorm-part-1>
package main

import (
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
)

var (
	db            *gorm.DB
	sqlConnection string
)

func main() {

	var err error
	//username:password@tcp(host)/databasename?additional connection params
	sqlConnection = "root:zhs@tcp(127.0.0.1:3306)/sen?parseTime=True"

	db, err = gorm.Open("mysql", sqlConnection)

	if err != nil {
		panic(err)
		return
	}

	m := martini.Classic()

	m.Use(render.Renderer())
	m.Get("/", func(r render.Render) {
		var retData struct {
			Items []Item
		}

		//db.Find is a basic gorm command, 
		//it’s simply going to run SELECT * FROM items
		db.Find(&retData.Items)
		r.HTML(200, "index", retData)

	})
	m.Run()
}
