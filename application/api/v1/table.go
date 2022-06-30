/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-30 12:52:28
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-06-30 17:11:07
 * @FilePath: \application\api\v1\table.go
 * @Description:
 */
package v1

import (
	"app/model"
	"app/utils"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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

func InsertATable(c *gin.Context) {
	// Parse description
	tType := c.PostForm("table_type")
	fmt.Println("Description:", tType)

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
		"message": "长传成功",
		"status":  200,
		"data":    ids,
	})
}

func formatFile(r io.Reader) {
	file, err := excelize.OpenReader(r)
	if err != nil {
		log.Println(err)
		return
	}

	sheets := file.GetSheetList()
	for _, s := range sheets {
		cols, err := file.GetCols(s)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(cols)
		return
	}
}
