/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-30 00:22:17
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-06-30 15:51:10
 * @FilePath: \application\model\report.go
 * @Description:
 */
package model

import "encoding/json"

type AttributeRes struct {
	ID          string      `json:"id"`
	ReqID       string      `json:"req_id"`
	TargetTable PublicTable `json:"target_table"`
	Service     string      `json:"service"`
	Result      float64     `json:"result"`
	Description string      `json:"description"`
	ResType     string      `json:"res_type"`
}

type AttributeResView struct {
	ID          string          `json:"id"`
	ReqID       string          `json:"req_id"`
	TargetTable PublicTableView `json:"target_table"`
	Service     string          `json:"service"`
	Result      float64         `json:"result"`
	Description string          `json:"description"`
	ResTypeStr  string          `json:"res_type"`
}

type TableRes struct {
	ID      string `json:"id"`
	ReqID   string `json:"req_id"`
	Res     Table  `json:"table"`
	ResType string `json:"res_type"`
}

type PublicTableView struct {
	ID           string   `json:"id"`
	Area         string   `json:"area"`
	Year         string   `json:"year"`
	Month        string   `json:"month"`
	Columns      []string `json:"columns"`
	NumOfObs     int      `json:"NumOfObs"`
	TableTypeStr string   `json:"table_type"`
	Label        []string `json:"label"`
}

// func ReadReport(id string)

func GetAllReports() ([]AttributeResView, error) {
	res, err := Contract.EvaluateTransaction("GetAllReports")
	if err != nil {
		return nil, err
	}

	var resRep []AttributeResView
	err = json.Unmarshal(res, &resRep)
	if err != nil {
		return nil, err
	}

	return resRep, nil
}

func ReadPurchasedTable() ([]*Table, error) {
	res, err := Contract.EvaluateTransaction("GetAllReports")
	if err != nil {
		return nil, err
	}

	var resRep []*Table
	err = json.Unmarshal(res, &resRep)
	if err != nil {
		return nil, err
	}

	return resRep, nil
}
