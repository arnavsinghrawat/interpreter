package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

// interface for a node in our AST
type Node interface {
	TokenLiteral() string // returns at actual litral of the node(from the token refering to this node)
	String() string       // return the semantic value of the node
}

// interface for a statement nodes
type Statement interface {
	Node
	statementNode()
}

// interface for an expression nodes
type Expression interface {
	Node
	expressionNode()
}

// root node of our AST
type Program struct {
	Statements []Statement
}

// the statements slice can hold all the different types of struct that implement statement

// returns the tokenLiteralFunctions of the nodes(that implement statement) in statments[]
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// returns a string that comprises of the (return of String() methods) for the each node(that implements the Statement node) in the Statements slice
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// node to represent a identifier(Implements Expression interface)
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

/*
let x = 5

let <identifier> = <expression>

	LetStatement{
		Token: Token{Type: LET, Literal: "let"}
		Name: Identifier{Value: "x" and also the Token}
		Value: in the above case 5
*/

// node for let statment (Implments Statment interface)
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal } // returns token literal in this case --> let

// return a string which give out --> "let (name of var) = (expression);"
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

// return statment parsing
/*
return 5;
return 10;
return add(15);
*/

// return <expression>;

// node for return statement (Implments Statment interface)
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// function returns the token literal which in this case is --> return
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// returns a string which in this case gives out --> "return expression;"
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// expression statements

/*
Examples:-

	Factorial(20);
	x + 4;  (given that x is assigned to value)
*/

// node for expressions in our ast(implments statment node)
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

/*
Example for x + 4
ExpressionStatement{
    Token:      token.Token{Type: IDENT, Literal: "x"}, // first token of expr
    Expression: &InfixExpression{
        Left:     &Identifier{Token: token.Token{Type: IDENT, Literal: "x"}, Value: "x"},
        Operator: "+",
        Right:    &IntegerLiteral{Token: token.Token{Type: INT, Literal: "10"}, Value: 10},
    },
}
*/

// integer literal
// for numbers like 4,5,3

// node for interger expressions in our ast
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// prefix Expression
// example := -3,-7 etc

// node for prefix expressions in our ast
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// infix parsing expression
// eg := 4 + 5, 6 * 8

// node for infix interger expressions in our ast
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}

// node for boolean expressions in our ast
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// if and else block
/*
if (<condition>) <consequence> else <alternative>
*/
// node for if expressions in our ast
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// node for bolck statement in our ast
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

//function literals
/*
example:-
fn(x, y) {
return x + y;
}
structure:-
fn <parameters> <block statement>
*/
// node for function definition(Implements Expression node)
type FunctionLiteral struct {
	Token      token.Token // would be "fn"
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

// this --> function will give a string like tokenLiteral like fn(params(separated using ',')) body
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}

// Call expression
/*
<expression>(<comma separated expressions>)
*/
// node for call expression
type CallExpression struct {
	Token     token.Token // would be "("
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

// Node for the string (implements expression node)
type StringLiteral struct {
	Token token.Token // would be "
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

// adding arrays
type ArrayLiteral struct {
	Token    token.Token // would be [
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// node for indexing of array literals (implements expression node)
type IndexExpression struct {
	Token token.Token // would be [
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

// hash map
// {<expression> : <expression>, <expression> : <expression>, ... }

// node for hash map in golang(implements Expression node)
type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

// node for while statement
type WhileStatement struct {
	Token     token.Token
	Condition Expression
	Body      *BlockStatement
}

func (we *WhileStatement) statementNode() {}

func (we *WhileStatement) TokenLiteral() string {
	return we.Token.Literal
}

func (we *WhileStatement) String() string {
	var out bytes.Buffer

	out.WriteString("while")
	out.WriteString(we.Condition.String())
	out.WriteString(" ")
	out.WriteString(we.Body.String())

	return out.String()
}
