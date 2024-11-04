package vm

import (
	"fmt"
	"knightcompiler/code"
	"knightcompiler/compiler"
	"knightcompiler/object"
)


const StackSize = 2048

type VM struct {
    constants []object.Object
    instructions code.Instructions

    stack   []object.Object
    sp      int     // Always point to the next value on Top of stack is stack[sp-1]
}


func New(bytecode *compiler.Bytecode) *VM {
    return &VM{
        instructions: bytecode.Instructions,
        constants: bytecode.Constants,

        stack: make([]object.Object, StackSize),
        sp: 0,
    }
}

func (vm *VM) StackTop() object.Object {
    // check if the stack if empty
    if vm.sp == 0 {
        return nil
    }

    // return the top element in the stack
    return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error {
    // fetch the instruction
    for ip := 0; ip < len(vm.instructions); ip++ {  // iterate through the instructions pointer
        op := code.Opcode(vm.instructions[ip])      // turn the byte into an opcode

        // decode the operands in the bytecode
        switch op {
        case code.OpConstant:
            constIndex := code.ReadUint16(vm.instructions[ip+1:])
            ip +=2
            err := vm.push(vm.constants[constIndex])    // push the constants onto the stack
            if err != nil {
                return err
            }
        }
    }

    return nil
}

func (vm *VM) push(o object.Object) error {
    if vm.sp >= StackSize {
        return fmt.Errorf("stack overflow")
    }

    vm.stack[vm.sp] = o
    vm.sp++


    return nil
}
