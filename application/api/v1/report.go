/*
 * @Author: Wen Jiajun
 * @Date: 2022-07-01 23:09:52
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-01 23:13:27
 * @FilePath: \application\api\v1\report.go
 * @Description:
 */
package v1

import (
	"app/model"
	"app/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func GetAllReports(c *gin.Context) {
	reps, err := model.GetAllReports()
	fmt.Println(reps)
	fmt.Println(err)
	c.JSON(http.StatusOK, utils.NewRes(err).WithData(reps))
}

// GET
func ReadPurchasedTable(c *gin.Context) {
	reps, err := model.ReadPurchasedTable()
	fmt.Println(reps)
	fmt.Println(err)
	c.JSON(http.StatusOK, utils.NewRes(err).WithData(reps))
}
