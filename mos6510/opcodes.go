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
	// 0x00-0x0F
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

	// 0x10-0x1F
	OpBPLrel     // BPL rel 2*
	OpORAizy     // ORA izy 5*
	OpKIL0x12    // KIL
	OpSLOizy     // SLO izy 8
	OpNOPzpx0x14 // NOP zpx 4
	OpORAzpx     // ORA zpx 4
	OpASLzpx     // ASL zpx 6
	OpSLOzpx     // SLO zpx 6
	OpCLC        // CLC 2
	OpORAaby     // ORA aby 4*
	OpNOP1a      // NOP 2
	OpSLOaby     // SLO aby 7
	OpNOPabx0x1C // NOP abx 4*
	OpORAabx     // ORA abx 4*
	OpASLabx     // ASL abx 7
	OpSLOabx     // SLO abx 7

	// 0x20-0x2F
	OpJSRabs     // JSR abs 6
	OpANDizx     // AND izx 6
	OpKIL0x22    // KIL
	OpRLAizx     // RLA izx 8
	OpBITzp      // BIT zp 3
	OpANDzp      // AND zp 3
	OpROLzp      // ROL zp 5
	OpRLAzp      // RLA zp 5
	OpPLP        // PLP 4
	OpANDimm     // AND imm 2
	OpROL        // ROL 2
	OpANCimm0x2b // ANC imm 2
	OpBITabs     // BIT abs 4
	OpANDabs     // AND abs 4
	OpROLabs     // ROL abs 6
	OpRLAabs     // RLA abs 6

	// 0x30-0x3F
	OpBMIrel     // BMI rel 2*
	OpANDizy     // AND izy 5*
	OpKIL0x32    // KIL
	OpRLAizy     // RLA izy 8
	OpNOPzpx0x34 // NOP zpx 4
	OpANDzpx     // AND zpx 4
	OpROLzpx     // ROL zpx 6
	OpRLAzpx     // RLA zpx 6
	OpSEC        // SEC 2
	OpANDaby     // AND aby 4*
	OpNOP0x3A    // NOP 2
	OpRLAaby     // RLA aby 7
	OpNOPabx0x3C // NOP abx 4*
	OpANDabx     // AND abx 4*
	OpROLabx     // ROL abx 7
	OpRLAabx     // RLA abx 7

	// 0x40-0x4F
	OpRTI       // RTI 6
	OpEORizx    // EOR izx 6
	OpKIL0x42   // KIL
	OpSREizx    // SRE izx 8
	OpNOPzp0x44 // NOP zp 3
	OpEORzp     // EOR zp 3
	OpLSRzp     // LSR zp 5
	OpSREzp     // SRE zp 5
	OpPHA       // PHA 3
	OpEORimm    // EOR imm 2
	OpLSR       // LSR 2
	OpALRimm    // ALR imm 2
	OpJMPabs    // JMP abs 3
	OpEORabs    // EOR abs 4
	OpLSRabs    // LSR abs 6
	OpSREabs    // SRE abs 6

	// 0x50-0x5F
	OpBVCrel     // BVC rel 2*
	OpEORizy     // EOR izy 5*
	OpKIL0x52    // KIL
	OpSREizy     // SRE izy 8
	OpNOPzpx0x54 // NOP zpx 4
	OpEORzpx     // EOR zpx 4
	OpLSRzpx     // LSR zpx 6
	OpSREzpx     // SRE zpx 6
	OpCLI        // CLI 2
	OpEORaby     // EOR aby 4*
	OpNOP5a      // NOP 2
	OpSREaby     // SRE aby 7
	OpNOPabx0x5C // NOP abx 4*
	OpEORabx     // EOR abx 4*
	OpLSRabx     // LSR abx 7
	OpSREabx     // SRE abx

	// 0x60-0x6F
	OpRTS       // RTS 6
	OpADCizx    // ADC izx 6
	OpKIL0x62   // KIL
	OpRRAizx    // RRA izx 8
	OpNOPzp0x64 // NOP zp 3
	OpADCzp     // ADC zp 3
	OpRORzp     // ROR zp 5
	OpRRAzp     // RRA zp 5
	OpPLA       // PLA 4
	OpADCimm    // ADC imm 2
	OpROR       // ROR 2
	OpARRimm    // ARR imm 2
	OpJMPind    // JMP ind 5
	OpADCabs    // ADC abs 4
	OpRORabs    // ROR abs 6
	OpRRAabs    // RRA abs 6

	// 0x70-0x7F
	OpBVSrel     // BVS rel 2*
	OpADCizy     // ADC izy 5*
	OpKIL0x72    // KIL
	OpRRAizy     // RRA izy 8
	OpNOPzpx0x64 // NOP zpx 4
	OpADCzpx     // ADC zpx 4
	OpRORzpx     // ROR zpx 6
	OpRRAzpx     // RRA zpx 6
	OpSEI        // SEI 2
	OpADCaby     // ADC aby 4*
	OpNOP0x7a    // NOP 2
	OpRRAaby     // RRA aby 7
	OpNOPabx0x7C // NOP abx 4*
	OpADCabx     // ADC abx 4*
	OpRORabx     // ROR abx 7
	OpRRAabx     // RRA abx 7

	// 0x80-0x8F
	OpNOPimm0x80 // NOP imm 2
	OpSTAizx     // STA izx 6
	OpNOPimm0x82 // NOP imm 2
	OpSAXizx     // SAX izx 6
	OpSTYzp      // STY zp 3
	OpSTAzp      // STA zp 3
	OpSTXzp      // STX zp 3
	OpSAXzp      // SAX zp 3
	OpDEY        // DEY 2
	OpNOPimm0x89 // NOP imm 2
	OpTXA        // TXA 2
	OpXAAimm     // XAA imm 2
	OpSTYabs     // STY abs 4
	OpSTAabs     // STA abs 4
	OpSTXabs     // STX abs 4
	OpSAXabs     // SAX abs 4

	// 0x90-0x9F
	OpBCCrel  // BCC rel 2*
	OpSTAizy  // STA izy 6
	OpKIL0x92 // KIL
	OpAHXizy  // AHX izy 6
	OpSTYzpx  // STY zpx 4
	OpSTAzpx  // STA zpx 4
	OpSTXzpy  // STX zpy 4
	OpSAXzpy  // SAX zpy 4
	OpTYA     // TYA 2
	OpSTAaby  // STA aby 5
	OpTXS     // TXS 2
	OpTASaby  // TAS aby 5
	OpSHYabx  // SHY abx 5
	OpSTAabx  // STA abx 5
	OpSHXaby  // SHX aby 5
	OpAHXaby  // AHX aby 5

	// 0xA0-0xAF
	OpLDYimm // LDY imm 2
	OpLDAizx // LDA izx 6
	OpLDXimm // LDX imm 2
	OpLAXizx // LAX izx 6
	OpLDYzp  // LDY zp 3
	OpLDAzp  // LDA zp 3
	OpLDXzp  // LDX zp 3
	OpLAXzp  // LAX zp 3
	OpTAY    // TAY 2
	OpLDAimm // LDA imm 2
	OpTAX    // TAX 2
	OpLAXimm // LAX imm 2
	OpLDYabs // LDY abs 4
	OpLDAabs // LDA abs 4
	OpLDXabs // LDX abs 4
	OpLAXabs // LAX abs 4

	// 0xB0-0xBF
	OpBCSrel  // BCS rel 2*
	OpLDAizy  // LDA izy 5*
	OpKIL0xB2 // KIL
	OpLAXizy  // LAX izy 5*
	OpLDYzpx  // LDY zpx 4
	OpLDAzpx  // LDA zpx 4
	OpLDXzpy  // LDX zpy 4
	OpLAXzpy  // LAX zpy 4
	OpCLV     // CLV 2
	OpLDAaby  // LDA aby 4*
	OpTSX     // TSX 2
	OpLASaby  // LAS aby 4*
	OpLDYabx  // LDY abx 4*
	OpLDAabx  // LDA abx 4*
	OpLDXaby  // LDX aby 4*
	OpLAXaby  // LAX aby 4*

	// 0xC0-0xCF
	OpCPYimm     // CPY imm 2
	OpCMPizx     // CMP izx 6
	OpNOPimm0xC2 // NOP imm 2
	OpDCPizx     // DCP izx 8
	OpCPYzp      // CPY zp 3
	OpCMPzp      // CMP zp 3
	OpDECzp      // DEC zp 5
	OpDCPzp      // DCP zp 5
	OpINY        // INY 2
	OpCMPimm     // CMP imm 2
	OpDEX        // DEX 2
	OpAXSimm     // AXS imm 2
	OpCPYabs     // CPY abs 4
	OpCMPabs     // CMP abs 4
	OpDECabs     // DEC abs 6
	OpDCPabs     // DCP abs 6

	// 0xD0-0xDF
	OpBNErel     // BNE rel 2*
	OpCMPizy     // CMP izy 5*
	OpKIL0xD2    // KIL
	OpDCPizy     // DCP izy 8
	OpNOPzpx0xD4 // NOP zpx 4
	OpCMPzpx     // CMP zpx 4
	OpDECzpx     // DEC zpx 6
	OpDCPzpx     // DCP zpx 6
	OpCLD        // CLD 2
	OpCMPaby     // CMP aby 4*
	OpNOP0xDA    // NOP 2
	OpDCPaby     // DCP aby 7
	OpNOPabx0xDC // NOP abx 4*
	OpCMPabx     // CMP abx 4*
	OpDECabx     // DEC abx 7
	OpDCPabx     // DCP abx 7

	// 0xE0-0xEF
	OpCPXimm     // CPX imm 2
	OpSBCizx     // SBC izx 6
	OpNOPimm0xE2 // NOP imm 2
	OpISCizx     // ISC izx 8
	OpCPXzp      // CPX zp 3
	OpSBCzp      // SBC zp 3
	OpINCzp      // INC zp 5
	OpISCzp      // ISC zp 5
	OpINX        // INX 2
	OpSBCimm0xE9 // SBC imm 2
	OpNOP0xEA    // NOP 2
	OpSBCimm0xEB // SBC imm 2
	OpCPXabs     // CPX abs 4
	OpSBCabs     // SBC abs 4
	OpINCabs     // INC abs 6
	OpISCabs     // ISC abs 6

	// 0xF0-0xFF
	OpBEQrel     // BEQ rel 2*
	OpSBCizy     // SBC izy 5*
	OpKIL0xF2    // KIL
	OpISCizy     // ISC izy 8
	OpNOPzpx0xF4 // NOP zpx 4
	OpSBCzpx     // SBC zpx 4
	OpINCzpx     // INC zpx 6
	OpISCzpx     // ISC zpx 6
	OpSED        // SED 2
	OpSBCaby     // SBC aby 4*
	OpNOP0xFA    // NOP 2
	OpISCaby     // ISC aby 7
	OpNOPabx0xFC // NOP abx 4*
	OpSBCabx     // SBC abx 4*
	OpINCabx     // INC abx 7
	OpISCabx     // ISC abx 7
)
