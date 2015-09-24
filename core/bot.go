package main

import (
	"errors"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

var moveCodes map[string]byte = map[string]byte{
	"forward" : 10,
	"back" : 11,
	"left" : 12,
	"right" : 13,
	"stop" : 14,
}

var lookCodes map[string]byte = map[string]byte{
	"center" : 20,
	"left" : 21,
	"right" : 22,
	"up" : 23,
	"down" : 24,
}

var (
    arduino byte = 4
)

type I2CBus interface {
	ReadBytes(addr byte, num int) (value []byte, err error)
	ReadByte(addr byte) (value byte, err error)
	WriteByte(addr, value byte) error
	Close() error
}

type Bot interface {
    Move(direction string) error
    Look(direction string) error
	LedOn(color string) error
	LedOff(color string) error
	Close() error
}

type CoreBot struct {
    bus I2CBus
	ledPin embd.DigitalPin
	leftDistanceSensor CoreDistanceSensor
	rightDistanceSensor CoreDistanceSensor
	centerDistanceSensor CoreDistanceSensor
}

func (bot CoreBot) Move(direction string) error {
	if code, valid := moveCodes[direction]; valid {
		return bot.bus.WriteByte(arduino, code)
	}
	return errors.New("invalid move direction")
}

func (bot CoreBot) Look(direction string) error {
	if code, valid := lookCodes[direction]; valid {
		return bot.bus.WriteByte(arduino, code)
	}
	return errors.New("invalid look direction")
}

func (bot CoreBot) LedOn(color string) error {
	return bot.ledPin.Write(embd.High)
}

func (bot CoreBot) LedOff(color string) error {
	return bot.ledPin.Write(embd.Low)
}

func (bot CoreBot) Close() error {
	return embd.CloseGPIO()
}

func (bot CoreBot) ArduinoStatus() []byte {
	bytes, err := bot.bus.ReadBytes(arduino, 10)
	if err != nil {
		panic(err)
	}
	return bytes
}

func NewCoreBot() CoreBot {
	b := embd.NewI2CBus(1)

	err := embd.InitGPIO()
	if err != nil {
		panic(err)
	}

	p, _ := embd.NewDigitalPin(17)
	p.SetDirection(embd.Out)

	leftSensor := NewCoreDistanceSensor(25, 14)
	centerSensor := NewCoreDistanceSensor(8, 15)
	rightSensor := NewCoreDistanceSensor(7, 23)

	return CoreBot{bus: b, ledPin: p,
			leftDistanceSensor: leftSensor,
			centerDistanceSensor: centerSensor,
			rightDistanceSensor: rightSensor}
}
