//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type opcode struct {
	op           int
	name         string
	addr         string
	cycles       int
	pageBoundary bool
}

func (op opcode) String() string {
	str := op.name
	if len(op.addr) > 0 {
		str += " " + op.addr
	}
	if op.cycles > 0 {
		str += fmt.Sprintf(" %d", op.cycles)
		if op.pageBoundary {
			str += "*"
		}
	}
	return str
}

var (
	opcodes []opcode
	names   = make(map[string]int)
)

func main() {
	output := flag.String("o", "", "output file name")
	flag.Parse()

	if len(*output) == 0 {
		log.Fatal("no output file specified")
	}
	out, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	for _, f := range flag.Args() {
		err := processFile(f)
		if err != nil {
			log.Fatal(err)
		}
	}

	var ops []string
	var maxLen int

	fmt.Fprint(out, `//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package mos6510

// Opcode defines 6510 opcodes.
type Opcode byte

// 6510 opcodes.
const (
`)

	for _, op := range opcodes {
		key := op.name + op.addr
		name := fmt.Sprintf("Op%s", key)
		if names[key] > 1 {
			name += fmt.Sprintf("0x%02X", op.op)
		}
		ops = append(ops, name)
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}
	for idx, op := range ops {
		fmt.Fprintf(out, "\t%s", op)
		if idx == 0 {
			for i := len(op); i < maxLen; i++ {
				fmt.Fprint(out, " ")
			}
			fmt.Fprintf(out, " Opcode = iota // %s\n", opcodes[idx])
		} else {
			for i := len(op); i < maxLen+15; i++ {
				fmt.Fprint(out, " ")
			}
			fmt.Fprintf(out, "// %s\n", opcodes[idx])
		}
	}

	fmt.Fprintln(out, ")")

	fmt.Fprint(out, `
// AddrMode defines 6510 instruction addressing modes.
type AddrMode byte

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

// Instructions define the 6510 instructions.
var Instructions = []Instr{
`)

	for idx, op := range opcodes {
		var addr string
		if len(op.addr) > 0 {
			addr = "Addr" + strings.ToUpper(op.addr)
		} else {
			addr = "AddrImp"
		}

		fmt.Fprintf(out, "\t{\n")
		fmt.Fprintf(out, "\t\tOp:           %s,\n", ops[idx])
		fmt.Fprintf(out, "\t\tName:         %q,\n", op.name)
		fmt.Fprintf(out, "\t\tAddr:         %s,\n", addr)
		fmt.Fprintf(out, "\t\tCycles:       %d,\n", op.cycles)
		fmt.Fprintf(out, "\t\tPageBoundary: %v,\n", op.pageBoundary)
		fmt.Fprintf(out, "\t},\n")
	}
	fmt.Fprintln(out, "}")
}

func processFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for lineno := 0; ; lineno++ {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")

		op := parts[0]
		var addr string
		var cycles int
		var pageBoundary bool

		switch len(parts) {
		case 1:
		case 2:
			cycles, pageBoundary, err = parseCycles(parts[1])
			if err != nil {
				return fmt.Errorf("%s:%d: %s", file, lineno, err.Error())
			}
		case 3:
			addr = parts[1]
			cycles, pageBoundary, err = parseCycles(parts[2])
			if err != nil {
				return err
			}
		}

		opcodes = append(opcodes, opcode{
			op:           lineno,
			name:         op,
			addr:         addr,
			cycles:       cycles,
			pageBoundary: pageBoundary,
		})
		key := op + addr
		names[key]++
	}
	return nil
}

func parseCycles(spec string) (cycles int, pageBoundary bool, err error) {
	l := len(spec)
	if spec[l-1] == '*' {
		pageBoundary = true
		spec = spec[:l-1]
	}
	cycles, err = strconv.Atoi(spec)
	return
}
