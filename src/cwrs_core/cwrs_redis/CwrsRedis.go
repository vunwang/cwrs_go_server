package cwrs_redis

import (
	"cwrs_go_server/src/cwrs_core/cwrs_viper"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var GlobalRedis *redis.Client

func init() {
	InitRedis()
}

func InitRedis() {
	//连接redis
	conn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cwrs_viper.GlobalViper.GetString("redis.path"), cwrs_viper.GlobalViper.GetInt("redis.port")),
		Password: cwrs_viper.GlobalViper.GetString("redis.password"),
		DB:       cwrs_viper.GlobalViper.GetInt("redis.db"),
	})

	GlobalRedis = conn

	fmt.Println("Redis Initialize OK !")

	//// 设置键值对并设置过期时间
	//err := GlobalRedis.Set(ctx, "key", "value", 10*time.Second).Err()
	//if err != nil {
	//	fmt.Printf("Set failed: %v\n", err)
	//	return
	//}
	//
	//val, err := GlobalRedis.Get(ctx, "key").Result()
	//if err == redis.Nil {
	//	fmt.Println("key does not exist")
	//} else if err != nil {
	//	fmt.Printf("Get failed: %v\n", err)
	//	return
	//} else {
	//	fmt.Println("key:", val)
	//}
}
