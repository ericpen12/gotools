package call

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
)

func GetFuncName(fn interface{}) string {
	r := reflect.ValueOf(fn)
	if r.Kind() == reflect.Func {
		return runtime.FuncForPC(r.Pointer()).Name()
	}
	return ""
}

type Func any
type FuncSet map[string]Func

var fSet = FuncSet{}

func Register(fn ...Func) {
	for _, v := range fn {
		RegisterByAlias(GetFuncName(v), v)
	}
}

func RegisterByAlias(name string, fn Func) {
	if _, ok := fSet[name]; ok {
		panic("方法名：%s 已注册")
	}
	fSet[name] = fn
}

func RegisterInBatchByAlias(m FuncSet) {
	for k, v := range m {
		RegisterByAlias(k, v)
	}
}

type Result struct {
	Err    error
	Return []reflect.Value
}

func Call(fnPath string, params ...any) Result {
	fn := fSet[fnPath]
	funcVal := reflect.ValueOf(fn)

	if !funcVal.IsValid() {
		return result(nil, fmt.Errorf("function: %s not found", fnPath))
	}

	rt := reflect.TypeOf(fn)
	if len(params) != rt.NumIn() {
		return result(nil, fmt.Errorf("need %d params, but input %d params", rt.NumIn(), len(params)))
	}
	returns := funcVal.Call(buildParams(rt, params...))
	return result(returns, nil)
}
func result(data []reflect.Value, err error) Result {
	return Result{err, data}
}
func buildParams(rt reflect.Type, params ...any) []reflect.Value {
	var result []reflect.Value
	for i := 0; i < rt.NumIn(); i++ {
		t := reflect.New(rt.In(i))
		b, _ := json.Marshal(params[i])
		_ = json.Unmarshal(b, t.Interface())
		result = append(result, reflect.ValueOf(t.Elem().Interface()))
	}
	return result
}
