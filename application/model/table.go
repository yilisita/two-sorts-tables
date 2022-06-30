/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-29 20:25:11
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-06-30 17:08:10
 * @FilePath: \application\model\table.go
 * @Description:
 */
package model

import (
	e "app/error"
	"encoding/json"
	"strconv"
)

type PublicTable struct {
	ID        string   `json:"id"`
	Area      string   `json:"area"`
	Year      string   `json:"year"`
	Month     string   `json:"month"`
	Columns   []string `json:"columns"`
	NumOfObs  int      `json:"num_of_obs"`
	TableType string   `json:"table_type"`
	Label     []string `json:"label"`
}

type Table struct {
	ID        string      `json:"id"`
	Area      string      `json:"area"`
	Year      string      `json:"year"`
	Month     string      `json:"month"`
	Columns   []string    `json:"columns"`
	Data      [][]float64 `json:"data"`
	Label     []string    `json:"label"`
	TableType string      `json:"table_type"`
}

func InsertATable(t *Table) (int, error) {
	tJSON, err := json.Marshal(t)
	if err != nil {
		return -1, e.JSON_PARSE_ERROR
	}
	res, err := Contract.SubmitTransaction("InsertATable", string(tJSON))
	if err != nil {
		return -1, err
	}

	id, _ := strconv.Atoi(string(res))
	return id, nil
}

func ReadPublicTableByID(id string) (*PublicTable, error) {
	res, err := Contract.EvaluateTransaction("ReadPublicTableByID", id)
	if err != nil {
		return nil, err
	}

	var resTable PublicTable
	err = json.Unmarshal(res, &resTable)
	if err != nil {
		return nil, e.JSON_PARSE_ERROR
	}

	return &resTable, nil
}

func ReadAllPublicTable() ([]*PublicTable, error) {
	res, err := Contract.EvaluateTransaction("ReadAllPublicTable")
	if err != nil {
		return nil, err
	}

	var resTable []*PublicTable
	err = json.Unmarshal(res, &resTable)
	if err != nil {
		return nil, e.JSON_PARSE_ERROR
	}

	return resTable, nil
}

func ReadMyTableByID(tableID string) (*Table, error) {
	res, err := Contract.EvaluateTransaction("ReadMyTableByID", tableID)
	if err != nil {
		return nil, err
	}

	var resTable Table
	err = json.Unmarshal(res, &resTable)
	if err != nil {
		return nil, err
	}
	return &resTable, nil
}
