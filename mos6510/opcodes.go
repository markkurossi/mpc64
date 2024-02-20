//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package mos6510

// Opcode defines 6510 opcodes.
type Opcode byte

// 6510 opcodes.
const (
	OpBRK        Opcode = iota // BRK 7
	OpORAizx                   // ORA izx 6
	OpKIL0x02                  // KIL
	OpSLOizx                   // SLO izx 8
	OpNOPzp0x04                // NOP zp 3
	OpORAzp                    // ORA zp 3
	OpASLzp                    // ASL zp 5
	OpSLOzp                    // SLO zp 5
	OpPHP                      // PHP 3
	OpORAimm                   // ORA imm 2
	OpASL                      // ASL 2
	OpANCimm0x0B               // ANC imm 2
	OpNOPabs                   // NOP abs 4
	OpORAabs                   // ORA abs 4
	OpASLabs                   // ASL abs 6
	OpSLOabs                   // SLO abs 6
	OpBPLrel                   // BPL rel 2*
	OpORAizy                   // ORA izy 5*
	OpKIL0x12                  // KIL
	OpSLOizy                   // SLO izy 8
	OpNOPzpx0x14               // NOP zpx 4
	OpORAzpx                   // ORA zpx 4
	OpASLzpx                   // ASL zpx 6
	OpSLOzpx                   // SLO zpx 6
	OpCLC                      // CLC 2
	OpORAaby                   // ORA aby 4*
	OpNOP0x1A                  // NOP 2
	OpSLOaby                   // SLO aby 7
	OpNOPabx0x1C               // NOP abx 4*
	OpORAabx                   // ORA abx 4*
	OpASLabx                   // ASL abx 7
	OpSLOabx                   // SLO abx 7
	OpJSRabs                   // JSR abs 6
	OpANDizx                   // AND izx 6
	OpKIL0x22                  // KIL
	OpRLAizx                   // RLA izx 8
	OpBITzp                    // BIT zp 3
	OpANDzp                    // AND zp 3
	OpROLzp                    // ROL zp 5
	OpRLAzp                    // RLA zp 5
	OpPLP                      // PLP 4
	OpANDimm                   // AND imm 2
	OpROL                      // ROL 2
	OpANCimm0x2B               // ANC imm 2
	OpBITabs                   // BIT abs 4
	OpANDabs                   // AND abs 4
	OpROLabs                   // ROL abs 6
	OpRLAabs                   // RLA abs 6
	OpBMIrel                   // BMI rel 2*
	OpANDizy                   // AND izy 5*
	OpKIL0x32                  // KIL
	OpRLAizy                   // RLA izy 8
	OpNOPzpx0x34               // NOP zpx 4
	OpANDzpx                   // AND zpx 4
	OpROLzpx                   // ROL zpx 6
	OpRLAzpx                   // RLA zpx 6
	OpSEC                      // SEC 2
	OpANDaby                   // AND aby 4*
	OpNOP0x3A                  // NOP 2
	OpRLAaby                   // RLA aby 7
	OpNOPabx0x3C               // NOP abx 4*
	OpANDabx                   // AND abx 4*
	OpROLabx                   // ROL abx 7
	OpRLAabx                   // RLA abx 7
	OpRTI                      // RTI 6
	OpEORizx                   // EOR izx 6
	OpKIL0x42                  // KIL
	OpSREizx                   // SRE izx 8
	OpNOPzp0x44                // NOP zp 3
	OpEORzp                    // EOR zp 3
	OpLSRzp                    // LSR zp 5
	OpSREzp                    // SRE zp 5
	OpPHA                      // PHA 3
	OpEORimm                   // EOR imm 2
	OpLSR                      // LSR 2
	OpALRimm                   // ALR imm 2
	OpJMPabs                   // JMP abs 3
	OpEORabs                   // EOR abs 4
	OpLSRabs                   // LSR abs 6
	OpSREabs                   // SRE abs 6
	OpBVCrel                   // BVC rel 2*
	OpEORizy                   // EOR izy 5*
	OpKIL0x52                  // KIL
	OpSREizy                   // SRE izy 8
	OpNOPzpx0x54               // NOP zpx 4
	OpEORzpx                   // EOR zpx 4
	OpLSRzpx                   // LSR zpx 6
	OpSREzpx                   // SRE zpx 6
	OpCLI                      // CLI 2
	OpEORaby                   // EOR aby 4*
	OpNOP0x5A                  // NOP 2
	OpSREaby                   // SRE aby 7
	OpNOPabx0x5C               // NOP abx 4*
	OpEORabx                   // EOR abx 4*
	OpLSRabx                   // LSR abx 7
	OpSREabx                   // SRE abx 7
	OpRTS                      // RTS 6
	OpADCizx                   // ADC izx 6
	OpKIL0x62                  // KIL
	OpRRAizx                   // RRA izx 8
	OpNOPzp0x64                // NOP zp 3
	OpADCzp                    // ADC zp 3
	OpRORzp                    // ROR zp 5
	OpRRAzp                    // RRA zp 5
	OpPLA                      // PLA 4
	OpADCimm                   // ADC imm 2
	OpROR                      // ROR 2
	OpARRimm                   // ARR imm 2
	OpJMPind                   // JMP ind 5
	OpADCabs                   // ADC abs 4
	OpRORabs                   // ROR abs 6
	OpRRAabs                   // RRA abs 6
	OpBVSrel                   // BVS rel 2*
	OpADCizy                   // ADC izy 5*
	OpKIL0x72                  // KIL
	OpRRAizy                   // RRA izy 8
	OpNOPzpx0x74               // NOP zpx 4
	OpADCzpx                   // ADC zpx 4
	OpRORzpx                   // ROR zpx 6
	OpRRAzpx                   // RRA zpx 6
	OpSEI                      // SEI 2
	OpADCaby                   // ADC aby 4*
	OpNOP0x7A                  // NOP 2
	OpRRAaby                   // RRA aby 7
	OpNOPabx0x7C               // NOP abx 4*
	OpADCabx                   // ADC abx 4*
	OpRORabx                   // ROR abx 7
	OpRRAabx                   // RRA abx 7
	OpNOPimm0x80               // NOP imm 2
	OpSTAizx                   // STA izx 6
	OpNOPimm0x82               // NOP imm 2
	OpSAXizx                   // SAX izx 6
	OpSTYzp                    // STY zp 3
	OpSTAzp                    // STA zp 3
	OpSTXzp                    // STX zp 3
	OpSAXzp                    // SAX zp 3
	OpDEY                      // DEY 2
	OpNOPimm0x89               // NOP imm 2
	OpTXA                      // TXA 2
	OpXAAimm                   // XAA imm 2
	OpSTYabs                   // STY abs 4
	OpSTAabs                   // STA abs 4
	OpSTXabs                   // STX abs 4
	OpSAXabs                   // SAX abs 4
	OpBCCrel                   // BCC rel 2*
	OpSTAizy                   // STA izy 6
	OpKIL0x92                  // KIL
	OpAHXizy                   // AHX izy 6
	OpSTYzpx                   // STY zpx 4
	OpSTAzpx                   // STA zpx 4
	OpSTXzpy                   // STX zpy 4
	OpSAXzpy                   // SAX zpy 4
	OpTYA                      // TYA 2
	OpSTAaby                   // STA aby 5
	OpTXS                      // TXS 2
	OpTASaby                   // TAS aby 5
	OpSHYabx                   // SHY abx 5
	OpSTAabx                   // STA abx 5
	OpSHXaby                   // SHX aby 5
	OpAHXaby                   // AHX aby 5
	OpLDYimm                   // LDY imm 2
	OpLDAizx                   // LDA izx 6
	OpLDXimm                   // LDX imm 2
	OpLAXizx                   // LAX izx 6
	OpLDYzp                    // LDY zp 3
	OpLDAzp                    // LDA zp 3
	OpLDXzp                    // LDX zp 3
	OpLAXzp                    // LAX zp 3
	OpTAY                      // TAY 2
	OpLDAimm                   // LDA imm 2
	OpTAX                      // TAX 2
	OpLAXimm                   // LAX imm 2
	OpLDYabs                   // LDY abs 4
	OpLDAabs                   // LDA abs 4
	OpLDXabs                   // LDX abs 4
	OpLAXabs                   // LAX abs 4
	OpBCSrel                   // BCS rel 2*
	OpLDAizy                   // LDA izy 5*
	OpKIL0xB2                  // KIL
	OpLAXizy                   // LAX izy 5*
	OpLDYzpx                   // LDY zpx 4
	OpLDAzpx                   // LDA zpx 4
	OpLDXzpy                   // LDX zpy 4
	OpLAXzpy                   // LAX zpy 4
	OpCLV                      // CLV 2
	OpLDAaby                   // LDA aby 4*
	OpTSX                      // TSX 2
	OpLASaby                   // LAS aby 4*
	OpLDYabx                   // LDY abx 4*
	OpLDAabx                   // LDA abx 4*
	OpLDXaby                   // LDX aby 4*
	OpLAXaby                   // LAX aby 4*
	OpCPYimm                   // CPY imm 2
	OpCMPizx                   // CMP izx 6
	OpNOPimm0xC2               // NOP imm 2
	OpDCPizx                   // DCP izx 8
	OpCPYzp                    // CPY zp 3
	OpCMPzp                    // CMP zp 3
	OpDECzp                    // DEC zp 5
	OpDCPzp                    // DCP zp 5
	OpINY                      // INY 2
	OpCMPimm                   // CMP imm 2
	OpDEX                      // DEX 2
	OpAXSimm                   // AXS imm 2
	OpCPYabs                   // CPY abs 4
	OpCMPabs                   // CMP abs 4
	OpDECabs                   // DEC abs 6
	OpDCPabs                   // DCP abs 6
	OpBNErel                   // BNE rel 2*
	OpCMPizy                   // CMP izy 5*
	OpKIL0xD2                  // KIL
	OpDCPizy                   // DCP izy 8
	OpNOPzpx0xD4               // NOP zpx 4
	OpCMPzpx                   // CMP zpx 4
	OpDECzpx                   // DEC zpx 6
	OpDCPzpx                   // DCP zpx 6
	OpCLD                      // CLD 2
	OpCMPaby                   // CMP aby 4*
	OpNOP0xDA                  // NOP 2
	OpDCPaby                   // DCP aby 7
	OpNOPabx0xDC               // NOP abx 4*
	OpCMPabx                   // CMP abx 4*
	OpDECabx                   // DEC abx 7
	OpDCPabx                   // DCP abx 7
	OpCPXimm                   // CPX imm 2
	OpSBCizx                   // SBC izx 6
	OpNOPimm0xE2               // NOP imm 2
	OpISCizx                   // ISC izx 8
	OpCPXzp                    // CPX zp 3
	OpSBCzp                    // SBC zp 3
	OpINCzp                    // INC zp 5
	OpISCzp                    // ISC zp 5
	OpINX                      // INX 2
	OpSBCimm0xE9               // SBC imm 2
	OpNOP0xEA                  // NOP 2
	OpSBCimm0xEB               // SBC imm 2
	OpCPXabs                   // CPX abs 4
	OpSBCabs                   // SBC abs 4
	OpINCabs                   // INC abs 6
	OpISCabs                   // ISC abs 6
	OpBEQrel                   // BEQ rel 2*
	OpSBCizy                   // SBC izy 5*
	OpKIL0xF2                  // KIL
	OpISCizy                   // ISC izy 8
	OpNOPzpx0xF4               // NOP zpx 4
	OpSBCzpx                   // SBC zpx 4
	OpINCzpx                   // INC zpx 6
	OpISCzpx                   // ISC zpx 6
	OpSED                      // SED 2
	OpSBCaby                   // SBC aby 4*
	OpNOP0xFA                  // NOP 2
	OpISCaby                   // ISC aby 7
	OpNOPabx0xFC               // NOP abx 4*
	OpSBCabx                   // SBC abx 4*
	OpINCabx                   // INC abx 7
	OpISCabx                   // ISC abx 7
)

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
	{
		Op:           OpBRK,
		Name:         "BRK",
		Addr:         AddrImp,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpORAizx,
		Name:         "ORA",
		Addr:         AddrIZX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpKIL0x02,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpSLOizx,
		Name:         "SLO",
		Addr:         AddrIZX,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpNOPzp0x04,
		Name:         "NOP",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpORAzp,
		Name:         "ORA",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpASLzp,
		Name:         "ASL",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpSLOzp,
		Name:         "SLO",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpPHP,
		Name:         "PHP",
		Addr:         AddrImp,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpORAimm,
		Name:         "ORA",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpASL,
		Name:         "ASL",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpANCimm0x0B,
		Name:         "ANC",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpNOPabs,
		Name:         "NOP",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpORAabs,
		Name:         "ORA",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpASLabs,
		Name:         "ASL",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpSLOabs,
		Name:         "SLO",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpBPLrel,
		Name:         "BPL",
		Addr:         AddrREL,
		Cycles:       2,
		PageBoundary: true,
	},
	{
		Op:           OpORAizy,
		Name:         "ORA",
		Addr:         AddrIZY,
		Cycles:       5,
		PageBoundary: true,
	},
	{
		Op:           OpKIL0x12,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpSLOizy,
		Name:         "SLO",
		Addr:         AddrIZY,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpNOPzpx0x14,
		Name:         "NOP",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpORAzpx,
		Name:         "ORA",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpASLzpx,
		Name:         "ASL",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpSLOzpx,
		Name:         "SLO",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpCLC,
		Name:         "CLC",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpORAaby,
		Name:         "ORA",
		Addr:         AddrABY,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpNOP0x1A,
		Name:         "NOP",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpSLOaby,
		Name:         "SLO",
		Addr:         AddrABY,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpNOPabx0x1C,
		Name:         "NOP",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpORAabx,
		Name:         "ORA",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpASLabx,
		Name:         "ASL",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpSLOabx,
		Name:         "SLO",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpJSRabs,
		Name:         "JSR",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpANDizx,
		Name:         "AND",
		Addr:         AddrIZX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpKIL0x22,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpRLAizx,
		Name:         "RLA",
		Addr:         AddrIZX,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpBITzp,
		Name:         "BIT",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpANDzp,
		Name:         "AND",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpROLzp,
		Name:         "ROL",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpRLAzp,
		Name:         "RLA",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpPLP,
		Name:         "PLP",
		Addr:         AddrImp,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpANDimm,
		Name:         "AND",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpROL,
		Name:         "ROL",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpANCimm0x2B,
		Name:         "ANC",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpBITabs,
		Name:         "BIT",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpANDabs,
		Name:         "AND",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpROLabs,
		Name:         "ROL",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpRLAabs,
		Name:         "RLA",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpBMIrel,
		Name:         "BMI",
		Addr:         AddrREL,
		Cycles:       2,
		PageBoundary: true,
	},
	{
		Op:           OpANDizy,
		Name:         "AND",
		Addr:         AddrIZY,
		Cycles:       5,
		PageBoundary: true,
	},
	{
		Op:           OpKIL0x32,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpRLAizy,
		Name:         "RLA",
		Addr:         AddrIZY,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpNOPzpx0x34,
		Name:         "NOP",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpANDzpx,
		Name:         "AND",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpROLzpx,
		Name:         "ROL",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpRLAzpx,
		Name:         "RLA",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpSEC,
		Name:         "SEC",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpANDaby,
		Name:         "AND",
		Addr:         AddrABY,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpNOP0x3A,
		Name:         "NOP",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpRLAaby,
		Name:         "RLA",
		Addr:         AddrABY,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpNOPabx0x3C,
		Name:         "NOP",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpANDabx,
		Name:         "AND",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpROLabx,
		Name:         "ROL",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpRLAabx,
		Name:         "RLA",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpRTI,
		Name:         "RTI",
		Addr:         AddrImp,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpEORizx,
		Name:         "EOR",
		Addr:         AddrIZX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpKIL0x42,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpSREizx,
		Name:         "SRE",
		Addr:         AddrIZX,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpNOPzp0x44,
		Name:         "NOP",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpEORzp,
		Name:         "EOR",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpLSRzp,
		Name:         "LSR",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpSREzp,
		Name:         "SRE",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpPHA,
		Name:         "PHA",
		Addr:         AddrImp,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpEORimm,
		Name:         "EOR",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpLSR,
		Name:         "LSR",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpALRimm,
		Name:         "ALR",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpJMPabs,
		Name:         "JMP",
		Addr:         AddrABS,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpEORabs,
		Name:         "EOR",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpLSRabs,
		Name:         "LSR",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpSREabs,
		Name:         "SRE",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpBVCrel,
		Name:         "BVC",
		Addr:         AddrREL,
		Cycles:       2,
		PageBoundary: true,
	},
	{
		Op:           OpEORizy,
		Name:         "EOR",
		Addr:         AddrIZY,
		Cycles:       5,
		PageBoundary: true,
	},
	{
		Op:           OpKIL0x52,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpSREizy,
		Name:         "SRE",
		Addr:         AddrIZY,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpNOPzpx0x54,
		Name:         "NOP",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpEORzpx,
		Name:         "EOR",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpLSRzpx,
		Name:         "LSR",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpSREzpx,
		Name:         "SRE",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpCLI,
		Name:         "CLI",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpEORaby,
		Name:         "EOR",
		Addr:         AddrABY,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpNOP0x5A,
		Name:         "NOP",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpSREaby,
		Name:         "SRE",
		Addr:         AddrABY,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpNOPabx0x5C,
		Name:         "NOP",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpEORabx,
		Name:         "EOR",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpLSRabx,
		Name:         "LSR",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpSREabx,
		Name:         "SRE",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpRTS,
		Name:         "RTS",
		Addr:         AddrImp,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpADCizx,
		Name:         "ADC",
		Addr:         AddrIZX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpKIL0x62,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpRRAizx,
		Name:         "RRA",
		Addr:         AddrIZX,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpNOPzp0x64,
		Name:         "NOP",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpADCzp,
		Name:         "ADC",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpRORzp,
		Name:         "ROR",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpRRAzp,
		Name:         "RRA",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpPLA,
		Name:         "PLA",
		Addr:         AddrImp,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpADCimm,
		Name:         "ADC",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpROR,
		Name:         "ROR",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpARRimm,
		Name:         "ARR",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpJMPind,
		Name:         "JMP",
		Addr:         AddrIND,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpADCabs,
		Name:         "ADC",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpRORabs,
		Name:         "ROR",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpRRAabs,
		Name:         "RRA",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpBVSrel,
		Name:         "BVS",
		Addr:         AddrREL,
		Cycles:       2,
		PageBoundary: true,
	},
	{
		Op:           OpADCizy,
		Name:         "ADC",
		Addr:         AddrIZY,
		Cycles:       5,
		PageBoundary: true,
	},
	{
		Op:           OpKIL0x72,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpRRAizy,
		Name:         "RRA",
		Addr:         AddrIZY,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpNOPzpx0x74,
		Name:         "NOP",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpADCzpx,
		Name:         "ADC",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpRORzpx,
		Name:         "ROR",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpRRAzpx,
		Name:         "RRA",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpSEI,
		Name:         "SEI",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpADCaby,
		Name:         "ADC",
		Addr:         AddrABY,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpNOP0x7A,
		Name:         "NOP",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpRRAaby,
		Name:         "RRA",
		Addr:         AddrABY,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpNOPabx0x7C,
		Name:         "NOP",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpADCabx,
		Name:         "ADC",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpRORabx,
		Name:         "ROR",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpRRAabx,
		Name:         "RRA",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpNOPimm0x80,
		Name:         "NOP",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpSTAizx,
		Name:         "STA",
		Addr:         AddrIZX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpNOPimm0x82,
		Name:         "NOP",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpSAXizx,
		Name:         "SAX",
		Addr:         AddrIZX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpSTYzp,
		Name:         "STY",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpSTAzp,
		Name:         "STA",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpSTXzp,
		Name:         "STX",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpSAXzp,
		Name:         "SAX",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpDEY,
		Name:         "DEY",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpNOPimm0x89,
		Name:         "NOP",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpTXA,
		Name:         "TXA",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpXAAimm,
		Name:         "XAA",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpSTYabs,
		Name:         "STY",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpSTAabs,
		Name:         "STA",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpSTXabs,
		Name:         "STX",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpSAXabs,
		Name:         "SAX",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpBCCrel,
		Name:         "BCC",
		Addr:         AddrREL,
		Cycles:       2,
		PageBoundary: true,
	},
	{
		Op:           OpSTAizy,
		Name:         "STA",
		Addr:         AddrIZY,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpKIL0x92,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpAHXizy,
		Name:         "AHX",
		Addr:         AddrIZY,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpSTYzpx,
		Name:         "STY",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpSTAzpx,
		Name:         "STA",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpSTXzpy,
		Name:         "STX",
		Addr:         AddrZPY,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpSAXzpy,
		Name:         "SAX",
		Addr:         AddrZPY,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpTYA,
		Name:         "TYA",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpSTAaby,
		Name:         "STA",
		Addr:         AddrABY,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpTXS,
		Name:         "TXS",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpTASaby,
		Name:         "TAS",
		Addr:         AddrABY,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpSHYabx,
		Name:         "SHY",
		Addr:         AddrABX,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpSTAabx,
		Name:         "STA",
		Addr:         AddrABX,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpSHXaby,
		Name:         "SHX",
		Addr:         AddrABY,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpAHXaby,
		Name:         "AHX",
		Addr:         AddrABY,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpLDYimm,
		Name:         "LDY",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpLDAizx,
		Name:         "LDA",
		Addr:         AddrIZX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpLDXimm,
		Name:         "LDX",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpLAXizx,
		Name:         "LAX",
		Addr:         AddrIZX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpLDYzp,
		Name:         "LDY",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpLDAzp,
		Name:         "LDA",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpLDXzp,
		Name:         "LDX",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpLAXzp,
		Name:         "LAX",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpTAY,
		Name:         "TAY",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpLDAimm,
		Name:         "LDA",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpTAX,
		Name:         "TAX",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpLAXimm,
		Name:         "LAX",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpLDYabs,
		Name:         "LDY",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpLDAabs,
		Name:         "LDA",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpLDXabs,
		Name:         "LDX",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpLAXabs,
		Name:         "LAX",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpBCSrel,
		Name:         "BCS",
		Addr:         AddrREL,
		Cycles:       2,
		PageBoundary: true,
	},
	{
		Op:           OpLDAizy,
		Name:         "LDA",
		Addr:         AddrIZY,
		Cycles:       5,
		PageBoundary: true,
	},
	{
		Op:           OpKIL0xB2,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpLAXizy,
		Name:         "LAX",
		Addr:         AddrIZY,
		Cycles:       5,
		PageBoundary: true,
	},
	{
		Op:           OpLDYzpx,
		Name:         "LDY",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpLDAzpx,
		Name:         "LDA",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpLDXzpy,
		Name:         "LDX",
		Addr:         AddrZPY,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpLAXzpy,
		Name:         "LAX",
		Addr:         AddrZPY,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpCLV,
		Name:         "CLV",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpLDAaby,
		Name:         "LDA",
		Addr:         AddrABY,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpTSX,
		Name:         "TSX",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpLASaby,
		Name:         "LAS",
		Addr:         AddrABY,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpLDYabx,
		Name:         "LDY",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpLDAabx,
		Name:         "LDA",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpLDXaby,
		Name:         "LDX",
		Addr:         AddrABY,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpLAXaby,
		Name:         "LAX",
		Addr:         AddrABY,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpCPYimm,
		Name:         "CPY",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpCMPizx,
		Name:         "CMP",
		Addr:         AddrIZX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpNOPimm0xC2,
		Name:         "NOP",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpDCPizx,
		Name:         "DCP",
		Addr:         AddrIZX,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpCPYzp,
		Name:         "CPY",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpCMPzp,
		Name:         "CMP",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpDECzp,
		Name:         "DEC",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpDCPzp,
		Name:         "DCP",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpINY,
		Name:         "INY",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpCMPimm,
		Name:         "CMP",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpDEX,
		Name:         "DEX",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpAXSimm,
		Name:         "AXS",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpCPYabs,
		Name:         "CPY",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpCMPabs,
		Name:         "CMP",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpDECabs,
		Name:         "DEC",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpDCPabs,
		Name:         "DCP",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpBNErel,
		Name:         "BNE",
		Addr:         AddrREL,
		Cycles:       2,
		PageBoundary: true,
	},
	{
		Op:           OpCMPizy,
		Name:         "CMP",
		Addr:         AddrIZY,
		Cycles:       5,
		PageBoundary: true,
	},
	{
		Op:           OpKIL0xD2,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpDCPizy,
		Name:         "DCP",
		Addr:         AddrIZY,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpNOPzpx0xD4,
		Name:         "NOP",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpCMPzpx,
		Name:         "CMP",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpDECzpx,
		Name:         "DEC",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpDCPzpx,
		Name:         "DCP",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpCLD,
		Name:         "CLD",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpCMPaby,
		Name:         "CMP",
		Addr:         AddrABY,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpNOP0xDA,
		Name:         "NOP",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpDCPaby,
		Name:         "DCP",
		Addr:         AddrABY,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpNOPabx0xDC,
		Name:         "NOP",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpCMPabx,
		Name:         "CMP",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpDECabx,
		Name:         "DEC",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpDCPabx,
		Name:         "DCP",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpCPXimm,
		Name:         "CPX",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpSBCizx,
		Name:         "SBC",
		Addr:         AddrIZX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpNOPimm0xE2,
		Name:         "NOP",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpISCizx,
		Name:         "ISC",
		Addr:         AddrIZX,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpCPXzp,
		Name:         "CPX",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpSBCzp,
		Name:         "SBC",
		Addr:         AddrZP,
		Cycles:       3,
		PageBoundary: false,
	},
	{
		Op:           OpINCzp,
		Name:         "INC",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpISCzp,
		Name:         "ISC",
		Addr:         AddrZP,
		Cycles:       5,
		PageBoundary: false,
	},
	{
		Op:           OpINX,
		Name:         "INX",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpSBCimm0xE9,
		Name:         "SBC",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpNOP0xEA,
		Name:         "NOP",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpSBCimm0xEB,
		Name:         "SBC",
		Addr:         AddrIMM,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpCPXabs,
		Name:         "CPX",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpSBCabs,
		Name:         "SBC",
		Addr:         AddrABS,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpINCabs,
		Name:         "INC",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpISCabs,
		Name:         "ISC",
		Addr:         AddrABS,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpBEQrel,
		Name:         "BEQ",
		Addr:         AddrREL,
		Cycles:       2,
		PageBoundary: true,
	},
	{
		Op:           OpSBCizy,
		Name:         "SBC",
		Addr:         AddrIZY,
		Cycles:       5,
		PageBoundary: true,
	},
	{
		Op:           OpKIL0xF2,
		Name:         "KIL",
		Addr:         AddrImp,
		Cycles:       0,
		PageBoundary: false,
	},
	{
		Op:           OpISCizy,
		Name:         "ISC",
		Addr:         AddrIZY,
		Cycles:       8,
		PageBoundary: false,
	},
	{
		Op:           OpNOPzpx0xF4,
		Name:         "NOP",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpSBCzpx,
		Name:         "SBC",
		Addr:         AddrZPX,
		Cycles:       4,
		PageBoundary: false,
	},
	{
		Op:           OpINCzpx,
		Name:         "INC",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpISCzpx,
		Name:         "ISC",
		Addr:         AddrZPX,
		Cycles:       6,
		PageBoundary: false,
	},
	{
		Op:           OpSED,
		Name:         "SED",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpSBCaby,
		Name:         "SBC",
		Addr:         AddrABY,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpNOP0xFA,
		Name:         "NOP",
		Addr:         AddrImp,
		Cycles:       2,
		PageBoundary: false,
	},
	{
		Op:           OpISCaby,
		Name:         "ISC",
		Addr:         AddrABY,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpNOPabx0xFC,
		Name:         "NOP",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpSBCabx,
		Name:         "SBC",
		Addr:         AddrABX,
		Cycles:       4,
		PageBoundary: true,
	},
	{
		Op:           OpINCabx,
		Name:         "INC",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
	{
		Op:           OpISCabx,
		Name:         "ISC",
		Addr:         AddrABX,
		Cycles:       7,
		PageBoundary: false,
	},
}
