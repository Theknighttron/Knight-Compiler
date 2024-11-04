package code

import (
    "fmt"
    "bytes"
    "encoding/binary"
)

// Instructions represents a series of bytes as machine code instructions.
type Instructions []byte

// Opcode is a byte representing a specific instruction (operation).
type Opcode byte

// Definition represents an instruction's metadata, including its name and the width (in bytes)
// of each operand that follows the instruction in a program.
type Definition struct {
    Name          string   // The name of the instruction.
    OperandWidths []int    // The number of bytes each operand takes.
}

const (
    OpConstant Opcode = iota // OpConstant represents an instruction to push a constant onto the stack.
    OpAdd
)

// definitions maps each Opcode to its corresponding Definition.
var definitions = map[Opcode]*Definition{
    OpConstant: {"OpConstant", []int{2}}, // OpConstant uses 2 bytes to store its operand.
    OpAdd:      {"OpAdd", []int{}},
}

// Lookup finds the definition for a given opcode.
func Lookup(op byte) (*Definition, error) {
    def, ok := definitions[Opcode(op)]
    if !ok {
        return nil, fmt.Errorf("opcode %d undefined", op)
    }

    return def, nil
}

// Make generates the bytecode for a given opcode and its operands.
// It constructs an instruction by encoding the opcode followed by the operands in the appropriate byte widths.
func Make(op Opcode, operands ...int) []byte {
    // Lookup the definition for the opcode.
    def, ok := definitions[op]
    if !ok {
        return []byte{}
    }

    // Calculate the total length of the instruction (opcode + operand widths).
    instructionLen := 1 // 1 byte for the opcode.
    for _, w := range def.OperandWidths {
        instructionLen += w // Add the width of each operand.
    }

    // Create a byte slice to hold the full instruction.
    instruction := make([]byte, instructionLen)
    instruction[0] = byte(op) // Set the opcode as the first byte.

    offset := 1 // Start placing operands after the opcode.
    // Encode each operand according to its width defined in the OperandWidths.
    for i, o := range operands {
        width := def.OperandWidths[i]
        switch width {
        case 2:
            // For 2-byte operands, use BigEndian encoding.
            binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
        }
        offset += width // Move the offset forward by the operand's width.
    }

    return instruction
}

// String representation of instructions
func (ins Instructions) String() string {
    var out bytes.Buffer

    i := 0
    for i < len(ins) {
        def, err := Lookup(ins[i])
        if err != nil {
            fmt.Fprintf(&out, "ERROR: %s\n", err)
            continue
        }

        operands, read := ReadOperands(def, ins[i+1:])

        fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))
         i +=  1 + read
    }

    return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
    operandCount := len(def.OperandWidths)

    if len(operands) != operandCount {
        return fmt.Sprintf("Error: operand len %d does not match defined %d\n", len(operands), operandCount)
    }

    switch operandCount {
    case 0:
        return def.Name
    case 1:
        return fmt.Sprintf("%s %d", def.Name, operands[0])
    }

    return fmt.Sprintf("Error: unhandled operandCount for %s\n", def.Name)
}

// Decode the operands along with the total number of bytes read
func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
    operands := make([]int, len(def.OperandWidths))
    offset := 0

    for i, width := range def.OperandWidths {
        switch width {
        case 2:
            operands[i] = int(ReadUint16(ins[offset:]))
        }

        offset += width
    }

    return operands, offset
}

// check for correct operands and number of byte reads
func ReadUint16(ins Instructions) uint16 {
    return binary.BigEndian.Uint16(ins)
}
