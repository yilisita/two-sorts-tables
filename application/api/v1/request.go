/*
 * @Author: Wen Jiajun
 * @Date: 2022-07-01 22:47:28
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 14:51:20
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

// type BasicRequest struct {
// 	ID            string `json:"id"`
// 	ReqType       string `json:"req_type"`  1
// 	Demander      string `json:"demander"`
// 	TargetTableID string `json:"target_table_id"` 1
// 	Service       int    `json:"service"` 0
// 	RequestTime   string `json:"request_time"` 2022-07-02 14:36
// 	State         uint   `json:"state"` 0
//  attributeID    attribute_id 1
// }
// POST
func SendRequest(c *gin.Context) {
	// A "Request" struct should be passed from the front end.
	raw, err := c.GetRawData()
	fmt.Println(err)
	var reqStr string = string(raw)

	fmt.Println(c.Request)

	// err := c.ShouldBindJSON(&reqStr)
	// fmt.Println(err)

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
