package redis

import (
	log "github.com/PanDa30ss/core/logManager"
	"github.com/PanDa30ss/core/service"

	"github.com/garyburd/redigo/redis"
)

var defaultMaxConn int = 0
var defaultMaxIdle int = 16

type Redis struct {
	db      *redis.Pool
	maxConn int
	maxIdle int
	url     string
	isOpen  bool
}

func (this *Redis) InitDB() {
	this.url = ""
	this.maxConn = defaultMaxConn
	this.maxIdle = defaultMaxIdle
	this.db = nil
	this.isOpen = false

}
func (this *Redis) SetUrl(url string) {
	this.url = url
}
func (this *Redis) SetMaxConn(maxConn int) {
	this.maxConn = maxConn
}
func (this *Redis) SetMaxIdle(maxIdle int) {
	this.maxIdle = maxIdle
}
func (this *Redis) Open() bool {
	if this.url == "" {
		return false
	}

	this.db = &redis.Pool{ //实例化一个连接池
		MaxIdle:     this.maxIdle, //最初的连接数量
		MaxActive:   this.maxConn, //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300,          //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", this.url)
		},
	}
	this.isOpen = true
	log.Info("redis open")
	return true
}
func (this *Redis) Query(cmd *redisCommand, op string, params ...interface{}) {
	go this.doQuery(cmd, op, params...)
}
func (this *Redis) doQuery(cmd *redisCommand, op string, params ...interface{}) {
	c := this.db.Get()
	defer c.Close()
	ret, err := c.Do(op, params...)
	cmd.result = &RedisResult{err, ret}
	service.Post(cmd)
}

// func (this *Redis) Set(key string, value interface{}, cmd *redisCommand) {
// 	go this.doSet(key, value, cmd)
// }
// func (this *Redis) Get(key string, cmd *redisCommand) {
// 	go this.doGet(key, cmd)
// }

// func (this *Redis) doSet(key string, value interface{}, cmd *redisCommand) {
// c := this.db.Get()
// defer c.Close() //函数运行结束 ，把连接放回连接池

// 	ret, err := c.Do("Set", key, value)
// 	fmt.Println(ret)
// 	cmd.result = &RedisResult{err, ret}
// 	service.Post(cmd)

// }
// func (this *Redis) doGet(key string, cmd *redisCommand) {
// 	c := this.db.Get()
// 	defer c.Close() //函数运行结束 ，把连接放回连接池
// 	ret, err := c.Do("Get", key)
// 	fmt.Println(ret)
// 	cmd.result = &RedisResult{err, ret}
// 	service.Post(cmd)
// }
