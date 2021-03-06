package lexer

import (
	"pascal_in_go/token"
	"unicode"
)

//ReservedKey hold the reserved key(内置保留字符)
var ReservedKey = map[string]token.Token{
	"BEGIN":     token.Token{Type: "BEGIN", Literal: "BEGIN"},
	"END":       token.Token{Type: "END", Literal: "END"},
	"VAR":       token.Token{Type: "VAR", Literal: "VAR"},
	"DIV":       token.Token{Type: "DIV", Literal: "DIV"},
	"INTEGER":   token.Token{Type: "INTEGER", Literal: "INTEGER"},
	"REAL":      token.Token{Type: "REAL", Literal: "REAL"},
	"PROGRAM":   token.Token{Type: "PROGRAM", Literal: "PROGRAM"},
	"PROCEDURE": token.Token{Type: "PROCEDURE", Literal: "PROCEDURE"},
}

type Lexer struct {
	Text    string `json:"text"`
	Pos     int    `json:"pos"`
	CurChar byte   `json:"curChar"`
}

func NewLexer(text string) Lexer {
	return Lexer{Text: text, Pos: 0, CurChar: text[0]}
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token
	for lexer.CurChar != 0 {
		if unicode.IsSpace(rune(lexer.CurChar)) {
			lexer.skipWhiteSpace()
			continue
		}

		if lexer.isalpha() {
			val := lexer.letter()
			tok = getIdentifier(val)
			return tok
		}
		if lexer.isnum() {
			tok = lexer.number()
			return tok
		}
		if lexer.CurChar == '+' {
			tok.Type = token.PLUS
			tok.Literal = "+"
			lexer.advance()
			return tok
		}

		if lexer.CurChar == '-' {
			tok.Type = token.MINUS
			tok.Literal = "-"
			lexer.advance()
			return tok
		}

		if lexer.CurChar == '*' {
			tok.Type = token.MUL
			tok.Literal = "*"
			lexer.advance()
			return tok
		}

		if lexer.CurChar == '/' {
			tok.Type = token.DIV
			tok.Literal = "/"
			lexer.advance()
			return tok
		}

		if lexer.CurChar == '(' {
			tok.Type = token.LPAREN
			tok.Literal = "("
			lexer.advance()
			return tok
		}

		if lexer.CurChar == ')' {
			tok.Type = token.RPAREN
			tok.Literal = ")"
			lexer.advance()
			return tok
		}

		if lexer.CurChar == '.' {
			tok.Type = token.DOT
			tok.Literal = "."
			lexer.advance()
			return tok
		}

		if lexer.CurChar == ':' && lexer.peek() == '=' {
			tok.Type = token.ASSIGN
			tok.Literal = ":="
			lexer.advance()
			lexer.advance()
			return tok
		}
		if lexer.CurChar == ';' {
			lexer.advance()
			tok.Type = token.SEMI
			tok.Literal = ";"
			return tok
		}

		if lexer.CurChar == ':' {
			lexer.advance()
			tok.Type = token.COLON
			tok.Literal = ":"
			return tok
		}

		if lexer.CurChar == ',' {
			lexer.advance()
			tok.Type = token.COMMA
			tok.Literal = ","
			return tok
		}

		if lexer.CurChar == 0 {
			tok.Type = token.EOF
			tok.Literal = ""
			lexer.advance()
			return tok
		}
	}
	return newToken(token.ILLEGAL, lexer.CurChar)
}

func (lexer *Lexer) isalpha() bool {
	ch := lexer.CurChar
	if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
		return true
	}
	return false
}

func getIdentifier(val string) token.Token {
	tok, ok := ReservedKey[val]
	if ok {
		return tok
	}
	return token.Token{
		Type:    token.ID,
		Literal: val,
	}
}

//
func (lexer *Lexer) isnum() bool {
	ch := lexer.CurChar
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

func (lexer *Lexer) integer() string {
	result := ""
	for lexer.CurChar != 0 && unicode.IsDigit(rune(lexer.CurChar)) {
		result += string(lexer.CurChar)
		lexer.advance()
	}
	return result
}

func (lexer *Lexer) number() token.Token {
	tok := token.Token{}
	result := ""
	for lexer.CurChar != 0 && unicode.IsDigit(rune(lexer.CurChar)) {
		result += string(lexer.CurChar)
		lexer.advance()
	}

	tok.Type = token.INTEGER
	tok.Literal = result

	if lexer.CurChar == '.' {
		result += "."
		lexer.advance()
		for lexer.CurChar != 0 && lexer.isnum() {
			result += string(lexer.CurChar)
			lexer.advance()
		}
		tok.Literal = result
		tok.Type = token.REAL
	}
	return tok
}

func (lexer *Lexer) letter() string {
	result := ""
	result += string(lexer.CurChar)
	lexer.advance()

	for lexer.CurChar != 0 && (lexer.isalpha() || lexer.isnum()) {
		result += string(lexer.CurChar)
		lexer.advance()
	}

	return result
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.CurChar == ' ' || lexer.CurChar == '\t' || lexer.CurChar == '\r' || lexer.CurChar == '\n' {
		lexer.advance()
	}
}

func (lexer *Lexer) advance() {
	lexer.Pos++
	if lexer.Pos > len(lexer.Text)-1 {
		lexer.CurChar = 0
	} else {
		lexer.CurChar = lexer.Text[lexer.Pos]
	}

}

func (lexer *Lexer) peek() byte {
	pos := lexer.Pos + 1
	var curChar byte = 0
	if pos < len(lexer.Text)-1 {
		curChar = lexer.Text[pos]
	}

	return curChar

}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
