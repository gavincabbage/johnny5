package main

import (
    "fmt"
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

    time.Sleep(10 * time.Microsecond) // necessary to sleep here?

    fmt.Println("Triggering measurement...")

    err := sensor.trigPin.Write(embd.High)
	if err != nil {
		panic(err)
	}

    time.Sleep(10 * time.Microsecond)

    err = sensor.trigPin.Write(embd.Low)
    if err != nil {
        panic(err)
    }

    fmt.Println("Reading result...")

    startTime := time.Now()
    e, _ := sensor.echoPin.Read()
    for e == embd.Low {
        startTime = time.Now()
        e, _ = sensor.echoPin.Read()
    }

    endTime := time.Now()
    e, _ = sensor.echoPin.Read()
    for e == embd.High {
        endTime = time.Now()
        e, _ = sensor.echoPin.Read()
    }

    duration := endTime.Sub(startTime).Seconds()

    fmt.Println("Duration: ", duration)
    distance := duration * 17150
    fmt.Println("Distance: ", distance)

    fmt.Println("Closing GPIO...")
    embd.CloseGPIO()

    fmt.Println("Done.")

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
