package compiler

import (
	"fmt"
	"knightcompiler/ast"
	"knightcompiler/code"
	"knightcompiler/lexer"
	"knightcompiler/parser"
	"testing"
)

type compilerTestCase struct {
    input                   string
    expectedConstants       []interface{}
    expectedInstructions    []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
    tests := []compilerTestCase{
        {
            input: "1 + 2",
            expectedConstants: []interface{}{1,2},
            expectedInstructions: []code.Instructions{
                // The instruction we expect the compiler to generate
                code.Make(code.OpConstant, 0),      // push constant 1 into the stack
                code.Make(code.OpConstant, 1),      // push constant 2 into the stack
            },
        },
    }

    runCompilerTests(t, tests)
}


func runCompilerTests(t *testing.T, tests []compilerTestCase) {
    t.Helper()

    for _, tt := range tests {
        // 1. Parse the input
        program := parse(tt.input)

        // 2. Create new compiler instance
        compiler := New()

        // 3. Compile the program
        err := compiler.Compile(program)
        if err != nil {
            t.Fatalf("Compiler error: %s", err)
        }

        // 4. Get the bytecode (instructions and constant) from the compiler
        bytecode := compiler.Bytecode()

        // 5. Test the instructions
        err := testInstructions(tt.expectedInstructions, bytecode.Instructions)
        if err != nil {
            t.Fatalf("Test instructions failed %s", err)
        }

        // 6. Test the constants
        err := testConstant(t, tt.expectedConstants, bytecode.Constants)
        if err != nil {
            t.Fatalf("Test Constants failed %s", err)
        }
    }
}

// Convert string into AST
func parse(input string) *ast.Program {
    l := lexer.New(input)
    p := parser.New(l)
    return p.ParseProgram()
}


// Comparte expected with actual bytecode
func testInstructions( expected []code.Instructions, actual code.Instructions) error {
    // combine all the expected instructions
    concatted := concatInstructions(expected)

    // Check if the length of the actual instructions and concatted expected instructions match
    if len(actual) != len(concatted) {
        return fmt.Errorf("Wrong instructions length. \nwant=%q \ngot= %q", concatted, actual)
    }

    // compare each byte of instructions
    for i, inst := range concatted {
        if actual[i] != inst {
            return fmt.Errorf("Wrong instruction at %d.\nwant=%q\nngot = %q", i, concatted, actual)
        }
    }

    // If everthing matches
    return nil
}

// Takes slice of instruction slices and concatenates them into a single slice
func concatInstructions(s []code.Instructions) code.Instructions {
    // Create an empyt slice to store the results
    out := code.Instructions{}

    for _, ins := range s {
        out = append(out, ins...)
    }

    return out
}

func testConstant() {

}



