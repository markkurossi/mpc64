//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package mos6510

import (
	"fmt"
	"os"
	"testing"

	"github.com/markkurossi/tabulate"
)

func TestOpcodeString(t *testing.T) {
	for i := 0; i < 256; i++ {
		fmt.Printf("%02X:\t%v\n", i, Opcode(i).String())
	}
}

func TestInstrString(t *testing.T) {
	for idx, i := range Instructions {
		fmt.Printf("%02X:\t%v\n", idx, i)
	}
}

func TestLogicalAndArithmetic(t *testing.T) {
	testCommandSet(t, logicalAndArithmeticInstructions)
}

func TestMove(t *testing.T) {
	testCommandSet(t, moveInstructions)
}

func TestJumpFlag(t *testing.T) {
	testCommandSet(t, jumpFlagInstructions)
}

func testCommandSet(t *testing.T, names []string) {
	tab := tabulate.New(tabulate.UnicodeLight)
	tab.Header("Op")
	tab.Header("imp")
	tab.Header("imm")
	tab.Header("zp")
	tab.Header("zpx")
	tab.Header("zpy")
	tab.Header("izx")
	tab.Header("izy")
	tab.Header("abs")
	tab.Header("abx")
	tab.Header("aby")
	tab.Header("ind")
	tab.Header("rel")

	for _, name := range names {
		instrs := make(map[AddrMode]Instr)

		for _, instr := range Instructions {
			if instr.Name == name {
				instrs[instr.Addr] = instr
			}
		}

		row := tab.Row()
		row.Column(name)
		for a := AddrImp; a <= AddrREL; a++ {
			instr, ok := instrs[a]
			if ok {
				row.Column(fmt.Sprintf("$%02X", int(instr.Op)))
			} else {
				row.Column("")
			}
		}
	}
	tab.Print(os.Stdout)
}
