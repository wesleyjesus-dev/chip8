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

	cpu.V[2] = 42
	cpu.V[1] = 0
	cpu.PC = 0x200

	cpu.Assign(0x8120)

	if cpu.V[1] != 42 {
		t.Errorf("V1 should be 42, but is %d", cpu.V[1])
	}

	if cpu.V[2] != 42 {
		t.Errorf("V2 don't should be changed")
	}

	if cpu.PC != 0x202 {
		t.Errorf("PC should be next to 0x202, but is 0x%X", cpu.PC)
	}
}

func TestBitOROp(t *testing.T) {
	cpu := NewCPU()
	cpu.V[1] = 100
	cpu.V[2] = 200

	cpu.BitOROp(0x8121)

	if cpu.V[1] != 236 {
		t.Errorf("The Bitwise OR is incorrect")
	}

	if cpu.PC != 0x202 {
		t.Errorf("PC not incremented")
	}
}

func TestBitANDOp(t *testing.T) {
	cpu := NewCPU()
	cpu.V[1] = 100
	cpu.V[2] = 200

	cpu.BitANDOp(0x8122)

	if cpu.V[1] != 64 {
		t.Errorf("The Bitwise AND is incorrect")
	}

	if cpu.PC != 0x202 {
		t.Errorf("PC not incremented")
	}
}

func TestBitXOROp(t *testing.T) {
	cpu := NewCPU()
	cpu.V[1] = 100
	cpu.V[2] = 200

	cpu.BitXOROp(0x8123)

	if cpu.V[1] != 172 {
		t.Errorf("The bitwise XOR is incorrect")
	}

	if cpu.PC != 0x202 {
		t.Errorf("PC not incoremented")
	}
}

func TestBitShiftRightOp(t *testing.T) {
	cpu := NewCPU()
	cpu.V[1] = 100

	cpu.BitShiftRightOp(0x8126)

	if cpu.V[1] != 50 {
		t.Errorf("The bitwise shift right is incorrect")
	}
	if cpu.V[0xF] != 0 {
		t.Errorf("The bitwise shift right is incorrect")
	}
	if cpu.PC != 0x202 {
		t.Errorf("PC not incremented")
	}
}
