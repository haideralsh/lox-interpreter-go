package token

import (
	"fmt"
	"os"
)

type Type string

const BEGINNING_OF_COMMENT = "//"

const (
	LeftParen    Type = "LEFT_PAREN"
	RightParen   Type = "RIGHT_PAREN"
	LeftBrace    Type = "LEFT_BRACE"
	RightBrace   Type = "RIGHT_BRACE"
	EOF          Type = "EOF"
	Comma        Type = "COMMA"
	Dot          Type = "DOT"
	Minus        Type = "MINUS"
	Plus         Type = "PLUS"
	SemiColon    Type = "SEMICOLON"
	Star         Type = "STAR"
	Error        Type = "Error"
	Equal        Type = "EQUAL"
	EqualEqual   Type = "EQUAL_EQUAL"
	Bang         Type = "BANG"
	BangEqual    Type = "BANG_EQUAL"
	Less         Type = "LESS"
	LessEqual    Type = "LESS_EQUAL"
	Greater      Type = "GREATER"
	GreaterEqual Type = "GREATER_EQUAL"
	Slash        Type = "SLASH"
)

type composite struct {
	Base Type
	Full Type
}

var Composites = []composite{
	{Equal, EqualEqual},
	{Bang, BangEqual},
	{Less, LessEqual},
	{Greater, GreaterEqual},
}

type Token struct {
	Type   Type
	Lexeme string
	Line   int
}

func (t Token) String() string {
	if t.Type == Error {
		return fmt.Sprintf("[line %d] Error: Unexpected character: %s", t.Line, t.Lexeme)
	}

	return fmt.Sprintf("%s %s null", t.Type, t.Lexeme)
}

func (t Token) Print() {
	if t.Type == Error {
		fmt.Fprintln(os.Stderr, t.String())
	} else {
		fmt.Fprintln(os.Stdout, t.String())
	}
}
