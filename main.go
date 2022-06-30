/*
 * @Author: Wen Jiajun
 * @Date: 2021-12-05 10:46:42
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-06-29 15:42:15
 * @FilePath: \electricity-data-trade\main.go
 * @Description:
 */
/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-samples/asset-transfer-basic/chaincode-go/chaincode"
)

func main() {
	if 1 == 0 {
		assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
		if err != nil {
			log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
		}

		if err := assetChaincode.Start(); err != nil {
			log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
		}
	}

	if 1 == 1 {
		var r = chaincode.TableRequest{
			BasicRequest: &chaincode.BasicRequest{
				ID:            "REQ-1",
				ReqType:       "2",
				Demander:      "Org2MSP",
				TargetTableID: "1",
				Service:       1,
				RequestTime:   "2022-06-27 21:48",
				State:         0},
		}
		rJSON, _ := json.Marshal(r)
		fmt.Println(string(rJSON))
		s := fmt.Sprint(string(rJSON))
		s = strings.Replace(s, `"`, `\"`, -1)
		fmt.Println(s)

		re, err := chaincode.ConvertRequest(string(rJSON))
		if err != nil {
			panic(err)
		}
		reJSON, err := json.Marshal(&re)
		fmt.Println(string(reJSON))
		fmt.Println(re.GetID())
		fmt.Println(re)
		fmt.Printf("re type: %T\n", re)

		var ri interface{}
		err = json.Unmarshal(rJSON, &ri)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%T\n", ri)
		// mapper := ri.(map[string]interface{})
		// rt := mapper["req_type"].(string)
		// fmt.Println(rt)
		// _, ok := ri.(chaincode.Request)
		// fmt.Println(ok)
		re.ChangeState(0)
		re.SetID("温家俊")
		fmt.Println("Altered ?,", re)

		reJSON, _ = json.Marshal(re)
		fmt.Println(string(reJSON))

		rep := chaincode.AttributeRes{
			ID:    "1",
			ReqID: "REQ-1",
			TargetTable: chaincode.PublicTable{
				ID:      "1",
				Area:    "Harbin",
				Year:    "2022",
				Month:   "12",
				Columns: []string{"列1", "列2"},
				// Data: [][]float64{
				// 	{1, 2, 3}, {10, 20, 30}, {100, 200, 300},
				// },
				NumOfObs:  3,
				Label:     []string{"Ob1", "Ob2", "Ob3"},
				TableType: "0",
			},
			Service:     "求和",
			Result:      1345.12,
			Description: "目标表格: 2022 年 12 月 Harbin 地区 全社会用电分类表; 目标属性: 列2; 计算服务: 求均值; 指标代码: 002; 指标名称: 全行业用电合计",
			ResType:     "0",
		}

		repJSON, _ := json.Marshal(rep)
		report, err := chaincode.ConvertReport(string(repJSON))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(report)

		t := chaincode.Table{
			ID:      "1",
			Area:    "Harbin",
			Year:    "2022",
			Month:   "12",
			Columns: []string{"列1", "列2"},
			Data: [][]float64{
				{1, 2, 3}, {10, 20, 30}, {100, 200, 300},
			},
			Label:     []string{"Ob1", "Ob2", "Ob3"},
			TableType: "0",
		}

		tJSON, _ := json.Marshal(t)
		fmt.Println(string(tJSON))
		s = fmt.Sprint(string(tJSON))
		s = strings.Replace(s, `"`, `\"`, -1)
		fmt.Println(s)
	}
}
