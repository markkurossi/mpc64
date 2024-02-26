//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

// Package mos6510 implements the MOS6510 CPU.
//
// Registers:
// - Program Counter (PC or IP): 16 bits
// - Status Register (.P): 8 bits
// - Accumulator (.A): 8 bits
// - Index Registres (.X and .Y): 8 bits
// - Stack Pointer (SP): 8 bits
package mos6510

import (
	"fmt"
	"strconv"
)

// Opcode defines 6510 opcodes.
//
//go:generate mos6510op -o opcodes.go opcodes.txt
type Opcode byte

func (op Opcode) String() string {
	return Instructions[op].Name
}

// Size returns the size of the opcode and its arguments in bytes.
func (op Opcode) Size() int {
	return 1 + Instructions[op].Addr.Size()
}

// ArgSize returns the size of the opcode's arguments in bytes.
func (op Opcode) ArgSize() int {
	return Instructions[op].Addr.Size()
}

// AddrMode returns the opcode's AddrMode.
func (op Opcode) AddrMode() AddrMode {
	return Instructions[op].Addr
}

// Data describes if the opcode reads or writes memory.
func (op Opcode) Data() bool {
	return Instructions[op].Data
}

// Jump describes if the opcode changes the program execution address.
func (op Opcode) Jump() bool {
	return Instructions[op].Jump
}

// BlockEnd describes if the opcode terminates a basic block i.e. the
// program will not advance to the next instruction.
func (op Opcode) BlockEnd() bool {
	return Instructions[op].BlockEnd
}

// AddrMode defines 6510 instruction addressing modes.
type AddrMode byte

// 6510 addressing modes.
const (
	// Implied or Accumulator addressing, no data followed by the
	// opcode.
	AddrImp AddrMode = iota

	// Immediate addressing, one byte of data: #64, $40
	AddrIMM

	// Absolute addressing, address in low-endian format: $lo,$hi
	AddrABS

	// Absolute addressing indexed with X: $lo,$hi + X
	AddrABX

	// Absolute addressing indexed with X: $lo,$hi + Y
	AddrABY

	// Zeropage addressing, one byte index in zeropage: $lo
	AddrZP

	// Zeropage addressing indexed with X: $lo + X
	AddrZPX

	// Zeropage addressing indexed with Y: $lo + Y
	AddrZPY

	// Relative addressing: signed 8-bit PC-relative offset: #-10
	AddrREL

	// Absolute-indirect addressing: vector to address: ($lo,$hi)
	AddrIND

	// Indexed-indirect addressing indexed with X: ($lo,X)
	AddrIZX

	// Indexed-indirect addressing indexed with Y: ($lo),Y
	AddrIZY
)

var addrModes = map[AddrMode]string{
	AddrImp: "",
	AddrIMM: "imm",
	AddrABS: "abs",
	AddrABX: "abx",
	AddrABY: "aby",
	AddrZP:  "zp",
	AddrZPX: "zpx",
	AddrZPY: "zpy",
	AddrREL: "rel",
	AddrIND: "ind",
	AddrIZX: "izx",
	AddrIZY: "izy",
}

func (m AddrMode) String() string {
	name, ok := addrModes[m]
	if ok {
		return name
	}
	return fmt.Sprintf("{AddrMode %d}", m)
}

var addrModeSizes = map[AddrMode]int{
	AddrImp: 0,
	AddrIMM: 1,
	AddrABS: 2,
	AddrABX: 2,
	AddrABY: 2,
	AddrZP:  1,
	AddrZPX: 1,
	AddrZPY: 1,
	AddrREL: 1,
	AddrIND: 2,
	AddrIZX: 1,
	AddrIZY: 1,
}

// Size returns the number of bytes of data the addressing mode has
// followed by the opcode.
func (m AddrMode) Size() int {
	size, ok := addrModeSizes[m]
	if !ok {
		panic("invalid addressing mode")
	}
	return size
}

// Instr provides 6510 instruction information.
type Instr struct {
	Op           Opcode
	Name         string
	Addr         AddrMode
	Cycles       int
	PageBoundary bool
	Data         bool
	Jump         bool
	BlockEnd     bool
}

func (i Instr) String() string {
	str := i.Name

	addr := i.Addr.String()
	if len(addr) > 0 {
		str += " "
		str += addr
	}
	if i.Cycles > 0 {
		str += " "
		str += strconv.Itoa(i.Cycles)
		if i.PageBoundary {
			str += "*"
		}
	}
	return str
}

// Size returns the instruction size (opcode + argument) in bytes.
func (i Instr) Size() int {
	return 1 + i.Addr.Size()
}

var logicalAndArithmeticInstructions = []string{
	"ORA", "AND", "EOR", "ADC", "SBC", "CMP", "CPX", "CPY",
	"DEC", "DEX", "DEY", "INC", "INX", "INY", "ASL", "ROL",
	"LSR", "ROR",
}

var moveInstructions = []string{
	"LDA", "STA", "LDX", "STX", "LDY", "STY", "TAX", "TXA",
	"TAY", "TYA", "TSX", "TXS", "PLA", "PHA", "PLP", "PHP",
}

var jumpFlagInstructions = []string{
	"BPL", "BMI", "BVC", "BVS", "BCC", "BCS", "BNE", "BEQ",
	"BRK", "RTI", "JSR", "RTS", "JMP", "BIT", "CLC", "SEC",
	"CLD", "SED", "CLI", "SEI", "CLV", "NOP",
}

var dataInstructions = map[string]bool{
	"ORA": true,
	"AND": true,
	"EOR": true,
	"ADC": true,
	"SBC": true,
	"CMP": true,
	"CPX": true,
	"CPY": true,
	"DEC": true,
	"INC": true,
	"ASL": true,
	"ROL": true,
	"LSR": true,
	"ROR": true,
	"LDA": true,
	"STA": true,
	"LDX": true,
	"STX": true,
	"LDY": true,
	"STY": true,
}

var jumpInstructions = map[string]bool{
	"BPL": true,
	"BMI": true,
	"BVC": true,
	"BVS": true,
	"BCC": true,
	"BCS": true,
	"BNE": true,
	"BEQ": true,
	"JSR": true,
	"JMP": true,
}

var blockEndInstuctions = map[string]bool{
	"BRK": true,
	"RTI": true,
	"RTS": true,
	"JMP": true,
}

func init() {
	for idx, instr := range Instructions {
		if dataInstructions[instr.Name] {
			Instructions[idx].Data = true
		}
		if jumpInstructions[instr.Name] {
			Instructions[idx].Jump = true
		}
		if blockEndInstuctions[instr.Name] {
			Instructions[idx].BlockEnd = true
		}
	}
}
