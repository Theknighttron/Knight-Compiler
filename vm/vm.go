package vm

import (
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
    if vm.sp == 0 {
        return nil
    }

    return vm.stack[vm.sp-1]
}