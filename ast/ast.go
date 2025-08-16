package ast

import (
	"monkey/token"
)

type Node interface { // this is a node in our AST
	TokenLiteral() string
}

type Statement interface { // this is a node(statement type) in our AST
	Node
	statementNode()
}

// this is an example of embedded interfaces what it basically says is that any struct implementing
// Statment should also implement Node and should also have one more function statmentNode
// (this function is called marker function see how it is not returning anything it is bascially there to make Go's typing system distinguish between Statemnt interface and Expression interface)

type Expression interface { // this is a node(expression type) in our AST
	Node
	expressionNode()
}

// first implementation of a node
type Program struct { 
	Statements []Statement
}
// the above is the root node of our source code 

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name *Identifier
	Value Expression
}

/*
let x = 5
LetStatement{
	Token: Token{Type: LET, Literal: "let"}
	Name: Identifier{Value: "x" and also the Token}
	Value: in the above case 5 
*/

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {return i.Token.Literal}