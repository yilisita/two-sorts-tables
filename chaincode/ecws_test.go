/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-28 01:44:57
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-06-28 01:47:45
 * @FilePath: \electricity-data-trade\chaincode\ecws_test.go
 * @Description:
 */
package chaincode

import (
	"fmt"
	"testing"
)

func Test_newCode(t *testing.T) {
	i := 10
	res := newCode(i)
	fmt.Println(res)
	t.Log("===============", res)
}
