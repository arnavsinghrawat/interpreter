package token

type TokenType string

type Token struct {
	Type    TokenType // name of the token
	Literal string // actual literal the token
}

const (
	//two letter keywords
	EQ     = "=="
	NOT_EQ = "!="

	//Special types
	ILLEGAL = "ILLEGAL" // To signify a illegal token
	EOF     = "EOF"     // To signify "end of line"

	//Identifirers+ literals
	IDENT = "IDENT"
	INT   = "INT"
	STRING = "STRING"

	//Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACKET = "["
	RBRACKET = "]"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func GetIdentfierType(ident string) TokenType {
	tokentype, ok := keywords[ident]
	if ok {
		return tokentype
	}
	return IDENT
}
