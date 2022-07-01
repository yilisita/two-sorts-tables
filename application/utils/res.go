/*
 * @Author: Wen Jiajun
 * @Date: 2022-07-01 22:14:25
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-01 22:36:31
 * @FilePath: \application\utils\res.go
 * @Description:
 */
package utils

import e "app/error"

type Res struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewRes(err error) *Res {
	c, ok := err.(e.ErrCode)
	if ok {
		return &Res{
			Status:  c.Code(),
			Message: c.Error(),
		}
	}
	return &Res{
		Status:  0,
		Message: err.Error(),
	}
}

func (r *Res) WithData(data interface{}) *Res {
	r.Data = data
	return r
}
