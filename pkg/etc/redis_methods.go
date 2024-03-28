package etc

import (
	"api_exam/api/handlers/models"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}

func GetRedis(key string) (*models.UserDetail, error) {
	rdb := ConnectToRedis()
	defer rdb.Close()

	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		log.Println("Error getting user from redis")
		return nil, err
	}

	var redisInfo models.UserDetail
	err = json.Unmarshal([]byte(val), &redisInfo)
	if err != nil {
		log.Println("Error marshaling value to interface")
		return nil, err
	}

	return &redisInfo, nil
}

func SetRedis(key string, value interface{}) error {
	rdb := ConnectToRedis()
	defer rdb.Close()

	userByte, err := json.Marshal(value)
	if err != nil {
		log.Println("Error marshalling interface to json as user")
		return err
	}

	_, err = rdb.Set(context.Background(), key, userByte, time.Minute*10).Result()
	if err != nil {
		log.Println("Error saving code and user info to redis", err.Error())
		return err
	}

	return nil
}
