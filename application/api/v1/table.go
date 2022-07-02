/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-30 12:52:28
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 20:37:42
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
/*
	curl --location --request POST 'http://localhost:4000/v1/tables' \
	--form 'table_type="1"' \
	-H 'content-type: multipart/form-data' \
	--form 'file=@"/home/wenjiajun/Desktop/go/src/github.com/hyperledger/fabric/scripts/fabric-samples/myChaincode/two-sorts-table/application/api/v1/prod.xlsx"'
*/
func InsertATable(c *gin.Context) {
	// Parse description
	tType := c.PostForm("table_type")
	fmt.Println("table_type:", tType)

	fmt.Println(c.Request)

	fileHeader, err := c.FormFile("file")
	fmt.Println(err)

	file, err := fileHeader.Open()
	// Parse file
	// file, fileHeader, err := c.Request.FormFile("file")
	fmt.Println(err)

	fileName := fileHeader.Filename
	log.Printf("File: %s uploaded\n", fileName)

	var tables []*model.Table
	switch tType {
	case prod:
		tables = utils.FormatProdFile(file)
	case ecws:
		tables = utils.FormatEcwsFile(file)
	}

	fmt.Println(tables)

	var ids []int
	for i := range tables {
		id, err := model.InsertATable(tables[i])
		fmt.Println(err)
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
	fmt.Println(id)
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
