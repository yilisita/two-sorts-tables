/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-28 01:48:49
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-06-29 15:48:42
 * @FilePath: \electricity-data-trade\chaincode\privateOrg1_test.go
 * @Description:
 */
/*
	这里编写数据拥有者的电力数据管理系统，包括增，删，查
	暂时先不做改的函数
*/
package chaincode

import (
	"encoding/json"
	"strings"
	"testing"
)

var testTable = Table{
	ID:      "1",
	Area:    "Harbin",
	Year:    "2022",
	Month:   "12",
	Columns: []string{"列1", "列2"},
	Data: [][]float64{
		{1, 2, 3}, {10, 20, 30}, {100, 200, 300},
	},
	Label:     []string{"Ob1", "Ob2", "Ob3"},
	TableType: ecws,
}

var r = EcwsRequest{
	BasicRequest: &BasicRequest{
		ID:            "REQ-1",
		ReqType:       ecws,
		Demander:      "Org2MSP",
		TargetTableID: "1",
		Service:       1,
		RequestTime:   "2022-06-27 21:48",
		State:         pending},

	AttributeID: 1,
	IndexCode:   2,
}

var pr = ProdRequest{
	BasicRequest: &BasicRequest{
		ID:            "REQ-2",
		ReqType:       prod,
		Demander:      "Org2MSP",
		TargetTableID: "1",
		Service:       1,
		RequestTime:   "2022-06-27 21:48",
		State:         pending,
	},
	AttributeID: 1,
}

var tr = TableRequest{
	BasicRequest: &BasicRequest{
		ID:            "REQ-2",
		ReqType:       prod,
		Demander:      "Org2MSP",
		TargetTableID: "1",
		Service:       1,
		RequestTime:   "2022-06-27 21:48",
		State:         pending,
	},
}

// {1 REQ-2 { 哈尔滨 2022 12 [列1 列2] 3 1 [Ob1 Ob2 Ob3]} 获取指标具体数值 20 目标表格: 2022 年 12 月 哈尔滨 地区 电力生产明细表; 目标属性: 列2; 计算服务: 获取指标具体数值; 指标代码: 001; 指标名称: 全社会用电总计 0}

var atRes = []AttributeRes{{
	ID:    "1",
	ReqID: "REQ-2",
	TargetTable: PublicTable{
		Area:      "哈尔滨",
		Year:      "2022",
		Month:     "12",
		Columns:   []string{"列1", "列2"},
		NumOfObs:  3,
		TableType: "1",
		Label:     []string{"Ob1", "Ob2", "Ob3"},
	},
	Description: "目标表格: 2022 年 12 月 哈尔滨 地区 电力生产明细表; 目标属性: 列2; 计算服务: 获取指标具体数值; 指标代码: 001; 指标名称: 全社会用电总计",
	Service:     "获取指标具体数值",
	Result:      20,
	ResType:     "0",
}}

// {2 REQ-3 {1 哈尔滨 2022 12 [列1 列2] [[1 2 3] [10 20 30] [100 200 300]] [Ob1 Ob2 Ob3] 1} 2}
var tres = TableRes{
	ID:    "2",
	ReqID: "REQ-3",
	Res: Table{
		ID:      "1",
		Area:    "Harbin",
		Year:    "2022",
		Month:   "12",
		Columns: []string{"列1", "列2"},
		Data: [][]float64{
			{1, 2, 3}, {10, 20, 30}, {100, 200, 300},
		},
		Label:     []string{"Ob1", "Ob2", "Ob3"},
		TableType: ecws,
	},
	ResType: table,
}

func Test_getReportDescription(t *testing.T) {
	// des := getReportDescription(testTable, r)
	// t.Errorf(des)

	// des = getReportDescription(testTable, pr)
	// t.Errorf(des)

	// des = getReportDescription(testTable, tr)
	// t.Errorf(des)

	atResJSON, _ := json.Marshal(atRes)
	s := string(atResJSON)
	t.Error(s)
	s = strings.Replace(s, `"`, `\"`, -1)
	t.Errorf(s)

	t.Error(atRes)

	var atB AttributeRes
	json.Unmarshal(atResJSON, &atB)
	t.Error(atB)

	tResJSON, _ := json.Marshal(tres)
	s = string(tResJSON)
	s = strings.Replace(s, `"`, `\"`, -1)
	t.Error(s)
}

func Test_getPublicTableFromPrivate(t *testing.T) {
	pt := getPublicTableFromPrivate(testTable)
	t.Error(pt)
}

func Test_getIndexValue(t *testing.T) {
	indexValue := getIndexValue(testTable, r)
	t.Error(indexValue)
}

func Test_mean(t *testing.T) {
	m := mean(testTable, pr)
	t.Error(m)
}

func Test_sum(t *testing.T) {
	s := sum(testTable, pr)
	t.Error(s)
}
