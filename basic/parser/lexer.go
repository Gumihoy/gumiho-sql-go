package parser

import (
	"strings"
	"unicode"
)

type ISQLLexer interface {
	Position() int
	Ch() rune

	Line() int
	Col() int
	Token() *Token
	SetToken(token *Token)

	PreToken() *Token
	SetPreToken(token *Token)

	Value() string
	PutChar()
	PutAndScanChar()
	ClearValue()

	Mark() *Mark
	Reset() bool

	IsSQLIdentifierStart() bool
	IsSQLIdentifierPart() bool

	SkipWhitespace()
	ScanChar()

	ScanDoubleQuota() *Kind
	ScanReverseQuota() *Kind
	ScanSingleQuota() *Kind

	ScanStartZero() *Kind
	ScanNumeric() *Kind

	ScanT() *Kind
	ScanF() *Kind

	ScanAnd() *Kind
	ScanOr() *Kind
	ScanNot() *Kind
	ScanXor() *Kind
	ScanExclamation() *Kind

	ScanPlus() *Kind
	ScanMinus() *Kind
	ScanStar() *Kind
	ScanSlash() *Kind
	ScanPercent() *Kind

	ScanGT() *Kind
	ScanEQ() *Kind
	ScanLT() *Kind

	ScanDot() *Kind
	ScanAt() *Kind
	ScanQuestion() *Kind

	ScanLeftParen() *Kind
	ScanRightParen() *Kind

	ScanLeftBracket() *Kind
	ScanRightBracket() *Kind

	ScanLeftBrace() *Kind
	ScanRightBrace() *Kind

	ScanComma() *Kind
	ScanSemi() *Kind

	// ScanIdent() *Kind
	GetKindMap() map[string]*Kind

	StringValue() string

	IsEOF() bool

	Error() string
}
type SQLLexer struct {
	content         []rune
	position, limit int
	ch              rune
	line, col       int

	valueBuff       []rune
	preToken, token *Token
	mark            *Mark
}

type Mark struct {
	position        int
	ch              rune
	line, col       int
	valueBuff       []rune
	preToken, token *Token
}

func NewLexer(sql string) *SQLLexer {
	var lexer SQLLexer
	buff := []rune(sql)
	lexer.content = buff
	lexer.limit = len(buff)
	lexer.position, lexer.line, lexer.col = 0, 1, 0
	lexer.ch = lexer.content[lexer.position]

	return &lexer
}

func (lexer *SQLLexer) Ch() rune {
	return lexer.ch
}

func (lexer *SQLLexer) Position() int {
	return lexer.position
}

func (lexer *SQLLexer) Line() int {
	return lexer.line
}

func (lexer *SQLLexer) Col() int {
	return lexer.col
}

func (lexer *SQLLexer) Token() *Token {
	return lexer.token
}

func (lexer *SQLLexer) SetToken(token *Token) {
	lexer.token = token
}

func (lexer *SQLLexer) PreToken() *Token {
	return lexer.preToken
}

func (lexer *SQLLexer) SetPreToken(token *Token) {
	lexer.preToken = token
}

func (lexer *SQLLexer) Value() string {
	return string(lexer.valueBuff)
}

func (lexer *SQLLexer) PutChar() {
	lexer.valueBuff = append(lexer.valueBuff, lexer.ch)
}

func (lexer *SQLLexer) PutCharWithCh(ch rune) {
	lexer.valueBuff = append(lexer.valueBuff, ch)
}

func (lexer *SQLLexer) PutAndScanChar() {
	lexer.valueBuff = append(lexer.valueBuff, lexer.ch)
	lexer.ScanChar()
}

func (lexer *SQLLexer) ClearValue() {
	if len(lexer.valueBuff) > 0 {
		lexer.valueBuff = lexer.valueBuff[0:0]
	}
}

func (lexer *SQLLexer) Mark() *Mark {
	mark := new(Mark)
	mark.position = lexer.position
	mark.ch = lexer.ch
	mark.line = lexer.line
	mark.col = lexer.col

	mark.preToken = lexer.preToken
	mark.token = lexer.token

	lexer.mark = mark

	return mark
}

func (lexer *SQLLexer) Reset() bool {
	return lexer.ResetWithMark(lexer.mark)
}

func (lexer *SQLLexer) ResetWithMark(mark *Mark) bool {
	if mark == nil {
		return false
	}
	lexer.position, lexer.ch = mark.position, mark.ch
	lexer.line, lexer.col = mark.line, mark.col
	lexer.preToken, lexer.token = mark.preToken, mark.token
	lexer.valueBuff = mark.valueBuff
	return true
}

func (lexer *SQLLexer) IsSQLIdentifierStart() bool {
	return lexer.IsSQLIdentifierStartWithCh(lexer.Ch())
}

func (lexer *SQLLexer) IsSQLIdentifierStartWithCh(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9') ||
		(ch == '$') ||
		(ch == '_') ||
		(ch >= 0x0080 && ch <= 0xFFFF)
}

func (lexer *SQLLexer) IsSQLIdentifierPart() bool {
	return lexer.IsSQLIdentifierPartWithCh(lexer.Ch())
}

func (lexer *SQLLexer) IsSQLIdentifierPartWithCh(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9') ||
		(ch == '$') ||
		(ch == '_') ||
		(ch >= 0x0080 && ch <= 0xFFFF)
}

func (lexer *SQLLexer) SkipWhitespace() {
	for {
		if !IsWhitespace(lexer.ch) {
			return
		}
		if lexer.ch == '\n' {
			lexer.line++
		}
		lexer.ScanChar()
	}
}

func (lexer *SQLLexer) ScanChar() {
	lexer.position++
	lexer.ch = lexer.charAt(lexer.position)
}
func (lexer *SQLLexer) unScan() {
	lexer.position--
	lexer.ch = lexer.charAt(lexer.position)
}

func (lexer *SQLLexer) SkipChar(count int)  {
	lexer.position+=count
	lexer.ch = lexer.charAt(lexer.position)
}

const END = -1

func (lexer *SQLLexer) charAt(position int) rune {
	if position >= lexer.limit {
		return END
	}
	return lexer.content[position]
}



// ""
// "xxx"" xxx"  => xxx" xxx
// "xxx" " xxx"  => xxx xxx
func (lexer *SQLLexer) ScanDoubleQuota() *Kind {
	for {
		lexer.ScanString()
		if lexer.Ch() == '"' {
			lexer.PutCharWithCh('\\')
			lexer.PutCharWithCh('"')
			continue
		}
		if IsWhitespace(lexer.Ch()) {
			lexer.ScanChar()
		}
		if lexer.Ch() == '"' || lexer.Ch() == '\'' {
			continue
		}
		break
	}

	return IDENTIFIER_DOUBLE_QUOTE
}

func (lexer *SQLLexer) ScanReverseQuota() *Kind {
	panic("UnSupport '`' error:" + lexer.Error())
}

// ''
// 'xxx' xxx'  => xxx" xxx
// 'xxx' ' xxx'  => xxx xxx
func (lexer *SQLLexer) ScanSingleQuota() *Kind {
	for {
		lexer.ScanString()
		if lexer.Ch() == '\'' {
			lexer.PutCharWithCh('\\')
			lexer.PutCharWithCh('\'')
			continue
		}
		if IsWhitespace(lexer.Ch()) {
			lexer.ScanChar()
		}
		if lexer.Ch() == '"' || lexer.Ch() == '\'' {
			continue
		}
		break
	}

	return LITERAL_STRING
}

// "string"
// 'string'
// `string`
func (lexer *SQLLexer) ScanString() {
	firstCh := lexer.ch
	lexer.ScanChar()
	for {
		if lexer.ch == '\\' {
			lexer.PutAndScanChar()
			lexer.PutAndScanChar()
		}
		if lexer.ch == firstCh {
			lexer.ScanChar()
			break
		}
		if lexer.ch == END {
			panic("")
		}
		lexer.PutAndScanChar()
	}
}

func (lexer *SQLLexer) ScanStartZero() *Kind {
	lexer.ScanChar()
	return LITERAL_INTEGER
}

/**
 * 1, .2, 3.4, -5, -6.78, +9.10
 * 1.2E3, 1.2E-3, -1.2E3, -1.2E-3
 */
func (lexer *SQLLexer) ScanNumeric() *Kind {
	kind := LITERAL_INTEGER

	if lexer.ch == '-' || lexer.ch == '+' {
		lexer.PutAndScanChar()
		// 处理空格
		lexer.SkipWhitespace()
	}

	lexer.ScanDigit()

	if lexer.ch == '.' {
		if lexer.charAt(lexer.position+1) == '.' {
			return kind
		}
		lexer.PutAndScanChar()
		lexer.ScanDigit()

		if lexer.ch == 'e' || lexer.ch == 'E' {
			lexer.PutAndScanChar()
			if lexer.ch == '-' || lexer.ch == '+' {
				lexer.PutAndScanChar()
			}
			lexer.ScanDigit()
		}
		kind = LITERAL_FLOATING_POINT
	} else if lexer.ch == 'E' || lexer.ch == 'e' {
		lexer.PutAndScanChar()
		if lexer.ch == '-' || lexer.ch == '+' {
			lexer.PutAndScanChar()
		}
		lexer.ScanDigit()
		kind = LITERAL_FLOATING_POINT
	}
	return kind
}

func IsDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func (lexer *SQLLexer) ScanDigit() {
	for {
		if IsDigit(lexer.ch) {
			lexer.PutAndScanChar()
			continue
		}
		break
	}
}

// true
func (lexer *SQLLexer) ScanT() *Kind {
	kind := LITERAL_BOOLEAN_TRUE
	if (lexer.charAt(lexer.position) == 't' || lexer.charAt(lexer.position) == 'T') &&
		(lexer.charAt(lexer.position+1) == 'r' || lexer.charAt(lexer.position+1) == 'R') &&
		(lexer.charAt(lexer.position+2) == 'u' || lexer.charAt(lexer.position+2) == 'U') &&
		(lexer.charAt(lexer.position+3) == 'e' || lexer.charAt(lexer.position+3) == 'E') &&
		!lexer.IsSQLIdentifierPartWithCh(lexer.charAt(lexer.position+4)) {

		lexer.SkipChar(4)

	} else {
		kind = ScanIdent(lexer)
	}
	return kind
}

// false
func (lexer *SQLLexer) ScanF() *Kind {
	kind := LITERAL_BOOLEAN_FALSE
	if (lexer.charAt(lexer.position) == 'f' || lexer.charAt(lexer.position) == 'F') &&
		(lexer.charAt(lexer.position+1) == 'a' || lexer.charAt(lexer.position+1) == 'A') &&
		(lexer.charAt(lexer.position+2) == 'l' || lexer.charAt(lexer.position+2) == 'L') &&
		(lexer.charAt(lexer.position+3) == 's' || lexer.charAt(lexer.position+3) == 'S') &&
		(lexer.charAt(lexer.position+4) == 'e' || lexer.charAt(lexer.position+4) == 'E') &&
		!lexer.IsSQLIdentifierPartWithCh(lexer.charAt(lexer.position+5)) {

		lexer.SkipChar(5)
	} else {
		kind = ScanIdent(lexer)
	}
	return kind
}

func (lexer *SQLLexer) ScanAnd() *Kind {
	lexer.ScanChar()
	return SYMB_BIT_AND
}

func (lexer *SQLLexer) ScanOr() *Kind {
	lexer.ScanChar()
	return SYMB_BIT_OR
}

func (lexer *SQLLexer) ScanNot() *Kind {
	lexer.ScanChar()
	return SYMB_BIT_NOT
}

// ^
func (lexer *SQLLexer) ScanXor() *Kind {
	lexer.ScanChar()
	return SYMB_BIT_XOR
}

// !
// !=
func (lexer *SQLLexer) ScanExclamation() *Kind {
	lexer.ScanChar()
	if lexer.ch == '=' {
		lexer.ScanChar()
		return SYMB_EXCLAMATION_EQUAL
	}
	return SYMB_EXCLAMATION
}

// +
// +1
func (lexer *SQLLexer) ScanPlus() *Kind {
	// 默认值
	kind := SYMB_PLUS
	nextPosition := lexer.position + 1
	next := lexer.charAt(nextPosition)

	// 处理空格
	for {
		if IsWhitespace(next) {
			if next == '\n' {
				lexer.line++
			}
			nextPosition := nextPosition + 1
			next = lexer.charAt(nextPosition)
			continue
		}
		break
	}

	if CanScanNumeric(lexer, next) {
		lexer.ScanNumeric()
	} else {
		lexer.ScanChar()
	}

	return kind
}

func (lexer *SQLLexer) ScanMinus() *Kind {
	// 默认值
	kind := SYMB_MINUS
	nextPosition := lexer.position + 1
	next := lexer.charAt(nextPosition)

	if next == '-' {
		// lexer.ScanMinusComment()
	} else {
		// 处理空格
		for {
			if IsWhitespace(next) {
				if next == '\n' {
					lexer.line++
				}
				nextPosition := nextPosition + 1
				next = lexer.charAt(nextPosition)
				continue
			}
			break
		}

		if CanScanNumeric(lexer, next) {
			kind = lexer.ScanNumeric()
		} else {
			lexer.ScanChar()
		}
	}
	return kind
}

func CanScanNumeric(lexer *SQLLexer, ch rune) bool {
	return IsDigit(ch) && (lexer.preToken == nil || (lexer.token != nil && (lexer.token.Kind == SYMB_LEFT_PAREN || lexer.token.Kind == SYMB_COMMA)))
}

// *
func (lexer *SQLLexer) ScanStar() *Kind {
	lexer.ScanChar()
	if lexer.ch == '=' {
		lexer.ScanChar()
		return SYMB_MULT_EQUAL
	}
	return SYMB_STAR
}

func (lexer *SQLLexer) ScanSlash() *Kind {
	panic("implement me")
}

func (lexer *SQLLexer) ScanPercent() *Kind {
	lexer.ScanChar()
	return SYMB_PERCENT
}

func (lexer *SQLLexer) ScanGT() *Kind {
	lexer.ScanChar()
	return SYMB_GREATER_THAN
}

// =

func (lexer *SQLLexer) ScanEQ() *Kind {
	lexer.ScanChar()
	return SYMB_EQUAL
}

//
func (lexer *SQLLexer) ScanLT() *Kind {
	lexer.ScanChar()
	return SYMB_LESS_THAN
}

func (lexer *SQLLexer) ScanDot() *Kind {
	lexer.ScanChar()
	if IsDigit(lexer.ch) {
		for {
			lexer.PutAndScanChar()
			if !IsDigit(lexer.ch) {
				break
			}
		}
		return LITERAL_FLOATING_POINT
	}
	return SYMB_DOT
}

func (lexer *SQLLexer) ScanAt() *Kind {
	lexer.ScanChar()
	return SYMB_AT
}
func (lexer *SQLLexer) ScanQuestion() *Kind {
	lexer.ScanChar()
	return SYMB_QUESTION
}

func (lexer *SQLLexer) ScanLeftParen() *Kind {
	lexer.ScanChar()
	return SYMB_LEFT_PAREN
}

func (lexer *SQLLexer) ScanRightParen() *Kind {
	lexer.ScanChar()
	return SYMB_RIGHT_PAREN
}

func (lexer *SQLLexer) ScanLeftBracket() *Kind {
	lexer.ScanChar()
	return SYMB_LERT_BRACKET
}

func (lexer *SQLLexer) ScanRightBracket() *Kind {
	lexer.ScanChar()
	return SYMB_RIGHT_BRACKET
}

func (lexer *SQLLexer) ScanLeftBrace() *Kind {
	lexer.ScanChar()
	return SYMB_LERT_BRACE
}

func (lexer *SQLLexer) ScanRightBrace() *Kind {
	lexer.ScanChar()
	return SYMB_RIGHT_BRACE
}

func (lexer *SQLLexer) ScanComma() *Kind {
	lexer.ScanChar()
	return SYMB_COMMA
}

func (lexer *SQLLexer) ScanSemi() *Kind {
	lexer.ScanChar()
	return SYMB_SEMI
}

func ScanIdent(lexer ISQLLexer) *Kind {
	lexer.PutChar()
	for {
		lexer.ScanChar()
		if !lexer.IsSQLIdentifierPart() {
			break
		}
		lexer.PutChar()
	}
	kind, ok := LookUp(lexer)
	if !ok {
		kind = IDENTIFIER
	}

	return kind
}

func LookUp(lexer ISQLLexer) (*Kind, bool) {
	kindMap := lexer.GetKindMap()
	kind, ok := kindMap[strings.ToUpper(lexer.StringValue())]
	return kind, ok
}

func (lexer *SQLLexer) GetKindMap() map[string]*Kind {
	return KindMap
}

func (lexer *SQLLexer) StringValue() string {
	return string(lexer.valueBuff)
}

func (lexer *SQLLexer) IsEOF() bool {
	return lexer.position >= lexer.limit
}

func (lexer *SQLLexer) Error() string {
	return string(lexer.content[lexer.position:])
}

func (lexer *SQLLexer) clearValueBuff() {
	lexer.valueBuff = lexer.valueBuff[:0]
}

func NextToken(lexer ISQLLexer) {
	if lexer == nil {
		panic("lexer is nil.")
	}

	lexer.ClearValue()

	var kind *Kind
	var start, end int
loop:
	for {
		lexer.SkipWhitespace()

		start = lexer.Position()
		switch lexer.Ch() {
		case '"':
			kind = lexer.ScanDoubleQuota()

		case '`':
			kind = lexer.ScanReverseQuota()

		case '\'':
			kind = lexer.ScanSingleQuota()

		case '0':
			kind = lexer.ScanStartZero()

		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			kind = lexer.ScanNumeric()

		// case 'b':
		// 	
		// case 'B':
		// 	
		// case 'x':
		// 	
		// case 'X':
		case 't', 'T':
			kind = lexer.ScanT()
		case 'f', 'F':
			kind = lexer.ScanF()

		case '&':
			kind = lexer.ScanAnd()

		case '|':
			kind = lexer.ScanOr()

		case '~':
			kind = lexer.ScanNot()

		case '^':
			kind = lexer.ScanXor()

		case '!':
			kind = lexer.ScanExclamation()

		case '+':
			kind = lexer.ScanPlus()

		case '-':
			kind = lexer.ScanMinus()

		case '*':
			kind = lexer.ScanStar()

		case '/':
			kind = lexer.ScanSlash()

		case '%':
			kind = lexer.ScanPercent()

		case '>':
			kind = lexer.ScanGT()

		case '=':
			kind = lexer.ScanEQ()

		case '<':
			kind = lexer.ScanLT()

		case '.':
			kind = lexer.ScanDot()

		case '@':
			kind = lexer.ScanAt()

		case '?', '？':
			kind = lexer.ScanQuestion()

		case '(', '（':
			kind = lexer.ScanLeftParen()

		case ')', '）':
			kind = lexer.ScanRightParen()

		case '[':
			kind = lexer.ScanLeftBracket()

		case ']':
			kind = lexer.ScanRightBracket()

		case '{':
			kind = lexer.ScanLeftBrace()

		case '}':
			kind = lexer.ScanRightBrace()

		case ',', '，':
			kind = lexer.ScanComma()

		case ';', '；':
			kind = lexer.ScanSemi()

		default:
			if lexer.IsSQLIdentifierStart() {
				kind = ScanIdent(lexer)
			} else if lexer.IsEOF() {
				kind = EOF
			} else {
				panic("TODO:" + string(lexer.Ch()))
			}
		}

		if kind == COMMENT {
			continue
		}

		break loop
	}

	lexer.SetPreToken(lexer.Token())

	end = lexer.Position() - 1
	token := NewToken(start, end, kind)
	lexer.SetToken(token)
}

func IsWhitespace(ch rune) bool {
	return unicode.IsSpace(ch)
}

var KindMap = make(map[string]*Kind)

func init() {

	KindMap["CREATE"] = CREATE

	KindMap["FROM"] = FROM

	KindMap["INSERT"] = INSERT

	KindMap["UPDATE"] = UPDATE

	KindMap["SELECT"] = SELECT

	KindMap["WHERE"] = WHERE

}
