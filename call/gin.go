package call

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

func Driver(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			response(c, nil, err.(error))
		}
	}()
	data, err := callMethod(c)
	response(c, data, err)
}

func callMethod(c *gin.Context) (any, error) {
	method, err := getMethod(c.Query("method"))
	if err != nil {
		return nil, err
	}
	params := make(map[string]any)
	_ = c.ShouldBind(&params)
	var ret Result
	if len(params) == 0 {
		ret = Call(method)
	} else {
		ret = Call(method, params)
	}
	if ret.Err != nil {
		return nil, ret.Err
	}
	return formatResult(ret.Return)
}

func getMethod(method string) (string, error) {
	if _, ok := fSet[method]; ok {
		return method, nil
	}
	var contains []string
	for k := range fSet {
		if strings.HasSuffix(k, method) {
			contains = append(contains, k)
		}
	}
	if len(contains) == 0 {
		return "", fmt.Errorf("method: %s not found", method)
	}
	if len(contains) > 1 {
		return "", fmt.Errorf("many method matched, what do you mean:\n %s", strings.Join(contains, "\n"))
	}
	return contains[0], nil
}

func formatResult(list []reflect.Value) (interface{}, error) {
	var (
		result interface{}
		err    error
	)
	for _, item := range list {
		if item.Type().String() == "error" {
			err, _ = item.Interface().(error)
		} else {
			result = item.Interface()
		}
	}
	return result, err
}

type ginResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func response(ginCtx *gin.Context, data interface{}, err error) {
	resp := ginResponse{
		Data: data,
	}
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = -1
	}
	ginCtx.JSON(200, resp)
}
