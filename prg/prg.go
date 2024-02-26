//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

// Package prg implements PRG file format.
package prg

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/markkurossi/mpc64/mos6510"
	"github.com/markkurossi/mpc64/petscii"
)

var (
	bo = binary.LittleEndian
)

// SegType specifies program segment type.
type SegType byte

// Program segment types.
const (
	SegNone SegType = iota
	SegBasic
	SegCode
	SegData
	SegAddr
)

// Prg defines a program.
type Prg struct {
	Load     uint16
	Start    uint16
	Data     []byte
	SegTypes []SegType
}

// MemToData maps an absolute memory addess into the Data array.
func (prg *Prg) MemToData(addr uint16) (offset int, err error) {
	if addr < prg.Load {
		return 0, fmt.Errorf("address %v underflow %v", addr, prg.Load)
	}
	if int(addr-prg.Load) >= len(prg.Data) {
		return 0, fmt.Errorf("address %v overflow %v",
			addr, int(prg.Load)+len(prg.Data))
	}
	return int(addr - prg.Load), nil
}

// DataToMem maps a Data offset to an absolute memory address.
func (prg *Prg) DataToMem(offset int) uint16 {
	return uint16(offset) + prg.Load
}

// Print prints the program to standard output.
func (prg *Prg) Print() error {
	pc, err := prg.MemToData(prg.Start)
	if err != nil {
		return err
	}
	for pc < len(prg.Data) {
		switch prg.SegTypes[pc] {
		case SegCode:
			op := mos6510.Opcode(prg.Data[pc])

			var i8 int8
			var addr uint16

			switch op.ArgSize() {
			case 0:
			case 1:
				i8 = int8(prg.Data[pc+1])
			case 2:
				addr = bo.Uint16(prg.Data[pc+1:])
			}
			fmt.Printf("%04X: %v", prg.DataToMem(pc), op)

			switch op.AddrMode() {
			case mos6510.AddrImp:
			case mos6510.AddrIMM:
				fmt.Printf(" #$%02X", uint8(i8))
			case mos6510.AddrABS:
				fmt.Printf(" $%04X", addr)
			case mos6510.AddrABX:
				fmt.Printf(" $%04X,X", addr)
			case mos6510.AddrABY:
				fmt.Printf(" $%04X,Y", addr)
			case mos6510.AddrZP:
				fmt.Printf(" $%02X", i8)
			case mos6510.AddrZPX:
				fmt.Printf(" $%02X,X", i8)
			case mos6510.AddrZPY:
				fmt.Printf(" $%02X,Y", i8)
			case mos6510.AddrREL:
				fmt.Printf(" #%d", i8)
			case mos6510.AddrIND:
				fmt.Printf(" ($%04X)", addr)
			case mos6510.AddrIZX:
				fmt.Printf(" ($%02X,X)", addr)
			case mos6510.AddrIZY:
				fmt.Printf(" ($%02X),Y", addr)
			}
			fmt.Println()

			pc += op.Size()

		case SegData:
			start := pc
			for pc < len(prg.Data) {
				if prg.SegTypes[pc] != SegData {
					break
				}
				pc++
			}
			prg.printData(start, pc, ".byte")

		case SegNone:
			start := pc
			for pc < len(prg.Data) {
				if prg.SegTypes[pc] != SegNone {
					break
				}
				pc++
			}
			prg.printData(start, pc, ".none")

		default:
			return fmt.Errorf("%04X: type %v not supported",
				prg.DataToMem(pc), prg.SegTypes[pc])
		}
	}
	return nil
}

func (prg *Prg) printData(from, to int, label string) {
	for from < to {
		fmt.Printf("%04X: %s", prg.DataToMem(from), label)
		var str string
		for i := 0; i < 8; i++ {
			if from < to {
				fmt.Printf(" $%02X", prg.Data[from])
				r := petscii.Shifted[prg.Data[from]]
				if r == 0 {
					r = '.'
				}
				str += string(r)
			} else {
				fmt.Printf("    ")
				str += " "
			}
			from++
		}
		fmt.Printf("\t%s\n", str)
	}
}

// Load loads program from the named file.
func Load(file string) (*Prg, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return Parse(data)
}

// Parse parses the program data.
func Parse(data []byte) (*Prg, error) {
	if len(data) < 7 {
		return nil, fmt.Errorf("data too short, need at least 7 bytes")
	}
	load := bo.Uint16(data)
	data = data[2:]

	prg := &Prg{
		Load:     load,
		Data:     data,
		SegTypes: make([]SegType, len(data)),
	}
	prg.SegTypes[0] = SegAddr
	prg.SegTypes[1] = SegAddr

	next, err := prg.MemToData(bo.Uint16(data))
	if err != nil {
		return nil, err
	}
	if next < 2 || next+2 >= len(data) {
		return nil, fmt.Errorf("next line out of bounds")
	}
	for i := 2; i < next; i++ {
		prg.SegTypes[i] = SegBasic
	}
	zero := bo.Uint16(data[next:])
	if zero != 0 {
		return nil, fmt.Errorf("next line not $0000")
	}
	prg.SegTypes[next] = SegAddr
	prg.SegTypes[next+1] = SegAddr

	prg.Start = prg.DataToMem(next + 2)

	err = prg.parseCodeFromAddr(prg.Start)
	if err != nil {
		return nil, err
	}

	// Parse all unmarked blocks preceded by code.
	for pc := 0; pc < len(prg.Data); pc++ {
		if prg.SegTypes[pc] == 0 && prg.SegTypes[pc-1] == SegCode {
			err = prg.parseCodeFromAddr(prg.DataToMem(pc))
			if err != nil {
				return nil, err
			}
		}
	}

	// Mark all unmarked blocks as data.
	for pc := 0; pc < len(prg.Data); pc++ {
		if prg.SegTypes[pc] == 0 {
			prg.SegTypes[pc] = SegData
		}
	}

	return prg, nil
}

func (prg *Prg) parseCodeFromAddr(start uint16) (err error) {
	var pending []uint16
	pending = append(pending, start)

	for len(pending) > 0 {
		pending, err = prg.parseCode(pending[0], pending[1:])
		if err != nil {
			return err
		}
	}
	return nil
}

func (prg *Prg) parseCode(addr uint16, pending []uint16) ([]uint16, error) {
	pc, err := prg.MemToData(addr)
	if err != nil {
		return nil, err
	}
	for pc < len(prg.Data) {
		if prg.SegTypes[pc] != 0 {
			break
		}
		op := mos6510.Opcode(prg.Data[pc])
		if pc+op.Size() > len(prg.Data) {
			return nil, fmt.Errorf("%04X: %v: truncated code",
				prg.DataToMem(pc), op)
		}
		for i := 0; i < op.Size(); i++ {
			prg.SegTypes[pc+i] = SegCode
		}
		var i8 int8
		var addr uint16

		switch op.ArgSize() {
		case 0:

		case 1:
			i8 = int8(prg.Data[pc+1])
			switch op.AddrMode() {
			case mos6510.AddrIMM:

			case mos6510.AddrZP, mos6510.AddrZPX, mos6510.AddrZPY:

			case mos6510.AddrREL:
				next := pc + op.Size() + int(i8)
				if next >= 0 && next < len(prg.Data) {
					addr := prg.DataToMem(next)
					if op.Jump() {
						pending = append(pending, addr)
					}
				}

			case mos6510.AddrIZX, mos6510.AddrIZY:

			default:
				return nil, fmt.Errorf("%04X: %v: %v not supported",
					prg.DataToMem(pc), op, op.AddrMode())
			}

		case 2:
			addr = bo.Uint16(prg.Data[pc+1:])
			ofs, err := prg.MemToData(addr)
			if err == nil {
				switch op.AddrMode() {
				case mos6510.AddrABS, mos6510.AddrABX, mos6510.AddrABY:
					if op.Jump() {
						pending = append(pending, addr)
					}
					if op.Data() {
						if prg.SegTypes[ofs] == 0 {
							prg.SegTypes[ofs] = SegData
						}
					}
				default:
					return nil, fmt.Errorf("%04X: %v: %v not supported",
						prg.DataToMem(pc), op, op.AddrMode())
				}

			}
		}
		pc += op.Size()
		if op.BlockEnd() {
			break
		}
	}

	return pending, nil
}
