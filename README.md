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

### Bytecode

Virtual Machine execute bytecode, bytecode is made up of instructions that tell the machine what to do.
It's called bytecode because the opcodes contained in each instructions are one byte in size,

- An "opcode" is the "operator" part of an instruction sometimes also called "op".

> Opcode is the part of machine language instruction that specifies the operation to be performed.
> It tells the processor what kind of operation needs to be executed, in assembly or machine code

- Operands are the values or data in which an instruction (such as opcode) operates.

> They are typically the inputs or parameters required by the operations specified by the opcode.
> In machine code or assembly language, operands can be:
>
> - Constants, such as numbers or characters.
> - Registers, which are small storage locations inside the CPU.
> - Memory Addresses, pointing to data stored elsewhere in memory.
> - Variables, that hold data needed for the operations.

Little endian -> means that the least significant byte of the original data comes first and is stored in the lowest memory address.
Big endian -> is the opposite whereby the most significant byte comes first

    1.Objective
    Compile and execute below knight expression:
        1 + 2
    - take the knight expression 1 +2
    - tokenize and parse it using our existing lexer, token and parser packages
    - take the resulting AST, whose node are defined in our ast package
    - pass it to the newly built compiler, which compile it to bytecode
    - take the bytecode and hand it over to the also newly-built virtual machine which will execute it
    - make sure that the virtual machine turned it into 3.

    | Lexer | --> | Parser | --> | Compiler | --> | Virtual Machine |

    In term of data structure:

    | String | --> | Tokens | --> | AST | --> | Bytecode | --> | Objects |

Type Assertion -> is a way to extract underlying value of an interface and convert it to a specific type.
Type Casting/Conversion -> refers to converting a value from one type to another.

    | Lexer | --> | Parser | --> | Compiler | --> | Virtual Machine |
    [-----------Compile Time----------------]     [-- Run Time------]

```
const (
    OpConstant Opcode = iota  -  0
    OpAdd                     -  1
    OpSubtract                -  2
    OpMultiply                -  3
)
```

The `const` keyword is used to define constants.
The `iota` keyword is a special identifier that simplifies the process of defining incrementing constants

Emit is compiler-speak for "generate" and "output" which translate to generate an instruction and add it to the result, either by printing it,
writing it to a file or by adding it to a collection in memory.

```
REPL(Read-Eval-Print-Loop)
It is an interactive programming environment that allows users to enter commands or expressions,
which are then processed and evaluated in real time.

Read:
The REPL reads the input from the user.
This input can be a single expression, a command, or a block of code.

Eval:
After reading the input, the REPL evaluates or executes the input.
This involves interpreting the code and performing the specified operations.

Print:
Once the evaluation is complete, the REPL prints the result of the evaluation back to the user.
This feedback loop allows users to see the output of their commands immediately.

Loop:
The process repeats in a loop,
allowing users to continuously enter new commands and see their results without restarting the session.
```
