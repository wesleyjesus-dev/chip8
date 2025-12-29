package main

import (
	"fmt"
	"os"
)

type CPU struct {
	V  [16]uint8
	I  uint16
	PC uint16

	Memory [4096]uint8

	Stack Stack

	Display [64 * 32]uint8

	DisplayTimer uint8
	SoundTimer   uint8

	Keypad [16]uint8
}

func NewCPU() *CPU {
	cpu := &CPU{
		PC: 0x200, // Program counter starts at 0x200
	}
	return cpu
}

func (cpu *CPU) LoadROM(romPath string) error {
	romData, err := os.ReadFile(romPath)
	if err != nil {
		return fmt.Errorf("failed to read ROM file: %v", err)
	}
	copy(cpu.Memory[0x200:], romData)
	return nil
}

func (cpu *CPU) FetchOpcode() uint16 {
	opcode := uint16(cpu.Memory[cpu.PC])<<8 | uint16(cpu.Memory[cpu.PC+1])
	fmt.Printf("Fetched opcode: 0x%04X at PC: 0x%X\n", opcode, cpu.PC)
	return opcode
}

func (cpu *CPU) Cycle() {
	for ; cpu.PC < cpu.PC+uint16(len(cpu.Memory)); cpu.PC += 2 {
		opcode := cpu.FetchOpcode()
		// Apply masking to decode opcode
		switch opcode & 0xF000 {
		case 0x0000:
			switch opcode & 0x00FF {
			case 0x0000:
			case 0x00E0:
				//display clear
			case 0x00EE:
				cpu.Flow00EEOpcode(opcode)
			}

			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0x1000: // 1NNN: Jump to address NNN
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0x2000:
			cpu.Flow2NNN(opcode)
		case 0x3000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0x4000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0x5000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0x6000: // 6XNN: Set VX to NN
			x := (opcode & 0x0F00) >> 8
			nn := byte(opcode & 0x00FF)
			//V[x] = nn
			fmt.Printf("Set V[%X] to %02X\n", x, nn)
			cpu.V[x] = nn
		case 0x7000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0x8000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0x9000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0xA000: // ANNN: Set I to NNN
			nnn := opcode & 0x0FFF
			//I = nnn
			fmt.Printf("Set I to %03X\n", nnn)
			cpu.I = nnn
		case 0xB000:
			cpu.FlowBNNN(opcode)
		case 0xC000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0xD000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0xE000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0xF000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		default:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		}
	}
}

// write the value of vX as BCD value at the addresses I, I+1 and I+2
func (cpu *CPU) BCDOpcode(opcode uint16) {
	x := opcode & 0x0F00 >> 8
	v := cpu.V[x]

	cpu.Memory[cpu.I] = v / 100         // here take the rest of division [C]
	cpu.Memory[cpu.I+1] = (v / 10) % 10 // removing the last digit [D]
	cpu.Memory[cpu.I+2] = v % 10        // take the unit [U]
}

func (cpu *CPU) Assign(opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	cpu.V[x] = cpu.V[y]
	cpu.PC += 2
}

func (cpu *CPU) Flow2NNN(opcode uint16) {
	nnn := opcode & 0x0FFF
	cpu.Stack.Push(cpu.PC + 2)
	cpu.PC = nnn
}

func (cpu *CPU) Flow00EEOpcode(opcode uint16) {
	address := cpu.Stack.Pop()
	cpu.PC = address
}

func (cpu *CPU) Flow1NNN(opcode uint16) {
	nnn := opcode & 0x0FFF
	cpu.PC = nnn
}

func (cpu *CPU) FlowBNNN(opcode uint16) {
	v0 := cpu.V[0]
	nnn := opcode & 0x0FFF
	cpu.PC = uint16(v0) + nnn
}
