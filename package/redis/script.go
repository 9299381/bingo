package redis

func luaLockScript() string {
	script := `
local lock_key = KEYS[1]
local lock_value = KEYS[2]
local lock_delay = KEYS[3]
local result = redis.call('SETNX',lock_key,lock_value)
if result == 1
then
redis.call('SETEX',lock_key,lock_delay,lock_value)
return result
else
return result
end
`
	return script
}

func luaUnLockScript() string {
	script := `
local lock_key = KEYS[1]
local result = redis.call('DEL',lock_key)
return result
`
	return script
}
