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

func TestBCD(t *testing.T) {
	cpu := NewCPU()
	cpu.V[0x1] = 123
	cpu.BCDOpcode(0xF133)

	if cpu.Memory[cpu.I] != 1 {
		t.Errorf("Centenas incorretas: esperado 1, obteve %d", cpu.Memory[cpu.I])
	}
	if cpu.Memory[cpu.I+1] != 2 {
		t.Errorf("Dezenas incorretas: esperado 2, obteve %d", cpu.Memory[cpu.I+1])
	}
	if cpu.Memory[cpu.I+2] != 3 {
		t.Errorf("Unidades incorretas: esperado 3, obteve %d", cpu.Memory[cpu.I+2])
	}
}

func TestAssign(t *testing.T) {
	cpu := NewCPU()
	cpu.Assign(0x8120)
	if cpu.V[0x1] != cpu.V[0x2] {
		t.Errorf("Program count is incorrect")
	}
}

func TestBitOROp(t *testing.T) {
	cpu := NewCPU()
	cpu.BitOROp(0x8111)
}
