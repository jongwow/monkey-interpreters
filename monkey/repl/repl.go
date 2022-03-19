package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/jongwow/monkey/evaluator"
	"github.com/jongwow/monkey/lexer"
	"github.com/jongwow/monkey/object"
	"github.com/jongwow/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func StartByLine(done chan bool, in chan string, out chan<- string) {
	env := object.NewEnvironment()
	var buffer bytes.Buffer
	fmt.Println("Executed")
	for {
		select {
		case <-done:
			return
		case s := <-in:
			fmt.Println("Recv: ", s)
			l := lexer.New(s)
			p := parser.New(l)
			pgm := p.ParseProgram()
			if len(p.Errors()) != 0 {
				out <- printParserErrorsString(p.Errors())
				continue
			}
			evaluated := evaluator.Eval(pgm, env)
			if evaluated != nil {
				buffer.WriteString(evaluated.Inspect())
				buffer.WriteString("\n")
				ret := buffer.String()
				fmt.Println("evaluated: ", ret)
				out <- ret
				buffer.Reset()
			}
		}
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

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func printParserErrorsString(errors []string) string {
	var buffer bytes.Buffer

	buffer.WriteString(MONKEY_FACE)
	buffer.WriteString("Woops! We ran into some monkey business here!\n")
	buffer.WriteString(" parser errors:\n")
	for _, msg := range errors {
		buffer.WriteString("\t" + msg + "\n")
	}
	return buffer.String()
}
