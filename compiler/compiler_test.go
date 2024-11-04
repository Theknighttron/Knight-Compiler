package compiler

import (
	"fmt"
	"knightcompiler/ast"
	"knightcompiler/code"
	"knightcompiler/lexer"
	"knightcompiler/object"
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
            input:                  "1 + 2",
            expectedConstants:      []interface{}{1,2},
            expectedInstructions:   []code.Instructions{
                // The instruction we expect the compiler to generate
                code.Make(code.OpConstant, 0),      // push constant 1 into the stack
                code.Make(code.OpConstant, 1),      // push constant 2 into the stack
                code.Make(code.OpAdd),              // Add the top stack values
            },
        },
    }

    runCompilerTests(t, tests)
}


func runCompilerTests(t *testing.T, tests []compilerTestCase) {
    t.Helper()

    for _, tt := range tests {
        // 1. Parse the input to produce an AST
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
        err = testInstructions(tt.expectedInstructions, bytecode.Instructions)
        if err != nil {
            t.Fatalf("Test instructions failed %s", err)
        }

        // 6. Test the constants
        err = testConstants(t, tt.expectedConstants, bytecode.Constants)
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

func testConstants(t *testing.T, expected []interface{}, actual []object.Object) error {
    if len(expected) != len(actual) {
        return fmt.Errorf("Wrong number of constants. got=%d, want=%d", len(actual), len(expected))
    }

    for i, constant := range expected {
        switch constant := constant.(type) {
        case int:
            err := testIntegerObject(int64(constant), actual[i])
            if err != nil {
                return fmt.Errorf("Constant %d - testIntegerObject failed: %s", i, err)
            }
        }
    }

    return nil
}

func testIntegerObject(expected int64, actual object.Object) error {
    result, ok := actual.(*object.Integer)
    if !ok {
        return fmt.Errorf("object is not Integer. got=%T (%+v)", actual, actual)
    }

    if result.Value != expected {
        return fmt.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
    }

    return nil
}



