package main

import (
	"os"
	"os/signal"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	 redisAddress string = "127.0.0.1:6379"
)

var (
	bot CoreBot
	redisConn redis.Conn
	redisSub redis.PubSubConn
)

func main() {

	bot = NewCoreBot()
	defer bot.Close()

	// catch interrupts so we close GPIO on Ctrl-C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			fmt.Println("Received an interrupt, stopping...")
			bot.Close()
			os.Exit(0)
		}
	}()

	redisConn, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		panic(err)
	}
	defer redisConn.Close()

	redisSubConn, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		panic(err)
	}
	redisSub = redis.PubSubConn{Conn: redisSubConn}
	defer redisSub.Close()

	fmt.Println("Arduino status: ", string(bot.ArduinoStatus()))

	redisSub.Subscribe("testchan")

	fmt.Println("Entering main loop")
	for {
		switch msg := redisSub.Receive().(type) {
		case redis.Message:
			fmt.Println("got redis message: ", string(msg.Data))
		case error:
			panic(msg)
		default:
			fmt.Println("default")
		}
	}
}
