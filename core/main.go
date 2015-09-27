package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	redisAddress string = "127.0.0.1:6379"
)

var (
	bot           CoreBot
	redisConn     redis.Conn
	redisSubConn  redis.PubSubConn
	subscriptions = [3]string{"test", "move", "look"}
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
			subChan <- processMessage(msg)
		case error:
			panic(msg)
		}
	}
}

func processMessage(msg redis.Message) string {
	switch msg.Channel {
	case "test":
		return string(msg.Data)
	case "move":
		bot.Move(string(msg.Data))
		return "move " + string(msg.Data)
	case "look":
		bot.Look(string(msg.Data))
		return "look " + string(msg.Data)
	default:
		return "default"
	}
}

func publishDistanceMeasurements() {
	for {
		time.Sleep(1 * time.Second)
		leftDistance, centerDistance, _ := bot.SenseDistance()
		if centerDistance < 3.0 {
			bot.Stop()
		}
		leftDistanceStr := strconv.FormatFloat(leftDistance, 'f', 2, 64)
		fmt.Println("leftDistanceString =", leftDistanceStr)
		redisConn.Do("PUBLISH", "distance.left", leftDistanceStr)
	}
}

func main() {

	bot = NewCoreBot()
	defer bot.Close()

	redisConn = newRedisConnection(redisAddress)
	defer redisConn.Close()
	redisSubConn = redis.PubSubConn{Conn: newRedisConnection(redisAddress)}
	defer redisSubConn.Close()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go signalHandler(signalChan)

	fmt.Println("arduino status: ", string(bot.ArduinoStatus()))

	leftDistance, centerDistance, rightDistance := bot.SenseDistance()
	fmt.Println("left distance: ", leftDistance)
	fmt.Println("center distance: ", centerDistance)
	fmt.Println("right distance: ", rightDistance)

	for _, sub := range subscriptions {
		fmt.Println("subscribing to ", sub)
		redisSubConn.Subscribe(sub)
	}

	fmt.Println("entering main loop")
	subChan := make(chan string)
	go listenToSubscriptions(subChan, redisSubConn)
	go publishDistanceMeasurements()
	for {
		select {
		case data := <-subChan:
			fmt.Println("message:", data)
		}
	}
}
