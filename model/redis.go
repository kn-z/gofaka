package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func InitRedis() {
	fmt.Println("Redis Initialized")
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("name", "Elliot", 0).Err()

	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json[12]))
	fmt.Println(len(string(json)))

	//	err = client.Set("id1234", json, 0).Err()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	val, err = client.Get("id1234").Result()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(val)
}
