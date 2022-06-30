/*
 * @Author: Wen Jiajun
 * @Date: 2022-03-25 16:31:39
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-06-29 23:44:47
 * @FilePath: \application\error\errors.go
 * @Description:
 */

package error

const SUCCESS ErrorCode = 200

const (
	JSON_PARSE_ERROR ErrorCode = 500 + iota*10
	TX_SUBMITION_ERROR
	TX_EVALUATION_ERROR
	WALLET_CREATION_ERROR
)

const (
	REQ_NOT_EXIST ErrorCode = 30000 + iota*1000
	TABLE_NOT_EXIST
	RES_NOT_EXIST
	NO_REQ
	NO_TABLE
	NO_RES
)
const (
	FILE_PARSE_ERROR ErrorCode = 1000 + iota*100
	READ_EXCEL_ERROR
)

const (
	ADD_KEY_ERROR ErrorCode = 3000 + iota*1000
	DELETE_KEY_ERROR
	QUERY_KEY_ERROR
)

type ErrorCode uint

func (e ErrorCode) Error() string {
	return GetErrMsg(e)
}

func (e ErrorCode) Code() uint {
	return uint(e)
}

func (e ErrorCode) IsErr(i ErrorCode) bool {
	return e == i
}

var errCodeMsg = map[ErrorCode]string{
	SUCCESS:               "成功",
	JSON_PARSE_ERROR:      "JSON对象处理错误，请联系管理员",
	TX_SUBMITION_ERROR:    "提交交易失败，请联系管理员",
	TX_EVALUATION_ERROR:   "评估交易失败，请联系管理员",
	WALLET_CREATION_ERROR: "钱包创建失败，请联系管理员",
	FILE_PARSE_ERROR:      "文件解析失败，请联系管理员",
	READ_EXCEL_ERROR:      "Excel文件读取失败，请联系管理员",
	ADD_KEY_ERROR:         "插入密钥失败",
	DELETE_KEY_ERROR:      "删除密钥失败",
	QUERY_KEY_ERROR:       "查询密钥失败",
	REQ_NOT_EXIST:         "请求不存在",
	TABLE_NOT_EXIST:       "表格不存在",
	RES_NOT_EXIST:         "数据报告不存在",
	NO_REQ:                "当前无请求",
	NO_TABLE:              "当前无表格",
	NO_RES:                "当前无数据报告",
}

func GetErrMsg(e ErrorCode) string {
	return errCodeMsg[e]
}
