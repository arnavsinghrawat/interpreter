# Interpreter Project

## Packages

### 1. Token
These are the tokens that are recieved from the source code using a lexer. In this package the structure and types of token are defined.

Token : {
    1. Type of Token(uses Type as a tokenType which is a string underhood) --> string
    2. The literal of the token --> string
}

example: let x = 5;
x is a token of type IDENT and literal is 'x'

There are also a few constant described for the types of token and a map to map them from literal to their type but only for few tokens like fn,let etc.

To use the the map getIdentifier function is used.

### 2. Lexer
In the token package we just define a token whereas in the lexer package we are going to define functions that read the source code and help create those tokens on the basis of the source code.

Lexer : {
    1. There is the input which is the source code (or a line of the source code) --> string
    2. position (consider this index of the input string) --> int
    3. readPosition is the next index to the position and helps us determining the type of token we have to create. --> int
    4. ch this is the character under the postion index --> byte
}

#### Function in lexer
1. peekChar() --> gives us the character at the readPosition

2. readChar() --> makes the ch equal to the character at readPostion and increments the postion and readPosition character

3. skipWhiteSpace() --> since in monkey language spaces hold not significance therefore we need to skips characters like spaces, tabs and etc.

4.  nextToken() --> we take the current l.ch under consideration and use a switch stament to create the correct type of token and return that token

5. readNumber() --> use for reading numbers(specially useful when number is bigger than one ccharacter like "9834")

6. readIdentifier() --> reads identifier

### 3. AST
After being done with the tokenization of the source code we need to need create a abtract syntax tree for the tokens that we have created

#### Interfaces in AST
1. Node: {
    1. tokenLiteral() --> gives string(helps in finding the error)
}

2. Statement: {
    1. Implements all the functions of the Node interface
    2. statementNode()
}

3. Expression: {
    1. Implements all the functions of the Node interface
    2. expressionNode()
}

#### Structs in AST
1. Program: {
    1. Statments array: {
        1. this is a basically a slice of all the struct that implement struct (the code is basically a series of statments)
    }
}

2. LetStatment:{
    1. token.Token
    2. indentifier(*Identifier)
    3. value(Expression)
    }  

3. Identifier :{
    1. token.Token
    2. Value --> this will be a string
}


//function to be added