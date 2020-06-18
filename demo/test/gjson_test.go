package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/9299381/bingo"
	"github.com/tidwall/gjson"
	"testing"
)

func TestGjson(t *testing.T) {
	ctx := bingo.NewContext(context.Background())
	order := `{"id":"123123","amount":"2323","title":"中文"}`
	channel := `{"id":"456456","amount":"2323","title":"中文"}`
	user := `{"id":"789789","amount":"2323","title":"中文"}`

	ctx.Set("attach.order", gjson.Parse(order))
	ctx.Set("attach.user", gjson.Parse(user))
	ctx.Set("attach.channel", gjson.Parse(channel))

	jsonstr, _ := json.Marshal(ctx.Get("attach.order"))
	str := string(jsonstr)
	//title := gjson.GetBytes(jsonstr,order)
	fmt.Println(str)
}

func TestJsonMap(t *testing.T) {
	userStr := `{"user":{"id":"123123123","level":"P1601"}}`
	attach := make(map[string]interface{})
	_ = json.Unmarshal([]byte(userStr), &attach)
	fmt.Println(attach)
}
