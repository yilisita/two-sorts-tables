/*
 * @Author: Wen Jiajun
 * @Date: 2022-03-25 16:31:39
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-01 22:01:20
 * @FilePath: \application\error\errors_test.go
 * @Description:
 */

package error

import (
	"testing"
)

func TestGetErrMsg(t *testing.T) {
	t.Error(SUCCESS)
	t.Error(SUCCESS.Code())
	t.Error(JSON_PARSE_ERROR.Error())
}
