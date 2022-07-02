/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-30 00:22:17
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 14:28:09
 * @FilePath: \application\model\report.go
 * @Description:
 */
package model

import (
	e "app/error"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

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
	txn, err := Contract.CreateTransaction(
		"GetAllReports",
		gateway.WithEndorsingPeers("peer0.org2.example.com:9051"),
	)
	if err != nil {
		fmt.Printf("Failed to create transaction: %s\n", err)
		return nil, e.TX_CREATION_ERROR
	}

	res, err := txn.Evaluate()
	if err != nil {
		log.Println(err)
		return nil, e.TX_EVALUATION_ERROR
	}

	if res == nil {
		return nil, e.NO_RES
	}

	var resRep []AttributeResView
	err = json.Unmarshal(res, &resRep)
	if err != nil {
		return nil, e.JSON_PARSE_ERROR
	}

	return resRep, e.SUCCESS
}

func ReadPurchasedTable() ([]*Table, error) {
	txn, err := Contract.CreateTransaction(
		"ReadPurchasedTable",
		gateway.WithEndorsingPeers("peer0.org2.example.com:9051"),
	)
	if err != nil {
		fmt.Printf("Failed to create transaction: %s\n", err)
		return nil, e.TX_CREATION_ERROR
	}

	res, err := txn.Evaluate()
	if err != nil {
		fmt.Printf("Failed to read purchased table: %v\n", err)
		return nil, err
	}

	if res == nil {
		return nil, e.NO_RES
	}

	var resRep []*Table
	err = json.Unmarshal(res, &resRep)
	if err != nil {
		return nil, e.JSON_PARSE_ERROR
	}

	return resRep, e.SUCCESS
}
