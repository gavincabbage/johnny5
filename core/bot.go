package main

import (
	"errors"
	"fmt"
	"time"
	"math"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
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
	aAddr byte = 0x1d
	bAddr byte = 0x6b
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
	RunMag()
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

func (bot CoreBot) determineHeading(xin, yin uint16) float64 {
	x := float64(xin)
	y := float64(yin)
	if y == 0.0 {
		if x < 0.0 {
			return 180.0
		} else {
			return 0.0
		}
	} else if y > 0.0 {
		return 90 - math.Atan(x/y)*180/math.Pi
	} else {
		return  270 - math.Atan(x/y)*180/math.Pi
	}
}

func (bot CoreBot) RunMag() {
	var accelMagAddr byte = 0x1d
	var accelMagCtrlReg5 byte = 0x24
	var accelMagCtrlReg6 byte = 0x25
	var accelMagCtrlReg7 byte = 0x26

	// enable mag
	fatal(bot.bus.WriteByteToReg(accelMagAddr, accelMagCtrlReg7, 0))
	oldVal, err := bot.bus.ReadByteFromReg(accelMagAddr, accelMagCtrlReg6)
	fatal(err)
	oldVal &= 0x9F
	oldVal |= 0x02 // mag gain
	fatal(bot.bus.WriteByteToReg(accelMagAddr, accelMagCtrlReg6, oldVal))
	// enable temp
	oldVal, err = bot.bus.ReadByteFromReg(accelMagAddr, accelMagCtrlReg5)
	fatal(err)
	newVal := oldVal | (1<<7)
	fatal(bot.bus.WriteByteToReg(accelMagAddr, accelMagCtrlReg5, newVal))

	var magRegister byte = 0x08
	//buffer := make([]byte, 6)
	for {
		t, err := bot.bus.ReadWordFromReg(accelMagAddr, 0x05)
		temp := 21.0 + float64(t)/8
		fatal(err)
		x, err := bot.bus.ReadWordFromReg(accelMagAddr, magRegister)
		fatal(err)
		y, err := bot.bus.ReadWordFromReg(accelMagAddr, magRegister+2)
		fatal(err)
		z, err := bot.bus.ReadWordFromReg(accelMagAddr, magRegister+4)
		fatal(err)
		heading := bot.determineHeading(x, y)
		// err = bot.bus.ReadFromReg(accelMagAddr, magRegister, buffer)
		// if err != nil {
		// 	panic(err)
		// }
		// xlo := int16(buffer[0])
		// xhi := int16(buffer[1])
		// ylo := int16(buffer[2])
		// yhi := int16(buffer[3])
		// zlo := int16(buffer[4])
		// zhi := int16(buffer[5])
		// xhi = xhi << 8
		// x := xhi | xlo
		// yhi = yhi << 8
		// y := yhi | ylo
		// zhi = zhi << 8
		// z := zhi | zlo
		fmt.Println("mag values: x = ", x, ", y = ", y, ", z = ", z, "and t = ", temp)
		fmt.Println("heading: ", heading)
		time.Sleep(0 * time.Second)
	}
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
