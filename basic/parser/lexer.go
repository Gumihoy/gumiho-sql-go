package parser

import "gumihoy.com/sql/basic/parser/kind"

type ILexer interface {
	Ch() rune

	GetMark() int
	SetMark(mark int)

	Pos() int

	Line() int
	Col() int
	Token() *Token
	SetToken(token *Token)

	GetValue() string
	AppendValue()

	IsSQLIdentifierStart() bool
	IsSQLIdentifierPart() bool

	ScanWhitespace()
	ScanRune()

	ScanDoubleQuota() kind.Kind
	ScanReverseQuota() kind.Kind
	ScanSingleQuota() kind.Kind

	ScanStartZero() kind.Kind
	ScanNumeric() kind.Kind
	ScanIdent() kind.Kind

	IsEOF() bool
}

type Lexer struct {
	content         []rune
	position, limit int
	mark, line, col int
	ch              rune

	valueBuff []rune
	token     *Token
}

func NewLexer(sql string) *Lexer {
	var lexer Lexer
	buff := []rune(sql)
	lexer.content = buff
	lexer.limit = len(buff)
	lexer.position, lexer.line, lexer.col = 0, 1, 0
	lexer.ch = lexer.content[lexer.position]

	return &lexer
}

func (lexer *Lexer) Ch() rune {
	return lexer.ch
}

func (lexer *Lexer) GetMark() int {
	return lexer.mark
}

func (lexer *Lexer) SetMark(mark int) {
	lexer.mark = mark
}

func (lexer *Lexer) Pos() int {
	return lexer.position
}

func (lexer *Lexer) Line() int {
	return lexer.line
}

func (lexer *Lexer) Col() int {
	return lexer.col
}

func (lexer *Lexer) Token() *Token {
	return lexer.token
}

func (lexer *Lexer) SetToken(token *Token) {
	lexer.token = token
}

func (lexer *Lexer) GetValue() string {
	return string(lexer.valueBuff)
}

func (lexer *Lexer) AppendValue() {
	lexer.valueBuff = append(lexer.valueBuff, lexer.ch)
}

func (lexer *Lexer) IsSQLIdentifierStart() bool {
	return (lexer.ch >= 'a' && lexer.ch <= 'z') ||
		(lexer.ch >= 'A' && lexer.ch <= 'Z') ||
		(lexer.ch == '_')
}

func (lexer *Lexer) IsSQLIdentifierPart() bool {
	return (lexer.ch >= 'a' && lexer.ch <= 'z') ||
		(lexer.ch >= 'A' && lexer.ch <= 'Z') ||
		(lexer.ch >= 0 && lexer.ch <= 9) ||
		(lexer.ch == '_')
}

func (lexer *Lexer) ScanWhitespace() {
	for {
		if !IsWhitespace(lexer.ch) {
			return
		}
		if lexer.ch == '\n' {
			lexer.line++
		}
		lexer.position++
	}
}

func (lexer *Lexer) ScanRune() {
	lexer.ch = lexer.content[lexer.position]
	lexer.position++
}

func (lexer *Lexer) ScanDoubleQuota() kind.Kind {
	return kind.IDENTIFIER_DOUBLE_QUOTE
}
func (lexer *Lexer) ScanReverseQuota() kind.Kind {
	panic(NewError(lexer.line, lexer.col, " UnSupport '`'"))
}
func (lexer *Lexer) ScanSingleQuota() kind.Kind {
	return kind.LITERAL_TEXT
}

func (lexer *Lexer) ScanStartZero() kind.Kind {
	return kind.LITERAL_INTEGER
}
func (lexer *Lexer) ScanNumeric() kind.Kind {
	return kind.LITERAL_INTEGER
}

func (lexer *Lexer) ScanIdent() kind.Kind {
	lexer.AppendValue()
	for {
		lexer.ScanRune()
		if !lexer.IsSQLIdentifierPart() {
			break
		}
		lexer.AppendValue()
	}
	return kind.LITERAL_INTEGER
}

func (lexer *Lexer) IsEOF() bool {
	return lexer.position >= lexer.limit
}

func (lexer *Lexer) clearValueBuff() {
	lexer.valueBuff = lexer.valueBuff[:0]
}
func (lexer *Lexer) Mark() {
	lexer.mark = lexer.position
}
func (lexer *Lexer) Reset() {
	lexer.position = lexer.mark
}

func NextToken(lexer ILexer) {
	var k kind.Kind
loop:
	for {
		lexer.ScanWhitespace()
		switch lexer.Ch() {
		case '"':
			k = lexer.ScanDoubleQuota()
			break loop
		case '`':
			k = lexer.ScanReverseQuota()
			break loop
		case '\'':
			k = lexer.ScanSingleQuota()
			break loop
		case '0':
			k = lexer.ScanStartZero()
			break loop
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			k = lexer.ScanNumeric()
			break loop
		default:
			if lexer.IsSQLIdentifierStart() {
				k = lexer.ScanIdent()
				break loop
			}

			if lexer.IsEOF() {
				k = kind.EOF
				break loop
			}
			break loop
		}
	}
	token := NewToken(lexer.GetMark(), 0, k)
	lexer.SetToken(token)
}

func IsWhitespace(ch rune) bool {
	return ch == ' ' ||
		ch == '\n' ||
		ch == '\t'
}
