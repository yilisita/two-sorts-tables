/*
 * @Author: Wen Jiajun
 * @Date: 2022-07-01 22:47:28
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-01 23:09:17
 * @FilePath: \application\api\v1\request.go
 * @Description:
 */
package v1

import (
	"app/model"
	"app/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 6 functions

// GET
func ReadAllRequest(c *gin.Context) {
	reqs, err := model.ReadAllRequest()
	fmt.Println(reqs)
	c.JSON(http.StatusOK, utils.NewRes(err).WithData(reqs))
}

// POST
func SendRequest(c *gin.Context) {
	// A "Request" struct should be passed from the front end.
	var reqStr string
	fmt.Println(c.Request)

	_ = c.ShouldBindJSON(&reqStr)

	fmt.Println(reqStr)

	reqID, err := model.SendRequest(reqStr)
	log.Println("reqID:", reqID)
	c.JSON(http.StatusOK, utils.NewRes(err).WithData(reqID))
}

// POST
func HandleSingle(c *gin.Context) {
	id := c.PostForm("id")
	fmt.Println(id)
	rep, err := model.HandleSingle(id)
	fmt.Println(rep)
	fmt.Println(err)

	err = model.SendReport(rep)
	c.JSON(http.StatusOK, utils.NewRes(err))
}

// POST
func HandleAll(c *gin.Context) {
	rep, err := model.HandleAll()
	fmt.Println(rep)
	fmt.Println(err)

	err = model.SendReport(rep)
	c.JSON(http.StatusOK, utils.NewRes(err))
}

// POST
func RefuseRequest(c *gin.Context) {
	id := c.PostForm("id")
	err := model.RefuseRequest(id)
	c.JSON(http.StatusOK, utils.NewRes(err))
}
