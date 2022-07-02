/*
 * @Author: Wen Jiajun
 * @Date: 2022-07-02 01:51:44
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 12:49:22
 * @FilePath: \two-sorts-table\chaincode\id.go
 * @Description:
 */
package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) GetReqID(ctx contractapi.TransactionContextInterface) (string, error) {
	reqs, err := s.GetAllRequests(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(len(reqs)), nil
}

func (s *SmartContract) GetTableID(ctx contractapi.TransactionContextInterface) (string, error) {
	tables, err := s.GetAllTable(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(len(tables)), nil
}

// TODO: GetResID

func (s *SmartContract) GetResID(ctx contractapi.TransactionContextInterface) (string, error) {
	queryString := `{
		"selector": {
		   "req_type": {
			  "$exists": true
		   },
		   "state": {
			  "$eq": 200
		   }
		}
	 }`
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return "", err
	}
	defer resultsIterator.Close()

	count := 0
	for resultsIterator.HasNext() {
		resultsIterator.Next()
		count++
	}
	return fmt.Sprint(count), nil
}
