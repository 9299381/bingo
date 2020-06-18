package bingo

import (
	"context"
	"github.com/9299381/bingo/package/logger"
	"strings"
	"time"
)

type Context struct {
	context.Context
	Keys map[string]interface{}
	Log  ILogger
	*Request
	cleanupFuncs []func()
}

func NewContext(ctx context.Context) Context {
	return Context{
		Context:      ctx,
		Keys:         map[string]interface{}{},
		Log:          logger.GetInstance(),
		cleanupFuncs: nil,
	}
}

func WithContext(ctx Context) (Context, context.CancelFunc) {
	newCtx := Context{
		Keys: map[string]interface{}{},
		Log:  ctx.Log,
	}
	c, cancel := context.WithCancel(ctx.Context)
	wrappedCancel := func() {
		cancel()
		for _, f := range ctx.cleanupFuncs {
			f()
		}
	}
	newCtx.Context = c
	return newCtx, wrappedCancel
}
func WithContextValue(ctx Context) (Context, context.CancelFunc) {
	keys := make(map[string]interface{})
	for k, v := range ctx.Keys {
		keys[k] = v
	}
	newCtx := Context{
		Keys: keys,
		Log:  ctx.Log,
	}
	c, cancel := context.WithCancel(ctx.Context)
	wrappedCancel := func() {
		cancel()
		for _, f := range ctx.cleanupFuncs {
			f()
		}
	}
	newCtx.Context = c
	return newCtx, wrappedCancel
}

// OnCancel executes f when ctx is canceled.
func (ctx Context) OnCancel(f func()) {
	ctx.cleanupFuncs = append(ctx.cleanupFuncs, f)
}

//------------------set---------------------
func (ctx Context) Set(key string, value interface{}) {
	m := strings.Split(key, ".")
	if len(m) > 1 {
		ctx.setKeys(key, value)
	} else {
		ctx.setKey(key, value)
	}
}
func (ctx Context) setKey(key string, value interface{}) {
	ctx.Keys[key] = value
}
func (ctx Context) setKeys(key string, value interface{}) {
	keyMap := strings.Split(key, ".")
	pos := ctx.getPos(keyMap)
	if pos > 0 {
		ctx.modifyValue(keyMap, pos, value)
	} else {
		ret := ctx.setKeyToMap(keyMap[1:], value)
		ctx.Set(keyMap[0], ret)
	}
}

//------------------get--------------------

func (ctx Context) Get(key string) interface{} {
	m := strings.Split(key, ".")
	if len(m) > 1 {
		return ctx.getKeys(key)
	}
	if value, exists := ctx.getKey(key); exists {
		return value
	}
	return nil
}

func (ctx Context) getKey(key string) (value interface{}, exists bool) {
	value, exists = ctx.Keys[key]
	return
}
func (ctx Context) getKeys(key string) (ret interface{}) {
	keyMap := strings.Split(key, ".")
	ret = ctx.getKeyFromMap(keyMap, ctx.Keys)
	return
}

func (ctx Context) GetString(key string) (s string) {
	if val := ctx.Get(key); val != nil {
		s, _ = val.(string)
	}
	return
}
func (ctx Context) GetBool(key string) (b bool) {
	if val := ctx.Get(key); val != nil {
		b, _ = val.(bool)
	}
	return
}

func (ctx Context) GetInt(key string) (i int) {
	if val := ctx.Get(key); val != nil {
		i, _ = val.(int)
	}
	return
}
func (ctx Context) GetInt64(key string) (i64 int64) {
	if val := ctx.Get(key); val != nil {
		i64, _ = val.(int64)
	}
	return
}
func (ctx Context) GetFloat64(key string) (f64 float64) {
	if val := ctx.Get(key); val != nil {
		f64, _ = val.(float64)
	}
	return
}

func (ctx Context) GetTime(key string) (t time.Time) {
	if val := ctx.Get(key); val != nil {
		t, _ = val.(time.Time)
	}
	return
}
func (ctx Context) GetDuration(key string) (d time.Duration) {
	if val := ctx.Get(key); val != nil {
		d, _ = val.(time.Duration)
	}
	return
}

func (ctx Context) GetStringSlice(key string) (ss []string) {
	if val := ctx.Get(key); val != nil {
		ss, _ = val.([]string)
	}
	return
}

func (ctx Context) GetStringMap(key string) (sm map[string]interface{}) {
	if val := ctx.Get(key); val != nil {
		sm, _ = val.(map[string]interface{})
	}
	return
}
func (ctx Context) GetStringMapString(key string) (sms map[string]string) {
	if val := ctx.Get(key); val != nil {
		sms, _ = val.(map[string]string)
	}
	return
}
func (ctx Context) GetStringMapStringSlice(key string) (smss map[string][]string) {
	if val := ctx.Get(key); val != nil {
		smss, _ = val.(map[string][]string)
	}
	return
}

//--------------

func (ctx Context) getPos(keyMap []string) (pos int) {
	l := len(keyMap)
	pos = 0
	//是否有值,如果有值应该叠加上,从最后开始
	for i := 0; i < l; i++ {
		newMap := keyMap[:l-i]
		old := ctx.getKeyFromMap(newMap, ctx.Keys)
		if old != nil {
			pos = l - i
			break
		}
	}
	return
}
func (ctx Context) setKeyToMap(keyMap []string, value interface{}) (ret map[string]interface{}) {
	ret = make(map[string]interface{})
	if len(keyMap) == 1 {
		ret[keyMap[0]] = value
	} else if len(keyMap) > 1 {
		ret[keyMap[0]] = ctx.setKeyToMap(keyMap[1:], value)
	}
	return
}
func (ctx Context) getKeyFromMap(keyMap []string, valueMap interface{}) (ret interface{}) {
	m, ok := valueMap.(map[string]interface{})
	if ok {
		if len(keyMap) == 0 {
			ret = nil
		} else if len(keyMap) == 1 {
			ret = m[keyMap[0]]
		} else {
			ret = ctx.getKeyFromMap(keyMap[1:], m[keyMap[0]])
		}
	}
	return
}
func (ctx Context) modifyValue(keyMap []string, pos int, value interface{}) {
	//在函数调用时，像切片（slice）、字典（map）、
	// 接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递
	// （即使没有显式的指出指针）。
	ret, ok := ctx.getKeyFromMap(keyMap[:pos], ctx.Keys).(map[string]interface{})
	if ok {
		for v, k := range ctx.setKeyToMap(keyMap[pos:], value) {
			//注意这里的ret为指针,修改其值则c.key中值发生变化
			ret[v] = k
		}
	} else {
		if pos > 1 {
			ctx.modifyValue(keyMap, pos-1, value)
		}
	}
}
