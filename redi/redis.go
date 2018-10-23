package redi

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"github.com/name5566/leaf/log"
)

var rediPool * redis.Pool

func newPool(addr string, pwd string) *redis.Pool {
	if addr == "" || pwd == "" {
		return nil
	}

	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH", pwd); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},
	}
}

//var (
//	pool *redis.Pool
//	redisServer = flag.String("redisServer", ":6379", "")
//)

//func main() {
//	flag.Parse()
//	pool = newPool(*redisServer)
//
//}

func RedisInit(url string, pwd string) *redis.Pool{
	rediPool := newPool(url, pwd)
	if rediPool == nil {
		log.Fatal("rdisInit fail")
	}

	con, err := rediPool.Dial()
	if con == nil {
		log.Fatal("rediPool dial error-%v", err)
	}

	return rediPool
	//c := rediPool.Get()
	//defer c.Close()
	//
	//if _, err := c.Do("AUTH", conf.Server.RedisPwdLogin); err != nil {
	//	log.Fatal("redis AUTH failed pd:b840fc02d524045429941cc15f59e41cb7be6c5288")
	//}
	//
	//
	////c.Do("hgetall, "myname", "superL")
	//table, err := c.Do("hget", "client1", "name")
	//if err != nil {
	//	log.Debug("redis get failed", err)
	//} else {
	//	t2 := table.([]byte)
	//	userName := string(t2)
	//	log.Debug("t2", userName)
	//	//var t1 = table.([]interface{})
	//	//for index, data := range t1 {
	//	//	log.Debug("get mykey:%v%v\n", index, string(data.([]byte)))
	//	//}
	//
	//}


	//string(table[0])

	//username, err := redis.String(c.Do("Get", "mykey"))
	//if err != nil {
	//	log.Debug("redis get failed", err)
	//} else {
	//	log.Debug("get mykey:%v\n", username)
	//}
}