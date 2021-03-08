package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"encoding/json"

	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var names = []string{
	"Hugo",
	"Paco",
	"Luis",
	"Carlos",
	"Alejandro",
	"Elliot",
}

func main() {

	start := time.Now()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//fmt.Println(names)

	for i := 0; i < 50000; i++ {

		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(names))
		nm := names[index]

		rand.Seed(time.Now().UnixNano())
		ag := rand.Intn(50)

		field := Author{
			Name: nm,
			Age:  ag,
		}

		json, err := json.Marshal(field)
		if err != nil {
			fmt.Println(err)
		}

		id := "ID"

		for {
			id += "0"
			if len(id)+len(string(strconv.Itoa(i))) == 10 {
				id += strconv.Itoa(i)
				break
			}
		}

		//fmt.Println(id)

		err = client.Set(id, json, 0).Err()
		if err != nil {
			fmt.Println(err)
		}

	}

	fmt.Printf("\nTiempo ejecución: %v ms\n", time.Since(start).Milliseconds())

	//val, err := client.Get("id1234").Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(val)

}
