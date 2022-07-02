/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-29 20:25:11
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 14:28:33
 * @FilePath: \application\model\table.go
 * @Description:
 */
package model

import (
	e "app/error"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
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
	t.ID = fmt.Sprint(TableID)
	tJSON, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return -1, e.JSON_PARSE_ERROR
	}
	res, err := Contract.SubmitTransaction("InsertATable", string(tJSON))
	if err != nil {
		log.Println(err)
		return -1, e.TX_SUBMITION_ERROR
	}

	id, _ := strconv.Atoi(string(res))
	TableID++
	return id, e.SUCCESS
}

func ReadPublicTableByID(id string) (*PublicTable, error) {
	res, err := Contract.EvaluateTransaction("ReadPublicTableByID", id)
	if err != nil {
		log.Println(err)
		return nil, e.TX_EVALUATION_ERROR
	}

	if res == nil {
		return nil, e.NO_TABLE
	}

	var resTable PublicTable
	err = json.Unmarshal(res, &resTable)
	if err != nil {
		log.Println(err)
		return nil, e.JSON_PARSE_ERROR
	}

	return &resTable, e.SUCCESS
}

func ReadAllPublicTable() ([]*PublicTable, error) {
	res, err := Contract.EvaluateTransaction("ReadAllPublicTable")
	if err != nil {
		log.Println(err)
		return nil, e.TX_EVALUATION_ERROR
	}

	if res == nil {
		return nil, e.NO_TABLE
	}

	var resTable []*PublicTable
	err = json.Unmarshal(res, &resTable)
	if err != nil {
		log.Println(err)
		return nil, e.JSON_PARSE_ERROR
	}

	return resTable, e.SUCCESS
}

func ReadMyTableByID(tableID string) (*Table, error) {
	txn, err := Contract.CreateTransaction(
		"ReadMyTableByID",
		gateway.WithEndorsingPeers("peer0.org1.example.com:7051"),
	)
	if err != nil {
		fmt.Printf("Failed to create transaction: %s\n", err)
		return nil, e.TX_CREATION_ERROR
	}

	res, err := txn.Evaluate(tableID)
	if err != nil {
		log.Println(err)
		return nil, e.TX_EVALUATION_ERROR
	}

	if res == nil {
		return nil, e.NO_TABLE
	}

	var resTable Table
	err = json.Unmarshal(res, &resTable)
	if err != nil {
		log.Println(err)
		return nil, e.JSON_PARSE_ERROR
	}
	return &resTable, e.SUCCESS
}

func GetAllTable() ([]*Table, error) {
	txn, err := Contract.CreateTransaction(
		"GetAllTable",
		gateway.WithEndorsingPeers("peer0.org1.example.com:7051"),
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
		return nil, e.NO_TABLE
	}

	var resTable []*Table
	err = json.Unmarshal(res, &resTable)
	if err != nil {
		log.Println(err)
		return nil, e.JSON_PARSE_ERROR
	}

	return resTable, e.SUCCESS
}
