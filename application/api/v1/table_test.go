/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-30 12:52:28
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-01 22:46:44
 * @FilePath: \application\api\v1\table_test.go
 * @Description:
 */

package v1

import (
	"os"
	"testing"
)

func Test_formatFile(t *testing.T) {
	wd, _ := os.Getwd()
	r, _ := os.Open(wd + "/prod.xlsx")
	t.Error(r)
}
