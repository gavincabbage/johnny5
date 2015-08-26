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
    arduino1, arduino2 byte = 4, 5
)

type Bot interface {
    Move(direction string) error
    Look(direction string) error
}

type I2CBus interface {
	ReadByte(addr byte) (value byte, err error)
	WriteByte(addr, value byte) error
	Close() error
}

type CoreBot struct {
    bus I2CBus
}

func (bot CoreBot) Move(direction string) error {
	if code, valid := moveCodes[direction]; valid {
		return bot.bus.WriteByte(arduino1, code)
	}
	return errors.New("invalid move direction")
}

func (bot CoreBot) Look(direction string) error {
	if code, valid := lookCodes[direction]; valid {
		return bot.bus.WriteByte(arduino2, code)
	}
	return errors.New("invalid look direction")
}

func (bot CoreBot) Close() error {
	return nil
}

func NewCoreBot() CoreBot {
	b := embd.NewI2CBus(1)
	return CoreBot{bus: b}
}
