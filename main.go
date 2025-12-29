package main

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

func main() {
	cpu := NewCPU()
	cpu.LoadROM("roms/Maze (alt) [David Winter, 199x].ch8")
	cpu.Cycle()
}
