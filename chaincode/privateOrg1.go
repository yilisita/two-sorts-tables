/*
	这里编写数据拥有者的电力数据管理系统，包括增，删，查
	暂时先不做改的函数
*/
package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const myCollection = "Org1PrivateCollection"

// type ECWS struct {
// 	Code                         int     `json:"code"`
// 	IndexName                    string  `json:"index_name"`
// 	NumOfUsers                   int     `json:"num_of_users"`
// 	UserInstalledCapacity        int     `json:"user_installed_capacity"`
// 	ThisMonthPowerConsumption    float64 `json:"this_month_power_consumption"`
// 	SMLYPowerConsumption         float64 `json:"smly_power_consumption"`
// 	AccumulativePowerConsumption float64 `json:"ACC_power_consumption"`
// 	LYAccPowerConsumption        float64 `json:"ly_acc_power_consumption"`
// }

var id int = 0
var indexID int = 0

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

var tableTypeMapper = map[string]string{
	ecws: "全社会用电分类表",
	prod: "电力生产明细表",
}

// Table用来描述一张表格的信息
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

func (s *SmartContract) InsertATable(ctx contractapi.TransactionContextInterface, tableStr string) (int, error) {
	var table Table
	err := json.Unmarshal([]byte(tableStr), &table)
	if err != nil {
		return -1, err
	}

	table.ID = fmt.Sprint(id)
	tableJSON, err := json.Marshal(table)
	if err != nil {
		return -1, err
	}

	err = ctx.GetStub().PutPrivateData(myCollection, fmt.Sprint(id), tableJSON)
	if err != nil {
		return -1, err
	}
	thisID := id

	var publicTable = PublicTable{
		ID:        table.ID,
		Area:      table.Area,
		Year:      table.Year,
		Month:     table.Month,
		Columns:   table.Columns,
		NumOfObs:  len(table.Label),
		TableType: table.TableType,
		Label:     table.Label,
	}

	publicTableJSON, err := json.Marshal(publicTable)
	if err != nil {
		return -1, err
	}

	err = ctx.GetStub().PutState(table.ID, publicTableJSON)
	if err != nil {
		return -1, err
	}
	id++
	return thisID, nil
}

func (s *SmartContract) ReadPublicTableByID(ctx contractapi.TransactionContextInterface, tableID string) (*PublicTable, error) {
	tableJSON, err := ctx.GetStub().GetState(tableID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from private collection: %v", err)
	}
	if tableJSON == nil {
		return nil, fmt.Errorf("Table %s does not exist", tableID)
	}
	var table PublicTable
	err = json.Unmarshal(tableJSON, &table)
	if err != nil {
		return nil, err
	}
	return &table, nil
}

// {"selector":{"table_type":{"$in": ["%s", "%s"]}}}
// OK
func (s *SmartContract) ReadAllPublicTable(ctx contractapi.TransactionContextInterface) ([]*PublicTable, error) {
	queryString := fmt.Sprintf(`{"selector":{"table_type":{"$in": ["%s", "%s"]}}}`, ecws, prod)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var res []*PublicTable
	for resultsIterator.HasNext() {
		tmp, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var t PublicTable
		err = json.Unmarshal(tmp.Value, &t)
		if err != nil {
			return nil, err
		}
		res = append(res, &t)
	}
	return res, nil
}

// 读取表格数据，如果是加密的数据，那么读取的结果保持加密状态。
// OK
func (s *SmartContract) ReadMyTableByID(ctx contractapi.TransactionContextInterface, tableID string) (Table, error) {
	tableJSON, err := ctx.GetStub().GetPrivateData(myCollection, tableID)
	if err != nil {
		return Table{}, fmt.Errorf("failed to read from private collection: %v", err)
	}
	if tableJSON == nil {
		return Table{}, fmt.Errorf("Table %s does not exist", tableID)
	}
	var table Table
	err = json.Unmarshal(tableJSON, &table)
	if err != nil {
		return Table{}, err
	}
	return table, nil
}

// 内部函数，用来检测表格是否存在
func (s *SmartContract) myTableExists(ctx contractapi.TransactionContextInterface, tableID string) (bool, error) {
	tableJSON, err := ctx.GetStub().GetPrivateData(myCollection, tableID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return tableJSON != nil, nil
}

// Ok
func (s *SmartContract) GetAllTable(ctx contractapi.TransactionContextInterface) ([]*Table, error) {
	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange(myCollection, "", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var res []*Table
	for resultsIterator.HasNext() {
		tmp, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var t Table
		err = json.Unmarshal(tmp.Value, &t)
		if err != nil {
			return nil, err
		}
		res = append(res, &t)
	}
	return res, nil
}

var serviceNameMapper = map[int]string{
	0: "求和",
	1: "求均值",
	2: "获取指标具体数值",
	3: "购买表格",
}

var serviceFucMapper = map[int]func(Table, Request) float64{
	0: sum,
	1: mean,
	2: getIndexValue,
}

// Pass
func sum(table Table, request Request) float64 {
	pr := request.(ProdRequest)
	var data = table.Data[pr.AttributeID]
	var res float64 = 0
	for _, v := range data {
		res += v
	}
	return res
}

// Pass
func mean(table Table, request Request) float64 {
	var s = sum(table, request)
	return s / float64(len(table.Label))
}

// Pass
func getIndexValue(table Table, request Request) float64 {
	er := request.(EcwsRequest)
	return table.Data[er.AttributeID][er.IndexCode]
}

// Pass
func getPublicTableFromPrivate(t Table) PublicTable {
	return PublicTable{
		Area:      t.Area,
		Year:      t.Year,
		Month:     t.Month,
		Columns:   t.Columns,
		NumOfObs:  len(t.Label),
		Label:     t.Label,
		TableType: t.TableType,
	}
}

// Pass
func getReportDescription(t Table, req Request) string {
	switch req.(type) {
	case EcwsRequest:
		var er = req.(EcwsRequest)
		index := newCode(er.IndexCode)
		return fmt.Sprintf("目标表格: %v 年 %v 月 %v 地区 %v; 目标属性: %v; 计算服务: %v; 指标代码: %v; 指标名称: %v",
			t.Year, t.Month, t.Area, tableTypeMapper[t.TableType], t.Columns[er.AttributeID],
			serviceNameMapper[er.Service], index, indexMapper[index])
	case ProdRequest:
		var pr = req.(ProdRequest)
		return fmt.Sprintf("目标表格: %v 年 %v 月 %v 地区 %v; 目标属性: %v; 计算服务: %v",
			t.Year, t.Month, t.Area, tableTypeMapper[t.TableType], t.Columns[pr.AttributeID],
			serviceNameMapper[pr.Service])
	case TableRequest:
		return fmt.Sprintf("购买目标表格: %v 年 %v 月 %v 地区 %v",
			t.Year, t.Month, t.Area, tableTypeMapper[t.TableType])
	}
	return "解析错误"
}
