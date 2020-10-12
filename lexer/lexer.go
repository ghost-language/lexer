package lexer

import "github.com/ghost-language/lexer/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}

	lexer.readChar()

	return lexer
}

// readChar gives us the next character and advances our
// position in the input string.
func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition++
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.ch {
	case '=':
		if lexer.peekChar() == '=' {
			tok = lexer.newTwoCharToken(token.EQ)
		} else {
			tok = lexer.newToken(token.ASSIGN)
		}
	case '+':
		tok = lexer.newToken(token.PLUS)
	case '-':
		tok = lexer.newToken(token.MINUS)
	case '!':
		if lexer.peekChar() == '=' {
			tok = lexer.newTwoCharToken(token.NOT_EQ)
		} else {
			tok = lexer.newToken(token.BANG)
		}
	case '/':
		tok = lexer.newToken(token.SLASH)
	case '*':
		tok = lexer.newToken(token.ASTERISK)
	case '<':
		tok = lexer.newToken(token.LT)
	case '>':
		tok = lexer.newToken(token.GT)
	case ';':
		tok = lexer.newToken(token.SEMICOLON)
	case ',':
		tok = lexer.newToken(token.COMMA)
	case '(':
		tok = lexer.newToken(token.LPAREN)
	case ')':
		tok = lexer.newToken(token.RPAREN)
	case '{':
		tok = lexer.newToken(token.LBRACE)
	case '}':
		tok = lexer.newToken(token.RBRACE)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lexer.ch) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = lexer.newToken(token.ILLEGAL)
		}
	}

	lexer.readChar()

	return tok
}

// peekChar "peeks" ahead at the next character without
// incrementing the current read position. This is useful
// for determining if the next character has any meaning
// to the current one.
func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	}

	return lexer.input[lexer.readPosition]
}

// readIdentifier loops through characters until it finds
// a non alphabetical character. Once this process ends, it returns
// the alphabetical value as a whole.
func (lexer *Lexer) readIdentifier() string {
	position := lexer.position

	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

// readNumber loops through characters until it no longer
// finds a number value. Once this process ends, it returns
// the numeric value as a whole.
func (lexer *Lexer) readNumber() string {
	position := lexer.position

	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

// isLetter determines if the passed character is
// an alphabetic value.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit determines if the passed character is
// a numeric value.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// lexer.newToken takes in the desired token type and current character
// and returns a new Token instance.
func (lexer *Lexer) newToken(tokenType token.TokenType) token.Token {
	return token.Token{Type: tokenType, Literal: string(lexer.ch)}
}

func (lexer *Lexer) newTwoCharToken(tokenType token.TokenType) token.Token {
	ch := lexer.ch
	lexer.readChar()
	literal := string(ch) + string(lexer.ch)

	return token.Token{Type: tokenType, Literal: literal}
}

// skipWhitespace loops through and skips over any "whitespace"
// characters. These are completely ignored by the tokenization
// process as they have no meaning with our programs.
func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}
