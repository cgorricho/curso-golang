package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"encoding/json"

	"github.com/go-redis/redis"
)

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
<<<<<<< HEAD
	TimeStamp int64  `json:"time"`
=======
	TimeStamp int    `json:"time"`
>>>>>>> bb15ae39f5d6019cac2ca3ffd0dffcba35427253
}

var names = []string{
	"Hugo",
	"Paco",
	"Luis",
	"Carlos",
	"Alejandro",
	"Elliot",
}

var wg sync.WaitGroup

func main() {

	start := time.Now()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//fmt.Println(names)

	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func(ind int) {
			rand.Seed(time.Now().UnixNano())
			index := rand.Intn(len(names))
			nm := names[index]

			rand.Seed(time.Now().UnixNano())
			ag := rand.Intn(50)

<<<<<<< HEAD
			ts := time.Now().UnixNano()
=======
			ts := int(time.Now().UnixNano())
>>>>>>> bb15ae39f5d6019cac2ca3ffd0dffcba35427253

			field := Author{
				Name:      nm,
				Age:       ag,
				TimeStamp: ts,
			}

			json, err := json.Marshal(field)
			if err != nil {
				fmt.Println(err)
			}

			var id strings.Builder
			id.WriteString("ID-")

			for {
				id.WriteString("0")
				if id.Len()+len(string(strconv.Itoa(ind))) == 10 {
					id.WriteString(strconv.Itoa(ind))
					break
				}
			}

			//fmt.Println(id)

			var tex time.Duration
			tex = 0 * 15 * 1e9 // dias * hrs * min * segundos * nanosegundos (1e9)

			err = client.Set(id.String(), json, tex).Err()
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()

	fmt.Printf("\nTiempo ejecución: %v ms\n", time.Since(start).Milliseconds())

	//val, err := client.Get("id1234").Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(val)

}
