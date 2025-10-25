# Custom Programming language interpreter

A complete interpreter for a custom programming language, written in Go. This interpreter implements a full programming language including lexing, parsing, and evaluation of code. This interpreter is made with the help of the book Writing An Interpreter In Go-Thorsten Ball.

## Overview

The Monkey programming language is a small, interpreted language that supports:
- Variable bindings
- Integers and booleans
- Arithmetic expressions
- Built-in functions
- First-class and higher-order functions
- Closures
- String data structure
- Array data structure
- Hash data structure

## Project Structure

### Core Components

1. **Token (`/token`)** 
   - Defines all tokens used in the language
   - Implements token types and structures
   - Handles keyword identification

2. **Lexer (`/lexer`)**
   - Performs lexical analysis
   - Converts source code into tokens
   - Handles character-by-character processing
   - Manages whitespace and comments

3. **Abstract Syntax Tree (`/ast`)**
   - Defines the structure for parsed code
   - Implements nodes for:
     - Expressions
     - Statements
     - Identifiers
     - Literals

4. **Parser (`/parser`)**
   - Converts tokens into an AST
   - Implements Pratt parsing
   - Handles:
     - Prefix expressions
     - Infix expressions
     - Grouped expressions
     - IF expressions
     - Function literals

5. **Object System (`/object`)**
   - Defines internal object system
   - Implements:
     - Integer objects
     - Boolean objects
     - String objects
     - Array objects
     - Hash objects
     - Function objects
     - Return values
     - Error handling

6. **Evaluator (`/evaluator`)**
   - Executes the parsed AST
   - Implements:
     - Expression evaluation
     - Built-in functions
     - Error handling
     - Environment management

7. **REPL (`/repl`)**
   - Provides interactive shell
   - Read-Eval-Print Loop implementation

## Language Features

### 1. Variable Bindings
```monkey
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
```

### 2. Functions and Closures
```monkey
let fibonacci = fn(x) {
    if (x < 2) { return x; }
    return fibonacci(x - 1) + fibonacci(x - 2);
};
```

### 3. Arrays and Built-in Functions
```monkey
let myArray = [1, 2, 3, 4, 5];
let length = len(myArray);
let first = first(myArray);
let last = last(myArray);
```

### 4. Hash Maps
```monkey
let prices = {"apple": 5, "banana": 3};
prices["apple"]
```

### 5. Built-in Functions
- `len()`: Returns length of strings and arrays
- `first()`: Returns first element of array
- `last()`: Returns last element of array
- `rest()`: Returns array without first element
- `push()`: Adds element to array
- `puts()`: Prints arguments to console

## Running the Interpreter

1. Clone the repository:
```bash
git clone <repository-url>
cd monkey
```

2. Build the interpreter:
```bash
go build
```

3. Run the REPL:
```bash
go run main.go
```

## Testing

Run the test suite:
```bash
go test ./...
```

## Implementation Details

### Lexical Analysis
- Character-by-character scanning
- Token identification
- Handling of identifiers, numbers, and operators

### Parsing
- Recursive descent parsing
- Pratt parsing for expressions
- Operator precedence handling
- Error recovery and reporting

### Evaluation
- Tree-walking interpreter
- Environment-based scope handling
- Built-in function implementation
- Error handling and propagation

## Contributing

Feel free to contribute by:
1. Forking the repository
2. Creating your feature branch
3. Committing your changes
4. Pushing to the branch
5. Creating a Pull Request

## License

This project is open source and available under the [MIT License].