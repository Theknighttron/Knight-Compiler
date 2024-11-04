package vm

import (
	"fmt"
	"knightcompiler/ast"
	"knightcompiler/lexer"
	"knightcompiler/object"
	"knightcompiler/parser"
	"knightcompiler/compiler"
    "testing"
)

type vmTestCase struct {
    input string
    expected interface{}
}

func runVmTests(t  *testing.T, tests []vmTestCase) {
    t.Helper()

    for _, tt := range tests {
        // parse the input to AST
        program := parse(tt.input)

        // compile the program
        comp := compiler.New()
        err := comp.Compile(program)
        if err != nil {
            t.Fatalf("Compiler error: %s ", err)
        }

        // create a vm with the bytecode and run it
        vm := New(comp.Bytecode())
        err = vm.Run()      // start the vm
        if err != nil {
            t.Fatalf("VM error %s", err)
        }

        // compare the output
        stackElem := vm.StackTop()
        testExpectedObject(t, tt.expected, stackElem)
    }

}

// compare the VM's output with the expected value
func testExpectedObject(
        t *testing.T,
        expected interface{},
        actual object.Object,
) {
    t.Helper()

    // check if expected is integer then call testIntegerObject
    switch expected := expected.(type) {
    case int:
        err := testIntegerObject(int64(expected), actual)
        if err != nil {
            t.Errorf("testIntegerObject failed: %s", err)
        }
    }
}

// Convert string into AST
func parse(input string) *ast.Program {
    l := lexer.New(input)
    p := parser.New(l)
    return p.ParseProgram()
}

// check the actual result return by VM is of type object.Integer
// with it's value matching the expected integer value
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

func TestIntegerArithmetic(t *testing.T) {
        tests := []vmTestCase{
            {"1", 1},
            {"2", 2},
            {"1 + 2", 2},
        }

        runVmTests(t, tests)
}


