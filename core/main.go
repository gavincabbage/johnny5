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
	measurementPrecision int = 2
)

var (
	bot           CoreBot
	redisConn     redis.Conn
	redisSubConn  redis.PubSubConn
	subscriptions = [3]string{"test", "move", "look"}
)

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}

func signalHandler(signalChan chan os.Signal) {
	for _ = range signalChan {
		fmt.Println("Received an interrupt, stopping...")
		bot.Close()
		os.Exit(0)
	}
}

func newRedisConnection(redisAddress string) redis.Conn {
	c, err := redis.Dial("tcp", redisAddress)
	fatal(err)
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

type redisMessage struct {
	channel string
	data string
}

func publishToRedis(pubChan chan redisMessage) {
	for {
		message := <-pubChan
		fmt.Println("publishing message <" + message.data + "> on channel <" + message.channel + ">")
		redisConn.Do("PUBLISH", message.channel, message.data)
	}
}

func publishMeasurements(channel string, fn func() float64, pubChan chan redisMessage) {
	for {
		time.Sleep(30 * time.Second)
		measurement := fn()
		measurementStr := strconv.FormatFloat(measurement, 'f', measurementPrecision, 64)
		pubChan <- redisMessage{channel: channel, data: measurementStr}
	}
}

func mouseOdometer() {
	fmt.Println("Mouse odometer engaged")
	f, err := os.Open("/dev/input/mice")
	fatal(err)

	readBuffer := make([]byte, 3)

	for {
		_, err := f.Read(readBuffer)
		fatal(err)
		//fmt.Println("mouse: x = ", int8(readBuffer[1]), ", y = ", int8(readBuffer[2]))
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

	for _, sub := range subscriptions {
		fmt.Println("subscribing to ", sub)
		redisSubConn.Subscribe(sub)
	}

	fmt.Println("entering main loop")
	subChan := make(chan string)
	pubChan := make(chan redisMessage)
	go listenToSubscriptions(subChan, redisSubConn)
	go publishToRedis(pubChan)
	go publishMeasurements("distance.left", bot.SenseLeftDistance, pubChan)
	go publishMeasurements("distance.right", bot.SenseRightDistance, pubChan)
	go publishMeasurements("distance.center", bot.SenseCenterDistance, pubChan)
	go mouseOdometer()
	for {
		select {
		case data := <-subChan:
			fmt.Println("message:", data)
		}
	}
}
