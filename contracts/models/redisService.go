package models

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/glide/path"
	"github.com/garyburd/redigo/redis"
	"github.com/robfig/config"
)

var (
	host,
	port string
)

const Redis_expire_time_EX = "EX"
const Redis_expire_time_PX = "PX"

// Initialize the redis configuration
func init() {
	goPath := path.Gopath()
	fmt.Println(goPath)
	var buffer bytes.Buffer
	buffer.WriteString(goPath)
	buffer.WriteString("/src/Browser-achain/conf/databaseConfig.ini")
	c, _ := config.ReadDefault(buffer.String())
	host, _ = c.String("redis", "host")
	port, _ = c.String("redis", "port")
	fmt.Println("\n The current redis IP and port are:", host, port)
}

// set key
func Set(key, value string) error {
	c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}
	defer c.Close()

	_, err = c.Do("SET", key, value)

	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}

	return nil
}

// EX :seconds
// PX :milliseconds
func SetWithExpire(key, value, expireType, expireTime string) error {
	c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}
	defer c.Close()

	_, err = c.Do("SET", key, value, expireType, expireTime)

	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}

	return nil
}

// get key
func Get(key string) (string, error) {
	c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return "", err
	}
	defer c.Close()

	value, err := redis.String(c.Do("GET", key))

	if err != nil {
		fmt.Println("redis get failed:", err)
		return "", nil
	}
	return value, nil
}

// delete key
func Delete(key string) error {
	c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}
	defer c.Close()

	_, err = c.Do("DEL", key)
	if err != nil {
		fmt.Println("redis delete failed:", err)
	}
	return err
}
