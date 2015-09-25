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
)

func signalHandler(signalChan chan os.Signal) {
	for _ = range signalChan {
		fmt.Println("Received an interrupt, stopping...")
		bot.Close()
		os.Exit(0)
	}
}

func newRedisConnection(redisAddress string) redis.Conn {
	c, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		panic(err)
	}
	return c
}

func listenToSubscriptions(subChan chan string, redisSubConn redis.PubSubConn) {
	for {
		switch msg := redisSubConn.Receive().(type) {
		case redis.Message:
			subChan <- string(msg.Data)
		case error:
			panic(msg)
		}
	}
}

func main() {

	bot = NewCoreBot()
	defer bot.Close()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go signalHandler(signalChan)

	redisConn := newRedisConnection(redisAddress)
	defer redisConn.Close()

	redisConn2 := newRedisConnection(redisAddress)
	redisSubConn := redis.PubSubConn{Conn: redisConn2}
	defer redisSubConn.Close()

	fmt.Println("Arduino status: ", string(bot.ArduinoStatus()))

	redisSubConn.Subscribe("testchan")

	fmt.Println("Entering main loop")

	subChan := make(chan string)
	go listenToSubscriptions(subChan, redisSubConn)
	for {
		select {
		case data := <-subChan:
			fmt.Println("received message on subchan: ", data)
		}
	}
}
