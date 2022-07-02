/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-30 00:07:29
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 15:20:55
 * @FilePath: \application\model\request.go
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

const REQ = "REQ-"

type BasicRequest struct {
	ID            string `json:"id"`
	ReqType       string `json:"req_type"`
	Demander      string `json:"demander"`
	TargetTableID string `json:"target_table_id"`
	Service       int    `json:"service"`
	RequestTime   string `json:"request_time"`
	State         uint   `json:"state"`
}

type ProdRequest struct {
	*BasicRequest
	AttributeID int `json:"attribute_id"`
}

type EcwsRequest struct {
	*BasicRequest
	AttributeID int `json:"attribute_id"`
	IndexCode   int `json:"index_code"`
}

type TableRequest struct {
	*BasicRequest
}

// TODO: ID
func SendRequest(reqStr string) (string, error) {
	var i interface{}
	fmt.Println("ReqStr:", reqStr)
	err := json.Unmarshal([]byte(reqStr), &i)
	iMapper := i.(map[string]interface{})
	iMapper["id"] = REQ + fmt.Sprint(ReqID)
	reqJSON, err := json.Marshal(iMapper)
	reqStr = string(reqJSON)

	res, err := Contract.SubmitTransaction("SendRequest", reqStr)
	if err != nil {
		log.Println(err)
		return "", e.TX_SUBMITION_ERROR
	}

	ReqID++
	return string(res), e.SUCCESS
}

func newReqID(id string) string {
	return REQ + id
}

type RequestView struct {
	ID            string `json:"id"`
	ReqTypeStr    string `json:"req_type_str"`
	Demander      string `json:"demander"`
	TargetTableID string `json:"target_table_id"`
	Service       string `json:"service"`
	RequestTime   string `json:"request_time"`
	State         string `json:"state"`
	Attribute     string `json:"attribute"`
	Index         string `json:"index"`
}

func ReadAllRequest() ([]RequestView, error) {
	res, err := Contract.EvaluateTransaction("ReadAllRequest")
	if err != nil {
		log.Println(err)
		return nil, e.TX_EVALUATION_ERROR
	}

	if res == nil {
		return nil, e.NO_REQ
	}

	var resReq []RequestView
	err = json.Unmarshal(res, &resReq)
	if err != nil {
		log.Println(err)
		return nil, e.JSON_PARSE_ERROR
	}
	return resReq, e.SUCCESS
}

func HandleAll() (string, error) {
	txn, err := Contract.CreateTransaction(
		"HandleAll",
		gateway.WithEndorsingPeers("peer0.org1.example.com:7051"),
	)
	if err != nil {
		fmt.Printf("Failed to create transaction: %s\n", err)
		return "", e.TX_CREATION_ERROR
	}
	res, err := txn.Evaluate()
	if err != nil {
		log.Println(err)
		return "", e.TX_EVALUATION_ERROR
	}

	return string(res), e.SUCCESS
}

func jsonSlice(s string) string {
	return "[" + s + "]"
}

func HandleSingle(id string) (string, error) {
	rid := newReqID(id)
	txn, err := Contract.CreateTransaction(
		"HandleSingle",
		gateway.WithEndorsingPeers("peer0.org1.example.com:7051"),
	)
	if err != nil {
		fmt.Printf("Failed to create transaction: %s\n", err)
		return "", e.TX_CREATION_ERROR
	}
	res, err := txn.Evaluate(rid)
	if err != nil {
		log.Println(err)
		return "", e.TX_EVALUATION_ERROR
	}

	return string(res), e.SUCCESS
}

// []Report
func SendReport(repStr string) error {

	fmt.Println("RepStr:", repStr)
	var i interface{}
	err := json.Unmarshal([]byte(repStr), &i)
	fmt.Println(err)
	iMapper := i.([]interface{})
	thisID := ResID
	for _, im := range iMapper {
		m := im.(map[string]interface{})
		m["id"] = fmt.Sprint(thisID)
		thisID++
	}
	repJSON, err := json.Marshal(iMapper)
	repStr = string(repJSON)

	_, err = Contract.SubmitTransaction("SendReport", repStr)
	if err != nil {
		log.Println(err)
		return e.TX_SUBMITION_ERROR
	}
	ResID = thisID
	return e.SUCCESS
}

func RefuseRequest(id string) error {
	rid := newReqID(id)
	_, err := Contract.SubmitTransaction("RefuseRequest", rid)
	if err != nil {
		log.Println(err)
		return e.TX_SUBMITION_ERROR
	}

	return e.SUCCESS
}
