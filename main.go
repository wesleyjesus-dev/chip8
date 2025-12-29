package main

import (
	"fmt"
	"io/ioutil"
)

/*
	1. Read ROM file
	2. Initialize memory and registers
	3. Load ROM into memory
	4. Emulation loop:
		- Fetch opcode
		- Decode opcode
		- Execute opcode
		- Update timers
	5. Handle input
	6. Render graphics
*/

type Registers struct {
	V [16]byte // General purpose registers V0 to VF

	I uint16 // Index register
}

func main() {
	// Step 1: Read ROM file
	romPath := "roms/Maze (alt) [David Winter, 199x].ch8"
	romData, err := ioutil.ReadFile(romPath)
	if err != nil {
		fmt.Printf("Failed to read ROM file: %v\n", err)
		return
	}
	fmt.Printf("ROM file '%s' loaded, size: %d bytes\n", romPath, len(romData))

	// Step 2: Initialize memory and registers
	var registers Registers // General purpose registers V0 to VF
	var memory [4096]byte
	//var V [16]byte        // General purpose registers V0 to VF
	//var I uint16          // Index register
	var pc uint16 = 0x200 // Program counter starts at 0x200
	fmt.Println("Memory and registers initialized.")

	// Step 3: Load ROM into memory starting at address 0x200
	copy(memory[pc:], romData)
	fmt.Printf("ROM loaded into memory at address 0x%X\n", pc)

	// Step 4: Emulation loop (simplified for demonstration)
	// Aqui ele esta:
	/*
		1. inicializando uma variavel int de 16bits para o pc com valor 0x200
		2. Utilizando o shift left para pegar o valor do byte na posiçao pc e deslocar 8 bits para esquerda
		3. Utilizando o operador OR para combinar o valor deslocado com o valor do byte na posiçao pc + 1
		4. Imprimindo o opcode buscado e o valor do pc
	*/
	fmt.Printf("Started flow!")
	for pc := pc; pc < pc+uint16(len(romData)); pc += 2 {
		// Fetch opcode (2 bytes)
		opcode := uint16(memory[pc])<<8 | uint16(memory[pc+1])
		fmt.Printf("Fetched opcode: 0x%04X at PC: 0x%X\n", opcode, pc)

		// Decode and execute opcode (not implemented in this snippet)
		switch opcode & 0xF000 {
		case 0x0000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0x1000: // 1NNN: Jump to address NNN
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
		case 0x2000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
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
			registers.V[x] = nn
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
			registers.I = nnn
		case 0xB000:
			fmt.Printf("Opcode 0x%04X not implemented.\n", opcode)
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
