package http

import (
	"encoding/json"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/util"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Post(host, service string, params map[string]interface{}) (ret *bingo.Response) {

	path := "http://" + host + "/" + strings.Replace(service, ".", "/", -1)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm(path, util.FormEncode(params))
	if err != nil {
		ret = bingo.Failed(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	response := &bingo.Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		ret = bingo.Failed(err)
	} else {
		m := response.Data.(map[string]interface{})
		m["call_method"] = "http"
		response.Data = m
		ret = response
	}
	return
}
