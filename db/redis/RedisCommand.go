package redis

type RedisCallBack func(result *RedisResult, params ...interface{})
type redisCommand struct {
	callback RedisCallBack
	result   *RedisResult
	params   []interface{}
}

func (this *redisCommand) Execute() {
	if this.callback == nil {
		return
	}
	this.callback(this.result, this.params...)
}

func defaultFunc(result *RedisResult, params ...interface{}) {
}
