package token

type TokenType string

// A Token is a small data structure that are fed into the parser. Like the chunks of the expression
type Token struct {
	Type    TokenType
	Literal string
}

// Token Types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	COMMENT = "#"
	NULL    = "NULL"

	// Identifiers + literals
	IDENT  = "IDENT"  // add, foobar, x, y, ...
	INT    = "INT"    // 1343456
	STRING = "STRING" // "foobar"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	MODULO   = "%"
	CARROT   = "^"

	PLUSASSIGN     = "+="
	MINUSASSIGN    = "-="
	ASTERISKASSIGN = "*="
	SLASHASSIGN    = "/="

	// Logical Comparison
	LT     = "<"
	GT     = ">"
	LTE    = "<="
	GTE    = ">="
	EQ     = "=="
	NOT_EQ = "!="
	AND    = "&&"
	OR     = "||"

	// Bitwise operators
	BITWISEOR     = "|"
	BITWISEAND    = "&"
	BITWISEANDNOT = "&^"
	SHIFTLEFT     = "<<"
	SHIFTRIGHT    = ">>"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	QMARK     = "?"
	DOT       = "."

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	FOR      = "FOR"
	IN       = "IN"
	RETURN   = "RETURN"
	IMPORT   = "IMPORT"
	FROM     = "FROM"
	AS       = "AS"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"in":     IN,
	"return": RETURN,
	"import": IMPORT,
	"from":   FROM,
	"as":     AS,
	"null":   NULL,
}

// LookupIdent finds the equivalent constant for a given identifier
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
