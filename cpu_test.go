package main

import (
	"testing"
)

func TestCPU(t *testing.T) {
	cpu := NewCPU()
	if cpu == nil {
		t.Fatal("NewCPU() returned nil")
	}
	if cpu.PC != 0x200 {
		t.Errorf("Expected PC to be 0x200, got 0x%X", cpu.PC)
	}
}

func TestFlow2Opcode(t *testing.T) {
	cpu := NewCPU()
	cpu.Flow2NNN(0x2111)
	programCount := cpu.PC
	if programCount != 0x111 {
		t.Errorf("Program count is incorrect")
	}
}

func TestFlow0Opcode(t *testing.T) {
	cpu := NewCPU()
	cpu.Stack.Push(0x0111)
	cpu.Flow00EEOpcode(0x111)
	if cpu.PC != 0x0111 {
		t.Errorf("Program coun is incorrect")
	}
}
