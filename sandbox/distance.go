package main

import (
    "fmt"
    "time"
    "github.com/kidoman/embd"
    _ "github.com/kidoman/embd/host/rpi"
)

func main() {

    fmt.Println("Setting up distance sensor...")

    err := embd.InitGPIO()
	if err != nil {
		panic(err)
	}

    trigPin, err := embd.NewDigitalPin(27)
	if err != nil {
		panic(err)
	}
	trigPin.SetDirection(embd.Out)
    err = trigPin.Write(embd.Low)
	if err != nil {
		panic(err)
	}

    echoPin, err := embd.NewDigitalPin(22)
	if err != nil {
		panic(err)
	}
	echoPin.SetDirection(embd.In)

    time.Sleep(2 * time.Second)
    fmt.Println("Triggering measurement...")

    err = trigPin.Write(embd.High)
	if err != nil {
		panic(err)
	}

    time.Sleep(10 * time.Microsecond)

    err = trigPin.Write(embd.Low)
    if err != nil {
        panic(err)
    }

    fmt.Println("Reading result...")

    startTime := time.Now()
    e, _ := echoPin.Read()
    for e == embd.Low {
        startTime = time.Now()
        e, _ = echoPin.Read()
    }

    endTime := time.Now()
    e, _ = echoPin.Read()
    for e == embd.High {
        endTime = time.Now()
        e, _ = echoPin.Read()
    }

    duration := endTime.Sub(startTime).Seconds()

    fmt.Println("Result: ", duration)

    fmt.Println("Closing GPIO...")
    embd.CloseGPIO()

    fmt.Println("Done.")

}
