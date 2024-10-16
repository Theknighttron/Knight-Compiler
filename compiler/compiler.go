package compiler

import (
    "knightcompiler/ast"
    "knightcompiler/code"
    "knightcompiler/object"
)

type Compiler struct {
    instructions code.Instructions  // store the compiled bytecode
    constants []object.Object       // object to store contants values during compilation
}


type Bytecode struct {
    Instructions code.Instructions
    Constants []object.Object
}

// Create a new instance of a compiler
func New()*Compiler {
    return &Compiler{
        instructions: code.Instructions{},
        constants: []object.Object{},
    }
}

func (c *Compiler) Compile(node ast.Node) error {
    return nil
}

func (c *Compiler) Bytecode() *Bytecode {
    return &Bytecode {
        Instructions: c.instructions,
        Constants: c.constants,
    }
}





