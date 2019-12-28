package request

import (
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type requestData struct {
	data map[string]string
}

func Data(r *http.Request) map[string]string {

	handler := new(requestData)
	handler.init()
	handler.setForm(r)
	return handler.Datas()
}

func (handler *requestData) init() {

	handler.data = make(map[string]string)
}

func (handler *requestData) setForm(r *http.Request) {

	r.ParseForm() //解析参数，默认是不会解析的

	formParams := handler.sortRequestParams(r.Form)
	for key, value := range formParams {
		handler.data[key] = value
	}
}

func (handler *requestData) Datas() map[string]string {
	return handler.data
}

func (handler *requestData) sortRequestParams(params url.Values) map[string]string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	formParams := make(map[string]string)

	for _, key := range keys {
		formParams[key] = strings.Join(params[key], "")
	}
	return formParams
}
