package main

import (
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

type DistanceSensor interface {
	Sense() (float64, error)
}

type CoreDistanceSensor struct {
	trigPin embd.DigitalPin
	echoPin embd.DigitalPin
}

func (sensor CoreDistanceSensor) Sense() (float64, error) {

	err := sensor.trigPin.Write(embd.High)
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Microsecond)

	err = sensor.trigPin.Write(embd.Low)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()
	holdTime := startTime
	holdDuration := startTime.Sub(holdTime).Seconds()
	e, _ := sensor.echoPin.Read()
	for e == embd.Low {
		startTime = time.Now()
		holdDuration = startTime.Sub(holdTime).Seconds()
		if holdDuration > 1.0 {
			break
		}
		e, _ = sensor.echoPin.Read()
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime).Seconds()
	e, _ = sensor.echoPin.Read()
	for e == embd.High {
		endTime = time.Now()
		duration = endTime.Sub(startTime).Seconds()
		if duration > 1.0 {
			break
		}
		e, _ = sensor.echoPin.Read()
	}

	duration = endTime.Sub(startTime).Seconds()
	distance := duration * 17150

	return distance, nil
}

func NewCoreDistanceSensor(t int, e int) CoreDistanceSensor {

	trigPin, err := embd.NewDigitalPin(t)
	if err != nil {
		panic(err)
	}
	trigPin.SetDirection(embd.Out)
	err = trigPin.Write(embd.Low)
	if err != nil {
		panic(err)
	}

	echoPin, err := embd.NewDigitalPin(e)
	if err != nil {
		panic(err)
	}
	echoPin.SetDirection(embd.In)

	return CoreDistanceSensor{trigPin: trigPin, echoPin: echoPin}
}
