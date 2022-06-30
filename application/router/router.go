/*
 * @Author: Wen Jiajun
 * @Date: 2022-03-26 18:48:42
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-06-29 20:07:59
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

	app.Static("/static", "frontend/dist")
	app.StaticFile("index", "frontend/dist/index.html")
	app.LoadHTMLGlob("frontend/dist/index.html")

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	table := app.Group("v1/table")
	{
		table.GET("/:id", v1.GetTable) // /:id or /?id=
		table.GET("/", v1.GetAllTable)
		table.GET("/search/:owner", v1.GetTableByOwner)
		table.GET("/public:id", v1.PublicTable)
		table.POST("/create", v1.CreateTable)
	}

	req := app.Group("v1/request")
	{
		req.GET("/", v1.GetAllRequest)
		req.GET("/:id", v1.GetRequest)

		req.DELETE("/:id", v1.DeleteRequest)
		req.POST("/handle", v1.HandleRequest)
		req.GET("/refuse", v1.RefuseRequest)
		req.POST("/sendRequest", v1.SendRequest)
	}

	return app
}
