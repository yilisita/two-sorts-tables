package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Report interface {
	GetReqID() string
	GetID() string
}

// 数据提供方返回响应的结构
/*---- 六个字段 ----*/
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

func NewPublicTableView(pt PublicTable) *PublicTableView {
	return &PublicTableView{
		ID:           pt.ID,
		Area:         pt.Area,
		Year:         pt.Year,
		Month:        pt.Month,
		Columns:      pt.Columns,
		NumOfObs:     pt.NumOfObs,
		TableTypeStr: tableTypeMapper[pt.TableType],
		Label:        pt.Label,
	}
}

func NewAttributeResView(r AttributeRes) *AttributeResView {
	return &AttributeResView{
		ID:          r.ID,
		ReqID:       r.ReqID,
		TargetTable: *NewPublicTableView(r.TargetTable),
		Service:     r.GetReqID(),
		Result:      r.Result,
		Description: r.GetReqID(),
		ResTypeStr:  reqTypeMapper[r.ResType],
	}
}

func (a AttributeRes) GetReqID() string {
	return a.ReqID
}

func (t TableRes) GetReqID() string {
	return t.ReqID
}

func (a AttributeRes) GetID() string {
	return a.ID
}

func (t TableRes) GetID() string {
	return t.ID
}

var org2Collection = "Org2PrivateCollection"

func (s *SmartContract) createReport(ctx contractapi.TransactionContextInterface, rep Report) (string, error) {
	responseJSON, err := json.Marshal(rep)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutPrivateData(org2Collection, rep.GetID(), responseJSON)
	if err != nil {
		return "", err
	}

	return rep.GetID(), nil
}

// ReadAsset returns the asset stored in the world state with given id.
// 读取方法有问题
// TODO: Error: json.Unmarshal(nil) ?
func (s *SmartContract) ReadReport(ctx contractapi.TransactionContextInterface, id string) (interface{}, error) {
	responseJSON, err := ctx.GetStub().GetPrivateData(org2Collection, id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if responseJSON == nil {
		return nil, fmt.Errorf("the response %s does not exist", id)
	}
	var res Report
	err = json.Unmarshal(responseJSON, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// OK
func (s *SmartContract) GetAllReports(ctx contractapi.TransactionContextInterface) ([]*AttributeResView, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	queryString := fmt.Sprintf(`{"selector":{"res_type":{"$in": ["%s", "%s"]}}}`, ecws, prod)
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(org2Collection, queryString)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var responses []*AttributeResView
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var temp AttributeRes
		err = json.Unmarshal(queryResponse.Value, &temp)
		if err != nil {
			return nil, err
		}

		responses = append(responses, NewAttributeResView(temp))

	}

	return responses, nil
}

// OK
func (s *SmartContract) ReadPurchasedTable(ctx contractapi.TransactionContextInterface) ([]*Table, error) {
	queryString := fmt.Sprintf(`{"selector":{"res_type":{"$in": ["%s"]}}}`, table)
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(org2Collection, queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	defer resultsIterator.Close()

	var tables []*Table
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var temp TableRes
		err = json.Unmarshal(queryResponse.Value, &temp)
		if err != nil {
			return nil, err
		}

		tables = append(tables, &(temp.Res))

	}

	return tables, nil
}

// OK
func ConvertReport(repStr string) (Report, error) {
	var r interface{}
	err := json.Unmarshal([]byte(repStr), &r)
	if err != nil {
		return nil, err
	}

	rm := r.(map[string]interface{})
	switch rm["res_type"].(string) {
	case ecws:
		var er AttributeRes
		err = json.Unmarshal([]byte(repStr), &er)
		if err != nil {
			return nil, err
		}
		return er, nil
	case prod:
		var pr AttributeRes
		err := json.Unmarshal([]byte(repStr), &pr)
		if err != nil {
			return nil, err
		}
		return pr, nil
	case table:
		var tr TableRes
		err := json.Unmarshal([]byte(repStr), &tr)
		if err != nil {
			return nil, err
		}
		return tr, nil
	}
	return nil, fmt.Errorf("Invalid Request Type")
}
