//@see <https://www.dougcodes.com/go-lang/building-a-web-application-with-martini-and-gorm-part-1>
package main

import (
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/binding"
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

	m.Get("/item/add", func(r render.Render) {
		var retData struct {
			Item Item
		}

		r.HTML(200, "item_edit", retData)
	})

	m.Post("/item/save", binding.Bind(Item{}), func(r render.Render, i Item) {
		db.Save(&i)
		r.Redirect("/")
	})

    m.Get("/item/edit/:id", func(r render.Render, p martini.Params) {
        var retData struct {
            Item Item
        }

        db.Where("id = ?", p["id"]).Find(&retData.Item)

        r.HTML(200, "item_edit", retData)
    })

    m.Get("/item/remove/:id", func(r render.Render, p martini.Params) {
        //为了让如上的Where语句知道是什么表而造的空结构
        var item Item
        db.Where("id = ?", p["id"]).Delete(&item)

        r.Redirect("/")
    })

	m.Run()
}
