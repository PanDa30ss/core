package redis

type RedisResult struct {
	Err    error
	Result interface{}
}
