/*
 * @Author: Wen Jiajun
 * @Date: 2022-03-25 15:43:26
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-01 23:26:59
 * @FilePath: \application\main.go
 * @Description:
 */

package main

import (
	// "log"
	//"seller-app/model"
	// "seller-app/router"
	"app/model"
	"app/router"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	model.InitFabric()
	app := router.InitRouter()

	app.Run(":4000")
}

// type Table struct {
// 	ID        string      `json:"id"`
// 	Area      string      `json:"area"`
// 	Year      string      `json:"year"`
// 	Month     string      `json:"month"`

// 	Columns   []string    `json:"columns"`
// 	Label     []string    `json:"label"`
// 	TableType string      `json:"table_type"`

// 	Data      [][]float64 `json:"data"`
// }
func formatProdFile(r io.Reader) []*model.Table {
	var res []*model.Table

	file, err := excelize.OpenReader(r)
	if err != nil {
		log.Println(err)
		return nil
	}

	sheets := file.GetSheetList()
	for _, s := range sheets {
		cols, err := file.GetCols(s)
		if err != nil {
			log.Println(err)
			return nil
		}

		// Year, Month at column 7, row 11 (6, 10)
		var ym = cols[6][10]
		year_month := strings.Split(ym, "年")
		year := year_month[0]
		month := year_month[1]
		month = strings.TrimRight(month, "月")
		// fmt.Println(year, month)

		// fmt.Println(cols)

		// Area at column 1, row 2
		area := cols[0][1]
		if strings.HasPrefix(area, "国网") {
			continue
		}
		area = strings.TrimPrefix(area, "黑龙江省")
		area = strings.TrimSuffix(area, "供电公司")

		// fmt.Println(area)
		// for _, col := range cols {
		// 	for i, v := range col {
		// 		if i > 20 {
		// 			continue
		// 		}
		// 		fmt.Println(i, v)
		// 	}
		// }

		// Columns, row 14
		columns := model.ProdCols
		// fmt.Println(len(columns))

		// Label
		labels := cols[0][16:len(cols[0])]
		// fmt.Println(labels)
		// fmt.Println("Length of label:", len(labels))

		// Data column ranges 1-7, 14-24, row index begins from 16
		var data [][]float64
		var index = []int{1, 2, 3, 4, 5, 6, 7, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
		for _, v := range index {
			d := cols[v][16:]
			// fmt.Println(d)
			// fmt.Println(len(d))
			data = append(data, strSlice2Float64Slice(d))
		}

		var t = model.Table{
			Year:      year,
			Month:     month,
			Area:      area,
			Data:      data,
			TableType: "1", // prod
			Label:     labels,
			Columns:   columns,
		}

		res = append(res, &t)
	}
	return res
}

func strSlice2Float64Slice(s []string) []float64 {
	var res []float64
	for _, v := range s {
		vf, _ := strconv.ParseFloat(v, 64)
		res = append(res, vf)
	}
	return res
}

func formatEcwsFile(r io.Reader) []*model.Table {
	var res []*model.Table

	file, err := excelize.OpenReader(r)
	if err != nil {
		log.Println(err)
		return nil
	}

	sheets := file.GetSheetList()
	for _, s := range sheets {
		cols, err := file.GetCols(s)
		if err != nil {
			log.Println(err)
			return nil
		}

		// Area at column 1, row 2
		area := cols[0][1]
		if strings.HasPrefix(area, "国网") {
			continue
		}
		area = strings.TrimPrefix(area, "黑龙江省")
		area = strings.TrimSuffix(area, "供电公司")
		// fmt.Println(area)

		// fmt.Println(cols)

		// for i, v := range cols[2] {
		// 	fmt.Println(i, v)
		// }

		// Year, Month at column 3, row 11 (2, 10)
		var ym = cols[2][10]
		// fmt.Println(ym)
		year_month := strings.Split(ym, "年")
		year := year_month[0]
		month := year_month[1]
		month = strings.TrimRight(month, "月")
		// fmt.Println(year, month)

		// Columns
		columns := model.EcwsCols

		// Labels
		labels := model.EcwsLabels

		// fmt.Println(len(columns), len(labels))

		// Data row index begins from 16, columns index ranges 2-7
		var index = []int{2, 3, 4, 5, 6, 7}
		var data [][]float64
		for _, v := range index {
			d := cols[v][16:len(cols[v])]
			// fmt.Println(strSlice2Float64Slice(d))
			data = append(data, strSlice2Float64Slice(d))
		}
		// fmt.Println(data)

		var t = model.Table{
			Year:      year,
			Month:     month,
			Area:      area,
			Data:      data,
			TableType: "0", // prod
			Label:     labels,
			Columns:   columns,
		}

		res = append(res, &t)

		// return nil
	}

	return res
}
