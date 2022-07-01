/*
 * @Author: Wen Jiajun
 * @Date: 2022-07-01 22:14:25
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-01 22:37:44
 * @FilePath: \application\utils\res_test.go
 * @Description:
 */

package utils

import (
	e "app/error"
	"testing"
)

func TestNewRes(t *testing.T) {
	msg := NewRes(e.JSON_PARSE_ERROR)
	t.Error(msg)
}

func TestRes_WithData(t *testing.T) {
	res := NewRes(e.TX_EVALUATION_ERROR).WithData(100)
	t.Error(res)

}
