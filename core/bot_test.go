package main

import (
	"testing"
)

func TestCoreBotInit(t *testing.T) {
	mockBus := &mockI2CBus{}
	bot := &CoreBot{bus: mockBus}
	if b := bot.bus; b == nil {
		t.Fail()
	}
}

func TestCoreBotMoveValid(t *testing.T) {
	mockBus := &mockI2CBus{}
	bot := &CoreBot{bus: mockBus}
	err := bot.Move("forward")
	if err != nil {
		t.Fail()
	}
}

func TestCoreBotMoveInvalid(t *testing.T) {
	mockBus := &mockI2CBus{}
	bot := &CoreBot{bus: mockBus}
	err := bot.Move("invalid")
	if err == nil {
		t.Fail()
	}
}

func TestCoreBotLookValid(t *testing.T) {
	mockBus := &mockI2CBus{}
	bot := &CoreBot{bus: mockBus}
	err := bot.Look("center")
	if err != nil {
		t.Fail()
	}
}

func TestCoreBotLookInvalid(t *testing.T) {
	mockBus := &mockI2CBus{}
	bot := &CoreBot{bus: mockBus}
	err := bot.Look("invalid")
	if err == nil {
		t.Fail()
	}
}
