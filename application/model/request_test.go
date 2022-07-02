/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-30 00:07:29
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 15:29:07
 * @FilePath: \application\model\request_test.go
 * @Description:
 */

package model

import (
	"encoding/json"
	"testing"
)

func TestSendRequest(t *testing.T) {
	pr := TableRequest{
		&BasicRequest{
			ReqType:       "2",
			TargetTableID: "1",
			RequestTime:   "2022-07-02 14:36",
			Service:       3,
			State:         0,
		},
	}
	prJSON, _ := json.Marshal(pr)
	t.Error(string(prJSON))
}
