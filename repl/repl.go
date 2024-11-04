package repl

import (
	"bufio"
	"fmt"
	"io"
	"knightcompiler/compiler"
	"knightcompiler/lexer"
	"knightcompiler/parser"
	"knightcompiler/vm"
    "os/user"
)



const PROMPT = "%s:~$ "

func terminalUser() (string, error)  {
    currentUser, err := user.Current()
    if err != nil {
        fmt.Printf("No user was found: %v", err)
    }

    return  currentUser.Username, nil
}

func Start(in io.Reader, out io.Writer) {
    // initialize the scanner to read input from the specified io.Reader
	scanner := bufio.NewScanner(in)

    //Retrieve the current user's name for the prompt
    username, err := terminalUser()
    if err != nil {
        fmt.Fprintln(out, err)
        username = "terminal" // fallback username if retrieval fails
    }
    // infinite for loop
    // displaying the prompt and waiting for user input
	for {
		fmt.Fprintf(out, PROMPT, username)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()  // reading line of user input
        // call lexer and parser
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

        // user input compilation
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}

        // running the bytecode generated
		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}

        // displaying the results
		stackTop := machine.StackTop()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

// Formats and prints any parsing errors encountered
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some knightcompiler business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
