package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
)

var (
	cpuProfile = flag.String("cpuprofile", "", "write CPU profile to file")
	memProfile = flag.String("memprofile", "", "write heap profile to file")
)

func main() {
	flag.Parse()

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Println("could not create CPU profile:", err)
			os.Exit(1)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Println("could not start CPU profile:", err)
			os.Exit(1)
		}
		defer pprof.StopCPUProfile()
	}

	args := flag.Args()
	if len(args) > 0 {
		runFile(args[0])

		if *memProfile != "" {
			time.Sleep(100 * time.Millisecond)
			runtime.GC()

			f, err := os.Create(*memProfile)
			if err != nil {
				fmt.Println("could not create memory profile:", err)
				os.Exit(1)
			}
			if err := pprof.WriteHeapProfile(f); err != nil {
				fmt.Println("could not write heap profile:", err)
			}
			f.Close()
		}

		return
	}

	// Else, start REPL
	fmt.Println("Monkey Interpreter (REPL mode)")
	fmt.Println("Type code and press Enter.")
	repl.Start(os.Stdin, os.Stdout)
}

func runFile(filename string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	source := string(bytes)
	l := lexer.New(source)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, e := range p.Errors() {
			fmt.Println("Parser Error:", e)
		}
		os.Exit(1)
	}

	env := object.NewEnvironment()
	result := evaluator.Eval(program, env)

	if result != nil {
		fmt.Println(result.Inspect())
	}
}
