/*
 * @Author: Wen Jiajun
 * @Date: 2022-07-02 01:43:42
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 12:24:31
 * @FilePath: \application\model\id.go
 * @Description:
	This deals with ID increasement in Hyperledger Fabric. Because it's
	depricatited to initialize an ID in chaincode, which may result unmatched
	proposal response. Therefore, it's always the server side's responsibility to
	pass every asset's ID to chaincode to maintain consistency.
*/
package model

import (
	e "app/error"
	"log"
	"strconv"
)

var ReqID int = 0
var ResID int = 0
var TableID int = 0

// When starting server, run these functions below to get the newest ID in the state DB.

func GetReqID() error {
	ReqID = 1
	res, err := Contract.EvaluateTransaction("GetReqID")
	if err != nil {
		log.Println(err)
		return e.TX_EVALUATION_ERROR
	}
	id, _ := strconv.Atoi(string(res))
	ReqID = id + 1
	return nil
}

func GetTableID() error {
	TableID = 1
	res, err := Contract.EvaluateTransaction("GetTableID")
	if err != nil {
		log.Println(err)
		return e.TX_EVALUATION_ERROR
	}
	id, _ := strconv.Atoi(string(res))
	TableID = id + 1
	return nil
}

func GetResID() error {
	ResID = 1
	res, err := Contract.EvaluateTransaction("GetResID")
	if err != nil {
		log.Println(err)
		return e.TX_EVALUATION_ERROR
	}
	id, _ := strconv.Atoi(string(res))
	ResID = id + 1
	return nil
}
