//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

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

// AddrMode defines 6510 instruction addressing modes.
type AddrMode byte

var addrModes = map[AddrMode]string{
	AddrImp: "",
	AddrIMM: "imm",
	AddrZP:  "zp",
	AddrZPX: "zpx",
	AddrZPY: "zpy",
	AddrIZX: "izx",
	AddrIZY: "izy",
	AddrABS: "abs",
	AddrABX: "abx",
	AddrABY: "aby",
	AddrIND: "ind",
	AddrREL: "rel",
}

func (m AddrMode) String() string {
	name, ok := addrModes[m]
	if ok {
		return name
	}
	return fmt.Sprintf("{AddrMode %d}", m)
}

// 6510 addressing modes.
const (
	AddrImp AddrMode = iota
	AddrIMM          // #$00
	AddrZP           // $00
	AddrZPX          // $00,X
	AddrZPY          // $00,Y
	AddrIZX          // ($00,X)
	AddrIZY          // ($00,Y)
	AddrABS          // $0000
	AddrABX          // $0000,X
	AddrABY          // $0000,Y
	AddrIND          // ($0000)
	AddrREL          // $0000 (PC-relative)
)

// Instr provides 6510 instruction information.
type Instr struct {
	Op           Opcode
	Name         string
	Addr         AddrMode
	Cycles       int
	PageBoundary bool
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
