const PUSH = "PUSH";
const ADD = "ADD";
const MINUS = "MINUS";

// The function accept an array of instruction from a program
let virtualMachine = function (program) {
  let programCounter = 0;
  let stack = [];
  let stackPointer = 0;

  while (programCounter < program.length) {
    let currentInstruction = program[programCounter];

    switch (currentInstruction) {
      case PUSH:
        stack[stackPointer] = program[programCounter + 1];
        console.log("Stack: ", stack[stackPointer]);
        stackPointer++;
        programCounter++;
        console.log("Stack Pointer: ", stackPointer);
        console.log("Program Counter: ", programCounter);
        break;

      case ADD:
        right = stack[stackPointer - 1];
        stackPointer--;
        left = stack[stackPointer - 1];
        stackPointer--;

        console.log("Stack: ", stack[stackPointer]);
        stack[stackPointer] = left + right;
        stackPointer++;
        console.log("Stack Pointer: ", stackPointer);
        console.log("Program Counter: ", programCounter);
        break;

      case MINUS:
        right = stack[stackPointer - 1];
        stackPointer--;
        left = stack[stackPointer - 1];
        stackPointer--;

        console.log("Stack: ", stack[stackPointer]);
        stack[stackPointer] = left - right;
        stackPointer++;
        console.log("Stack Pointer: ", stackPointer);
        console.log("Program Counter: ", programCounter);
        break;
    }

    programCounter++;
  }

  console.log("StackTop: ", stack[stackPointer - 1]);
};

let program = [PUSH, 3, PUSH, 4, ADD, PUSH, 5, MINUS];
virtualMachine(program);
