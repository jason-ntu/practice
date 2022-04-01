package cache

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func New() {
	var err error
	Client = redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "",
		DB:       0,
	})
	err = Client.Set("Get", "0", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	err = Client.Set("GetAll", "0", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	err = Client.Set("Create", "0", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	err = Client.Set("Update", "0", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	err = Client.Set("Delete", "0", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	return
}
