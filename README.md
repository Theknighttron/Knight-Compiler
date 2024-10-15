# Compilers

| CODE | -> | COMPILER | -> | VIRTUAL MACHINE | -> | OUTPUT |

Compiler and virtual machines are ideas, patterns just like interpreter or web server's are idea that can have multiple implementation
ranging from tiny to massive.

Compiler take in source code in one language and generate source code in another one.

        [ Source Code ]
            |
            | -> (Lexer and Parser)
            |
        [   AST ]
            |
            | -> (Optimizer)
            |
        [   Internal Representation ]
            |
            | -> (Code Generator)
            |
        [   Machine Code ]

AST (Abstract Syntax Tree)
is the tree representation of the source code of a computer program that conveys the structure of the source code.

### Compilation Process

- First the source code is tokenized and parsed by lexer and parser. Then the source code is turned from text into an AST
- Then the component called optimizer translate the AST into another Internal Representation (IR)
  IR might be another Syntax tree, or binary format, or even textual format the reason is that IR might lend itself better to optimization
  and translation into the target language than AST would.
- Finally the code generator also called backend, generates code in the target language. This is where the compilation happens.
  Here is where the code hit's the file system. And after that we can execute and see the computer performs what we instructed it to
  in the original source code.

### Virtual machines

These are program that emulated a computer, including a disk drive, hard drive, graphics card etc...

### How Computer Works

- Von Neumann Architecture by John Von Neumann in 1945

                            Input Device
                    _________________________________
                                |
                                |
                    _________________________________
                    |   Central Processing Unit     |
                    |           |                   |
                    |   [                           |
                    |       Control Unit &          |
                    |    Arithmetic / Logic Unit    |
                    |                           ]   |
                    |           |                   |
                    |        Memory  Unit           |
                    _________________________________
                                |
                                |
                    _________________________________
                            Output Device

- As soon as the computer in turned on the CPU:

  1. Fetches an instruction from memory: The program counter tell the CPU where in the memory it can find the next instruction.
  2. Decodes the instruction: To identify which operation should be executed
  3. Executes the instruction: This mean either to modify the content of it's registers, or transferring data from the register to memory
     or moving data around in memory or generating output or reading input...

and then go to step 1 again....

Program Counter
is the part of the CPU that keeps tract of where to fetch the next instruction.

Computer Memory is segmented into "words"
Word is the smallest addressable region of the memory, it's the base unit when accessing the memory
The size of the word varies and depends on the CPU type among other things. However 32 and 64 bit word sizes are the standards.

### Memory

Stack
It is the region in memory where data is managed in a last-in-first-out (LIFO)
The data in it grows and shrinks, you push elements on the stack and later pop them off. The stack is used to implement "call stack".

Call stack
is a stack data structure that store's information about the active subroutines of a computer program.

Return Address
a value that is used to indicate where a particular function should return control after it finishes executing.

Stack Pointer
is a small register that stores the memory address of the last element added to the stack or in some cases the first available address in the stack.

### What is a Virtual Machines

A virtual machine is a computer built with software, it's software entity that mimic how a computer works.
It can be anything a function, a struct, an object, or even a whole program.

What matter's is what it does:

- A virtual machine has a run loop that goes through the fetch-decode-execute cycle.
- It has a program counter: it fetches instructions, decodes and execute them.
- It has a stack
- It has call stack and sometimes even registers all built in one software
