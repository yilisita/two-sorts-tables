/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-30 12:52:28
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-01 23:30:52
 * @FilePath: \application\api\v1\table.go
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

// type Table struct {
// 	ID        string      `json:"id"`
// 	Area      string      `json:"area"`
// 	Year      string      `json:"year"`
// 	Month     string      `json:"month"`
// 	Columns   []string    `json:"columns"`
// 	Data      [][]float64 `json:"data"`
// 	Label     []string    `json:"label"`
// 	TableType string      `json:"table_type"`
// }

const (
	ecws string = "0"
	prod string = "1"
)

// POST
func InsertATable(c *gin.Context) {
	// Parse description
	tType := c.PostForm("table_type")
	fmt.Println("table_type:", tType)

	// Parse file
	file, fileHeader, err := c.Request.FormFile("file")
	fmt.Println(err)

	fileName := fileHeader.Filename
	log.Printf("File: %s uploaded\n", fileName)

	var tables []*model.Table
	switch tType {
	case prod:
		tables = utils.FormatProdFile(file)
	case ecws:
		tables = utils.FormatProdFile(file)
	}

	fmt.Println(tables)

	var ids []int
	for i := range tables {
		id, err := model.InsertATable(tables[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
				"status":  http.StatusInternalServerError,
			})
		}
		ids = append(ids, id)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "上传成功",
		"status":  200,
		"data":    ids,
	})
}

func GetAllTable(c *gin.Context) {
	tables, err := model.GetAllTable()
	fmt.Println(tables)
	c.JSON(http.StatusOK, utils.NewRes(err).WithData(tables)) // [Object Object]
}

func ReadMyTableByID(c *gin.Context) {
	id := c.Param("id")
	table, err := model.ReadMyTableByID(id)
	fmt.Println(table)
	c.JSON(http.StatusOK, utils.NewRes(err).WithData(table))
}

func ReadAllPublicTable(c *gin.Context) {
	tables, err := model.ReadAllPublicTable()
	fmt.Println(tables)
	c.JSON(http.StatusOK, utils.NewRes(err).WithData(tables))

}

func ReadPublicTableByID(c *gin.Context) {
	id := c.Param("id")
	table, err := model.ReadPublicTableByID(id)
	fmt.Println(table)
	c.JSON(http.StatusOK, utils.NewRes(err).WithData(table))
}
