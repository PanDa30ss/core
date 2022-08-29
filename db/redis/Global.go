package redis

import "github.com/garyburd/redigo/redis"

func MakeRedisCommand(callback RedisCallBack, params ...interface{}) *redisCommand {
	ret := &redisCommand{}
	ret.callback = callback
	ret.params = params
	return ret
}

func Strings(reply interface{}) ([]string, error) {
	var err error
	return redis.Strings(reply, err)
}
func Ints(reply interface{}) ([]int, error) {
	var err error
	return redis.Ints(reply, err)
}
func Int(reply interface{}) (int, error) {
	var err error
	return redis.Int(reply, err)
}
