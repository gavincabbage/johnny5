package main

import (
	"errors"

	"github.com/gavincabbage/embd"
	_ "github.com/gavincabbage/embd/host/rpi"
)

var moveCodes map[string]byte = map[string]byte{
	"forward": 10,
	"back":    11,
	"left":    12,
	"right":   13,
	"stop":    14,
}

var lookCodes map[string]byte = map[string]byte{
	"center": 20,
	"left":   21,
	"right":  22,
	"up":     23,
	"down":   24,
}

var (
	arduinoAddr byte = 0x04
	aAddr       byte = 0x1d
	bAddr       byte = 0x6b
)

type I2CBus interface {
	ReadByte(addr byte) (byte, error)
	ReadBytes(addr byte, num int) ([]byte, error)
	WriteByte(addr, value byte) error
	WriteByteToReg(addr, reg, value byte) error
	ReadByteFromReg(addr, reg byte) (byte, error)
	ReadWordFromReg(addr, reg byte) (uint16, error)
	Close() error
}

type Bot interface {
	Move(direction string) error
	Stop() error
	Look(direction string) error
	LedOn(color string) error
	LedOff(color string) error
	ArduinoStatus() []byte
	SenseLeftDistance() float64
	SenseRightDistance() float64
	SenseCenterDistance() float64
	Close() error
}

type CoreBot struct {
	bus                  I2CBus
	ledPin               embd.DigitalPin
	leftDistanceSensor   CoreDistanceSensor
	rightDistanceSensor  CoreDistanceSensor
	centerDistanceSensor CoreDistanceSensor
}

func (bot CoreBot) Move(direction string) error {
	if code, valid := moveCodes[direction]; valid {
		return bot.bus.WriteByte(arduinoAddr, code)
	}
	return errors.New("invalid move direction")
}

func (bot CoreBot) Stop() error {
	return bot.Move("stop")
}

func (bot CoreBot) Look(direction string) error {
	if code, valid := lookCodes[direction]; valid {
		return bot.bus.WriteByte(arduinoAddr, code)
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
	bytes, err := bot.bus.ReadBytes(arduinoAddr, 10)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (bot CoreBot) senseDistance(sensor DistanceSensor) float64 {
	distance, err := sensor.Sense()
	if err != nil {
		panic(err)
	}

	return distance
}

func (bot CoreBot) SenseLeftDistance() float64 {
	return bot.senseDistance(bot.leftDistanceSensor)
}

func (bot CoreBot) SenseRightDistance() float64 {
	return bot.senseDistance(bot.rightDistanceSensor)
}

func (bot CoreBot) SenseCenterDistance() float64 {
	return bot.senseDistance(bot.centerDistanceSensor)
}

func NewCoreBot() CoreBot {
	b := embd.NewI2CBus(1)

	err := embd.InitGPIO()
	fatal(err)

	p, _ := embd.NewDigitalPin(17)
	p.SetDirection(embd.Out)

	leftSensor := NewCoreDistanceSensor(25, 14)
	centerSensor := NewCoreDistanceSensor(8, 15)
	rightSensor := NewCoreDistanceSensor(7, 23)

	return CoreBot{bus: b, ledPin: p,
		leftDistanceSensor:   leftSensor,
		centerDistanceSensor: centerSensor,
		rightDistanceSensor:  rightSensor}
}
