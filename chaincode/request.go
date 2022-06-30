package chaincode

import (
	//"encoding/json"
	//"fmt"
	//"log"

	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

const (
	pending uint = iota * 100
	refused
	completed
)

// ResType, ReqType, TableType
const (
	ecws  string = "0"
	prod  string = "1"
	table string = "2"
)

var reqTypeMapper = map[string]string{
	ecws:  "全社会用电查询",
	prod:  "电力生产情况查询",
	table: "整表购买",
}
var stateMapper = map[uint]string{
	pending:   "待处理",
	refused:   "已拒绝处理",
	completed: "已完成",
}

const reqPrefix = "REQ-"

var reqID = 1

// 数据需求方发送请求的结构
/*--- 七个字段 ----*/
type Request interface {
	SetID(string)
	GetID() string
	SetDemander(string)
	GetState() uint
	GetTargetTableID() string
	GetService() int
	ChangeState(uint)
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

func (b *BasicRequest) SetID(id string) {
	b.ID = id
}

func (b *BasicRequest) SetDemander(d string) {
	b.Demander = d
}

func (b *BasicRequest) GetID() string {
	return b.ID
}

func (b *BasicRequest) GetState() uint {
	return b.State
}

func (b *BasicRequest) GetTargetTableID() string {
	return b.TargetTableID
}

func (b *BasicRequest) GetService() int {
	return b.Service
}

func (b *BasicRequest) ChangeState(s uint) {
	b.State = s
}
func newReqID(id string) string {
	return reqPrefix + id
}

// 用来购买字段的
// OKOK
func (s *SmartContract) SendRequest(ctx contractapi.TransactionContextInterface, reqStr string) (string, error) {
	request, err := ConvertRequest(reqStr)
	if err != nil {
		return "", err
	}

	// 获取发起者的MSPID，e.g. Org2MSP
	demander, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	request.SetDemander(demander)
	request.SetID(newReqID(fmt.Sprint((reqID))))

	log.Println(request)

	requestJSON, err := json.Marshal(&request)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(request.GetID(), requestJSON)
	if err != nil {
		return "", err
	}
	reqID++
	return fmt.Sprint(reqID - 1), nil
}

// OK
func (s *SmartContract) ReadAllRequest(ctx contractapi.TransactionContextInterface) ([]RequestView, error) {
	queryString := fmt.Sprintf(`{"selector":{"req_type":{"$in": ["%s", "%s", "%s"]}}}`, ecws, prod, table)
	requestIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer requestIterator.Close()

	var reqs []RequestView
	for requestIterator.HasNext() {
		requestJSON, err := requestIterator.Next()
		if err != nil {
			return nil, err
		}
		temp, err := ConvertRequest(string(requestJSON.Value))
		if err != nil {
			return nil, err
		}
		t, _ := s.ReadPublicTableByID(ctx, temp.GetTargetTableID())
		var tmpv = NewRequestView(*t, temp)
		reqs = append(reqs, tmpv)
	}
	return reqs, nil
}

// {"selector":{"req_type":{"$in": ["%s", "%s", "%s"]}}}
// Ok
func (s *SmartContract) GetAllRequests(ctx contractapi.TransactionContextInterface) ([]interface{}, error) {
	queryString := fmt.Sprintf(`{"selector":{"req_type":{"$in": ["%s", "%s", "%s"]}}}`, ecws, prod, table)
	requestIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer requestIterator.Close()

	var reqs []interface{}
	for requestIterator.HasNext() {
		requestJSON, err := requestIterator.Next()
		if err != nil {
			return nil, err
		}
		temp, err := ConvertRequest(string(requestJSON.Value))
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, temp)
	}
	return reqs, nil
}

func (s *SmartContract) HandleAll(ctx contractapi.TransactionContextInterface) ([]interface{}, error) {
	var requests, err = s.GetAllRequests(ctx)
	if err != nil {
		return nil, err
	}
	var res []interface{}
	for _, re := range requests {
		// Skip handled and refused request
		requestJSON, err := json.Marshal(re)
		if err != nil {
			return nil, err
		}
		request, err := ConvertRequest(string(requestJSON))
		if request.GetState() != pending {
			continue
		}
		report, err := s.HandleSingle(ctx, request.GetID()) // "REQ-1"
		if err != nil {
			return nil, err
		}
		res = append(res, report)
	}
	return res, nil
}

// REQ-1
// OK
func (s *SmartContract) HandleSingle(ctx contractapi.TransactionContextInterface, requestID string) (interface{}, error) {
	// 读取这个请求
	reqID := requestID
	requestJSON, err := ctx.GetStub().GetState(reqID)
	if err != nil {
		return nil, err
	}
	if requestJSON == nil {
		return nil, fmt.Errorf("No such request.")
	}

	// 反序列化请求JSON => struct
	request, err := ConvertRequest(string(requestJSON))
	if err != nil {
		return nil, err
	}

	// 如果已经解决了，则结束
	if request.GetState() != pending {
		return nil, fmt.Errorf("Unable to handle, status: %v", request.GetState())
	}

	targetTable, err := s.ReadMyTableByID(ctx, request.GetTargetTableID())
	if err != nil {
		return nil, err
	}

	return deriveReport(request, targetTable), nil

}

// TODO: test
func (s *SmartContract) RefuseRequest(ctx contractapi.TransactionContextInterface, id string) error {
	requestJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return err
	}

	if requestJSON == nil {
		return fmt.Errorf("Nil request %v", id)
	}

	r, err := ConvertRequest(string(requestJSON))
	r.ChangeState(refused)

	rJSON, err := json.Marshal(r)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(r.GetID(), rJSON)
}

// Ok
func (s *SmartContract) SendReport(ctx contractapi.TransactionContextInterface, reportStr string) error {
	// 在private data 1上做了range query就不能在private data 2上做write，所有handleRequest和sendResponse必须分开定义
	var reports interface{}
	err := json.Unmarshal([]byte(reportStr), &reports)
	if err != nil {
		return err
	}
	rs := reports.([]interface{})
	for _, v := range rs {
		vJSON, err := json.Marshal(v)
		if err != nil {
			return err
		}
		rep, err := ConvertReport(string(vJSON))
		if err != nil {
			return err
		}
		_, err = s.createReport(ctx, rep)
		if err != nil {
			return err
		}

		// 将处理过的请求标记为已经完成
		requestJSON, err := ctx.GetStub().GetState(rep.GetReqID())
		if err != nil {
			return err
		}
		request, err := ConvertRequest(string(requestJSON))
		if err != nil {
			return err
		}

		// 标记为已经完成处理该请求
		request.ChangeState(completed)

		requestJSON, err = json.Marshal(request)
		if err != nil {
			return err
		}

		// 将完成状态的请求放回请求ledger中
		err = ctx.GetStub().PutState(rep.GetReqID(), requestJSON)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
func (s *SmartContract) HandleTableBuy(ctx contractapi.TransactionContextInterface, id string) (*TableRes, error) {
	// 读取这个请求
	reqID := newReqID(id)
	requestJSON, err := ctx.GetStub().GetState(reqID)
	if err != nil {
		return nil, err
	}
	if requestJSON == nil {
		return nil, fmt.Errorf("No such request.")
	}

	// 反序列化请求JSON => struct
	var request Request
	err = json.Unmarshal(requestJSON, &request)
	if err != nil {
		return nil, err
	}
	// 如果已经解决了，则结束
	if request.State != completed {
		return nil, fmt.Errorf("Unable to handle, status: %v", request.State)
	}

	targetTable, err := s.ReadMyTableByID(ctx, request.TargetTableID)
	if err != nil {
		return nil, err
	}

	var res = TableRes{
		ID:      fmt.Sprint(resID),
		Res:     targetTable,
		ResType: "Table",
	}

	return &res, nil
}

func (s *SmartContract) HanldeAllTableBuy(ctx contractapi.TransactionContextInterface) ([]*TableRes, error) {
	queryString := fmt.Sprintf(`{"selector":{"req_type":"%s"}}`, "table")
	reqIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer reqIterator.Close()

	var res []*TableRes
	for reqIterator.HasNext() {
		tmp, err := reqIterator.Next()
		if err != nil {
			return nil, err
		}
		var r Request
		err = json.Unmarshal(tmp.Value, &r)
		if err != nil {
			return nil, err
		}
		tRes, err := s.HandleTableBuy(ctx, r.GetID())
		if err != nil {
			return nil, err
		}
		res = append(res, tRes)
	}
	return res, nil
}
*/

// ID 和 TableRes
func ConvertRequest(reqStr string) (Request, error) {
	var r interface{}
	err := json.Unmarshal([]byte(reqStr), &r)
	if err != nil {
		return nil, err
	}

	rm := r.(map[string]interface{})
	switch rm["req_type"].(string) {
	case ecws:
		var er EcwsRequest
		err = json.Unmarshal([]byte(reqStr), &er)
		if err != nil {
			return nil, err
		}
		return er, nil
	case prod:
		var pr ProdRequest
		err := json.Unmarshal([]byte(reqStr), &pr)
		if err != nil {
			return nil, err
		}
		return pr, nil
	case table:
		var tr TableRequest
		err := json.Unmarshal([]byte(reqStr), &tr)
		if err != nil {
			return nil, err
		}
		return tr, nil
	}
	return nil, fmt.Errorf("Invalid Request Type")
}

func deriveReport(r Request, t Table) Report {
	switch r.(type) {
	case EcwsRequest:
		if f, ok := serviceFucMapper[r.GetService()]; ok {
			var res = f(t, r)
			var ar = AttributeRes{
				ID:          fmt.Sprint(resID),
				ReqID:       r.GetID(),
				TargetTable: getPublicTableFromPrivate(t),
				Service:     serviceNameMapper[r.GetService()],
				Result:      res,
				ResType:     ecws,
				Description: getReportDescription(t, r),
			}
			return ar
		}
	case ProdRequest:
		if f, ok := serviceFucMapper[r.GetService()]; ok {
			var res = f(t, r)
			var ar = AttributeRes{
				ID:          fmt.Sprint(resID),
				ReqID:       r.GetID(),
				TargetTable: getPublicTableFromPrivate(t),
				Service:     serviceNameMapper[r.GetService()],
				Result:      res,
				ResType:     prod,
				Description: getReportDescription(t, r),
			}
			return ar
		}
	case TableRequest:
		var tr = TableRes{
			ID:      fmt.Sprint(resID),
			ReqID:   r.GetID(),
			Res:     t,
			ResType: table,
		}
		return tr
	}
	return nil
}

func NewRequestView(t PublicTable, r Request) RequestView {
	var res = RequestView{
		ID:            r.GetID(),
		TargetTableID: r.GetTargetTableID(),
		Service:       serviceNameMapper[r.GetService()],
		// RequestTime   string `json:"request_time"`
		State: stateMapper[r.GetState()],
		// Attribute     string `json:"attribute"`
		// Index         string `json:"index"`
	}

	switch r.(type) {
	case EcwsRequest:
		var er = r.(EcwsRequest)
		res.Demander = er.Demander
		res.Attribute = t.Columns[er.AttributeID] // TODO: Table Attribute Name
		res.Index = indexMapper[newCode(er.IndexCode)]
		res.RequestTime = er.RequestTime
		res.ReqTypeStr = reqTypeMapper[er.ReqType]
	case ProdRequest:
		var pr = r.(ProdRequest)
		res.Demander = pr.Demander
		res.Attribute = t.Columns[pr.AttributeID] // TODO: Table Attribute Name
		res.Index = "无"
		res.RequestTime = pr.RequestTime
		res.ReqTypeStr = reqTypeMapper[pr.ReqType]
	case TableRequest:
		var tr = r.(TableRequest)
		res.Demander = tr.Demander
		res.Attribute = "无" // TODO: Table Attribute Name
		res.Index = "无"
		res.RequestTime = tr.RequestTime
		res.ReqTypeStr = reqTypeMapper[tr.ReqType]
	}
	return res
}
