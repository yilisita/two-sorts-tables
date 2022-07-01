/*
 * @Author: Wen Jiajun
 * @Date: 2022-03-26 18:48:42
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-01 23:26:32
 * @FilePath: \application\router\router.go
 * @Description:
 */

package router

import (
	v1 "app/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	app := gin.Default()

	// recovery middleware == error middleware in Express
	app.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.String(http.StatusInternalServerError, "an error occured in the server: %v", err)

		// c.Abort will stop pending handlers from being excuted
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	// app.Static("/static", "frontend/dist")
	// app.StaticFile("index", "frontend/dist/index.html")
	// app.LoadHTMLGlob("frontend/dist/index.html")

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello, world",
			"status":  200,
		})
		// c.HTML(http.StatusOK, "index.html", nil)
	})
	table := app.Group("v1/tables")
	{
		table.GET("/:id", v1.ReadMyTableByID) // /:id or /?id=
		table.GET("/", v1.GetAllTable)
		table.POST("/", v1.InsertATable)
	}

	pt := app.Group("v1/public_tables")
	{
		pt.GET("/", v1.ReadAllPublicTable)
		pt.GET("/:id", v1.ReadPublicTableByID)
	}

	req := app.Group("v1/requests")
	{
		req.GET("/", v1.ReadAllRequest)

		req.POST("/handle", v1.HandleSingle)
		req.POST("handle_all", v1.HandleAll)
		req.POST("/refuse", v1.RefuseRequest)
		req.POST("/send", v1.SendRequest)
	}

	return app
}
