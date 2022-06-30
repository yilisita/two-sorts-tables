/*
 * @Author: Wen Jiajun
 * @Date: 2022-03-25 16:31:39
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-06-29 23:47:16
 * @FilePath: \application\error\errors_test.go
 * @Description:
 */

package error

import (
	"testing"
)

func TestGetErrMsg(t *testing.T) {
	t.Error(SUCCESS.Code())
	t.Errorf("%T", SUCCESS)
	var e error
	e = SUCCESS
	t.Errorf("%t, %T, %T", e, e, e.(ErrorCode))
}
