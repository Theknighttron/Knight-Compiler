package knightcompiler

import (
    "fmt"
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
)

// definitions maps each Opcode to its corresponding Definition.
var definitions = map[Opcode]*Definition{
    OpConstant: {"OpConstant", []int{2}}, // OpConstant uses 2 bytes to store its operand.
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
