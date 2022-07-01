/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-30 00:07:29
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-01 22:04:54
 * @FilePath: \application\model\request.go
 * @Description:
 */
package model

import (
	e "app/error"
	"encoding/json"
	"log"
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

func SendRequest(reqStr string) (string, error) {
	res, err := Contract.SubmitTransaction("SendRequest", reqStr)
	if err != nil {
		log.Println(err)
		return "", e.TX_SUBMITION_ERROR
	}

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
	var resReq []RequestView
	err = json.Unmarshal(res, &resReq)
	if err != nil {
		log.Println(err)
		return nil, e.JSON_PARSE_ERROR
	}
	return resReq, nil
}

func HandleAll() (string, error) {
	res, err := Contract.EvaluateTransaction("HandleAll")
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
	res, err := Contract.EvaluateTransaction("HandleSingle", rid)
	if err != nil {
		log.Println(err)
		return "", e.TX_EVALUATION_ERROR
	}

	return jsonSlice(string(res)), e.SUCCESS
}

func SendReport(repStr string) error {
	_, err := Contract.SubmitTransaction("SendReport", repStr)
	if err != nil {
		log.Println(err)
		return e.TX_SUBMITION_ERROR
	}
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
