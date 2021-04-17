package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"strings"
	"unicode"
)

type ISQLLexer interface {
	Position() int
	Ch() rune

	Limit() int

	Line() int
	Col() int
	Token() *SQLToken
	SetToken(token *SQLToken)

	PreToken() *SQLToken
	// SetPreToken(token *SQLToken)

	Value() string
	PutChar()
	PutAndScanChar()
	ClearValue()

	Mark() *Mark
	Reset() bool
	ResetWithMark(mark *Mark) bool

	IsSQLIdentifierStart() bool
	IsSQLIdentifierPart() bool

	SkipWhitespace()
	ScanChar()
	ScanMinusCommentRest()
	ScanMultiLineCommentRest()
	ScanSharpCommentRest()

	ScanDoubleQuota() *Kind
	ScanReverseQuota() *Kind
	ScanSingleQuota() *Kind

	ScanStartZero() *Kind
	ScanNumeric() *Kind

	ScanStartX(child ISQLLexer) *Kind

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

	ScanColon(child ISQLLexer) *Kind

	ScanSharp(child ISQLLexer) *Kind

	// ScanIdent() *Kind
	GetKindMap() map[string]*Kind

	StringValue() string

	Comments() []ast.ISQLComment
	AddComment(comment ast.ISQLComment)
	ClearComments()

	IsEOF() bool

	UseNearErrorMsg() string
	UnSupport() string
}
type SQLLexer struct {
	content         []rune
	position, limit int
	ch              rune
	line, col       int

	valueBuff       []rune
	preToken, token *SQLToken
	mark            *Mark

	comments []ast.ISQLComment
}

type Mark struct {
	position        int
	ch              rune
	line, col       int
	valueBuff       []rune
	preToken, token *SQLToken
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

func (lexer *SQLLexer) Limit() int {
	return lexer.limit
}

func (lexer *SQLLexer) Line() int {
	return lexer.line
}

func (lexer *SQLLexer) Col() int {
	return lexer.col
}

func (lexer *SQLLexer) Token() *SQLToken {
	return lexer.token
}

func (lexer *SQLLexer) SetToken(token *SQLToken) {
	lexer.preToken = lexer.token
	lexer.token = token
}

func (lexer *SQLLexer) PreToken() *SQLToken {
	return lexer.preToken
}

// func (lexer *SQLLexer) SetPreToken(token *SQLToken) {
// 	lexer.preToken = token
// }

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

	mark.position, mark.ch = lexer.position, lexer.ch
	mark.line, mark.col = lexer.line, lexer.col

	mark.preToken, mark.token = lexer.preToken, lexer.token

	mark.valueBuff = make([]rune, len(lexer.valueBuff))
	copy(mark.valueBuff, lexer.valueBuff)

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
		(ch >= 0x0080 && ch <= 0xFFFF && ch != '（' && ch != '）' && ch != '，')
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
		(ch >= 0x0080 && ch <= 0xFFFF && ch != '（' && ch != '）' && ch != '，')
}

func (lexer *SQLLexer) SkipWhitespace() {
	for {
		if !IsWhitespace(lexer.ch) {
			return
		}
		lexer.ScanChar()
	}
}

func (lexer *SQLLexer) ScanChar() {
	if lexer.ch == '\n' {
		lexer.line++
		lexer.col = 0
	} else {
		lexer.col++
	}
	lexer.position++
	lexer.ch = lexer.charAt(lexer.position)
}
func (lexer *SQLLexer) unScan() {
	lexer.position--
	lexer.ch = lexer.charAt(lexer.position)
}

func (lexer *SQLLexer) SkipChar(count int) {
	lexer.position += count
	lexer.ch = lexer.charAt(lexer.position)
}

const END = -1

func (lexer *SQLLexer) charAt(position int) rune {
	if position >= lexer.limit {
		return END
	}
	return lexer.content[position]
}

/**
 * -- comment
 */
func (lexer *SQLLexer) ScanMinusCommentRest() {
	lexer.ClearValue()
	lexer.SkipWhitespace()
	for {
		lexer.PutAndScanChar()
		if lexer.ch == '\n' || lexer.ch == END {
			break
		}
	}
	comment := ast.NewMinusCommentWithComment(lexer.StringValue())
	lexer.AddComment(comment)
}

// /* */
func (lexer *SQLLexer) ScanMultiLineCommentRest() {
	lexer.ClearValue()

	for {
		if lexer.Ch() == '*' && lexer.charAt(lexer.position+1) == '/' {
			lexer.ScanChar()
			lexer.ScanChar()
			break
		}
		if lexer.ch == END {
			break
		}
		lexer.PutAndScanChar()
	}
	comment := ast.NewMultiLineCommentWithComment(lexer.StringValue())
	lexer.AddComment(comment)
}

/**
 * # comment
 */
func (lexer *SQLLexer) ScanSharpCommentRest() {
	lexer.ClearValue()
	lexer.SkipWhitespace()
	for {
		lexer.PutAndScanChar()
		if lexer.ch == '\n' || lexer.ch == END {
			break
		}
	}
	comment := ast.NewSharpCommentWithComment(lexer.StringValue())
	lexer.AddComment(comment)
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
	panic("UnSupport '`' error:" + lexer.UseNearErrorMsg())
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

/**
 * 0、0.1
 * 0x01AF
 */
func (lexer *SQLLexer) ScanStartZero() *Kind {

	if lexer.charAt(lexer.position+1) == '.' {
		lexer.PutAndScanChar()
		lexer.ScanChar()

		// 0.
		if !IsDigit(lexer.ch) {
			return LITERAL_INTEGER
		}

		lexer.PutCharWithCh('.')

		// 0.1xxx
		lexer.ScanDigit()

		return LITERAL_FLOATING_POINT

	} else if lexer.ch == 'x' {

		lexer.ScanChar()
		lexer.ScanHexadecimal()

		return LITERAL_HEXADECIMAL_0X

	} else if lexer.ch == 'X' {

		panic("")

	}

	lexer.PutAndScanChar()

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

/**
 * X'01AF'、X'01af'、x'01AF'、x'01af'
 * https://dev.mysql.com/doc/refman/8.0/en/hexadecimal-literals.html
 */
func (lexer *SQLLexer) ScanStartX(child ISQLLexer) *Kind {

	if lexer.charAt(lexer.position+1) == '\'' {
		lexer.ScanChar()

		lexer.ScanHexadecimal()

		return LITERAL_HEXADECIMAL_X
	}

	return ScanIdent(child)
}

func IsHexadecimal(ch rune) bool {
	return (ch >= '0' && ch <= '9') ||
		(ch >= 'a' && ch <= 'f') ||
		(ch >= 'A' && ch <= 'F')
}

/**
 * '01AF'
 * 01AF
 */
func (lexer *SQLLexer) ScanHexadecimal() {

	isQuote := lexer.ch == '\''
	if isQuote {
		lexer.ScanChar()
	}

	for {
		if !IsHexadecimal(lexer.ch) {
			break
		}
		lexer.PutAndScanChar()
	}

	if isQuote {
		if lexer.ch == '\'' {
			lexer.ScanChar()

		} else {
			panic("")
		}
	}

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
		lexer.ScanChar()
		lexer.ScanChar()
		kind = COMMENT_MINUS
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

/*
 * /
 * /*
 */
func (lexer *SQLLexer) ScanSlash() *Kind {
	kind := SYMB_SLASH

	lexer.ScanChar()

	if lexer.Ch() == '*' {
		lexer.ScanChar()
		kind = COMMENT_MULTI_LINE
		lexer.ScanMultiLineCommentRest()
	}

	return kind
}

func (lexer *SQLLexer) ScanPercent() *Kind {
	lexer.ScanChar()
	return SYMB_PERCENT
}

// >
// >>、>=
func (lexer *SQLLexer) ScanGT() *Kind {
	lexer.ScanChar()

	if lexer.ch == '>' {
		lexer.ScanChar()
		return SYMB_GREATER_THAN_GREATER_THAN

	} else if lexer.ch == '=' {
		lexer.ScanChar()
		return SYMB_GREATER_THAN_EQUAL
	}

	return SYMB_GREATER_THAN
}

// =
func (lexer *SQLLexer) ScanEQ() *Kind {
	lexer.ScanChar()
	return SYMB_EQUAL
}

// <
// <<、<>、<=、<=>
func (lexer *SQLLexer) ScanLT() *Kind {
	lexer.ScanChar()

	if lexer.ch == '<' {
		lexer.ScanChar()
		return SYMB_LESS_THAN_LESS_THAN

	} else if lexer.ch == '>' {
		lexer.ScanChar()
		return SYMB_LESS_THAN_GREATER_THAN

	} else if lexer.ch == '=' {

		lexer.ScanChar()

		if lexer.ch == '>' {
			lexer.ScanChar()
			return SYMB_LESS_THAN_EQUAL_GREATER_THAN
		}

		return SYMB_LESS_THAN_EQUAL
	}

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
	panic(lexer.UnSupport())
}

func (lexer *SQLLexer) ScanRightBrace() *Kind {
	panic(lexer.UnSupport())
}

func (lexer *SQLLexer) ScanComma() *Kind {
	lexer.ScanChar()
	return SYMB_COMMA
}

func (lexer *SQLLexer) ScanSemi() *Kind {
	lexer.ScanChar()
	return SYMB_SEMI
}

func (lexer *SQLLexer) ScanColon(child ISQLLexer) *Kind {
	panic(lexer.UnSupport())
}

func (lexer *SQLLexer) ScanSharp(child ISQLLexer) *Kind {
	panic(lexer.UnSupport())
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

func (lexer *SQLLexer) Comments() []ast.ISQLComment {
	return lexer.comments
}
func (lexer *SQLLexer) AddComment(comment ast.ISQLComment) {
	lexer.comments = append(lexer.comments, comment)
}
func (lexer *SQLLexer) ClearComments() {
	lexer.comments = ClearComments(lexer.comments)
}

func ClearComments(comments []ast.ISQLComment) []ast.ISQLComment {
	return comments[:0]
}

func (lexer *SQLLexer) IsEOF() bool {
	return lexer.position >= lexer.limit
}

func (lexer *SQLLexer) UseNearErrorMsg() string {
	return string(lexer.content[0:lexer.position])
}

func (lexer *SQLLexer) UnSupport() string {
	if lexer.token.Kind == IDENTIFIER {
		return "TODO: " + lexer.token.UnSupport() + " " + lexer.StringValue() + ", error:\"" + lexer.UseNearErrorMsg() + "\""
	} else {
		return "TODO: " + lexer.token.UnSupport() + ", error:\"" + lexer.UseNearErrorMsg() + "\""
	}

}

func (lexer *SQLLexer) clearValueBuff() {
	lexer.valueBuff = lexer.valueBuff[:0]
}

func NextToken(lexer ISQLLexer) {
	if lexer == nil {
		panic("lexer is nil.")
	}

	var kind *Kind
	var start, end int
	var line, col int
loop:
	for {
		lexer.ClearValue()
		lexer.SkipWhitespace()

		start = lexer.Position()
		line = lexer.Line()
		col = lexer.Col()

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
		case 'x', 'X':
			kind = lexer.ScanStartX(lexer)

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

		case ':':
			kind = lexer.ScanColon(lexer)

		case '#':
			kind = lexer.ScanSharp(lexer)

		default:
			if lexer.IsSQLIdentifierStart() {
				kind = ScanIdent(lexer)
			} else if lexer.IsEOF() {
				kind = EOF
			} else {
				panic("TODO:" + string(lexer.Ch()))
			}
		}

		break loop
	}

	end = lexer.Position() - 1
	token := NewToken(kind, start, end, line, col)
	lexer.SetToken(token)
}

func IsWhitespace(ch rune) bool {
	return unicode.IsSpace(ch)
}

var KindMap = make(map[string]*Kind)

func init() {

	// KindMap["ACCESSIBLE"] = ACCESSIBLE
	// KindMap["ADD"] = ADD
	KindMap["ALL"] = ALL
	KindMap["ALTER"] = ALTER
	// KindMap["ANALYZE"] = ANALYZE
	KindMap["AND"] = AND
	KindMap["AS"] = AS
	KindMap["ASC"] = ASC
	// KindMap["ASENSITIVE"] = ASENSITIVE

	// KindMap["BEFORE"] = BEFORE
	// KindMap["BETWEEN"] = BETWEEN
	// KindMap["BIGINT"] = BIGINT
	// KindMap["BINARY"] = BINARY
	// KindMap["BLOB"] = BLOB
	// KindMap["BOTH"] = BOTH
	KindMap["BY"] = BY

	// KindMap["CALL"] = CALL
	// KindMap["CASCADE"] = CASCADE
	// KindMap["CASE"] = CASE
	// KindMap["CHANGE"] = CHANGE
	// KindMap["CHAR"] = CHAR
	// KindMap["CHARACTER"] = CHARACTER
	// KindMap["CHECK"] = CHECK
	// KindMap["COLLATE"] = COLLATE
	// KindMap["COLUMN"] = COLUMN
	// KindMap["CONDITION"] = CONDITION
	// KindMap["CONSTRAINT"] = CONSTRAINT
	// KindMap["CONTINUE"] = CONTINUE
	// KindMap["CONVERT"] = CONVERT
	// KindMap["CREATE"] = CREATE
	// KindMap["CROSS"] = CROSS
	// KindMap["CUBE"] = CUBE
	// KindMap["CUME_DIST"] = CUME_DIST
	// KindMap["CURRENT_DATE"] = CURRENT_DATE
	// KindMap["CURRENT_TIME"] = CURRENT_TIME
	// KindMap["CURRENT_TIMESTAMP"] = CURRENT_TIMESTAMP
	// KindMap["CURRENT_USER"] = CURRENT_USER
	// KindMap["CURSOR"] = CURSOR

	// KindMap["DATABASE"] = DATABASE
	// KindMap["DATABASES"] = DATABASES
	// KindMap["DAY_HOUR"] = DAY_HOUR
	// KindMap["DAY_MICROSECOND"] = DAY_MICROSECOND
	// KindMap["DAY_MINUTE"] = DAY_MINUTE
	// KindMap["DAY_SECOND"] = DAY_SECOND
	// KindMap["DEC"] = DEC
	// KindMap["DECIMAL"] = DECIMAL
	// KindMap["DECLARE"] = DECLARE
	// KindMap["DEFAULT"] = DEFAULT
	// KindMap["DELAYED"] = DELAYED
	// KindMap["DELETE"] = DELETE
	// KindMap["DENSE_RANK"] = DENSE_RANK
	// KindMap["DESC"] = DESC
	// KindMap["DESCRIBE"] = DESCRIBE
	// KindMap["DETERMINISTIC"] = DETERMINISTIC
	// KindMap["DISTINCT"] = DISTINCT
	// KindMap["DISTINCTROW"] = DISTINCTROW
	// KindMap["DIV"] = DIV
	// KindMap["DOUBLE"] = DOUBLE
	// KindMap["DROP"] = DROP
	// KindMap["DUAL"] = DUAL

	// KindMap["EACH"] = EACH
	// KindMap["ELSE"] = ELSE
	// KindMap["ELSEIF"] = ELSEIF
	// KindMap["EMPTY"] = EMPTY
	// KindMap["ENCLOSED"] = ENCLOSED
	// KindMap["ESCAPED"] = ESCAPED
	// KindMap["EXCEPT"] = EXCEPT
	// KindMap["EXISTS"] = EXISTS
	// KindMap["EXIT"] = EXIT
	// KindMap["EXPLAIN"] = EXPLAIN

	KindMap["FALSE"] = FALSE
	// KindMap["FETCH"] = FETCH
	// KindMap["FIRST_VALUE"] = FIRST_VALUE
	// KindMap["FLOAT"] = FLOAT
	// KindMap["FLOAT4"] = FLOAT4
	// KindMap["FLOAT8"] = FLOAT8
	// KindMap["FOR"] = FOR
	// KindMap["FORCE"] = FORCE
	// KindMap["FOREIGN"] = FOREIGN
	KindMap["FROM"] = FROM
	// KindMap["FULLTEXT"] = FULLTEXT
	// KindMap["FUNCTION"] = FUNCTION
	// KindMap["GENERATED"] = GENERATED
	// KindMap["GET"] = GET
	// KindMap["GRANT"] = GRANT
	KindMap["GROUP"] = GROUP
	// KindMap["GROUPING"] = GROUPING
	// KindMap["GROUPS"] = GROUPS
	// KindMap["HAVING"] = HAVING
	// KindMap["HIGH_PRIORITY"] = HIGH_PRIORITY
	// KindMap["HOUR_MICROSECOND"] = HOUR_MICROSECOND
	// KindMap["HOUR_MINUTE"] = HOUR_MINUTE

	// HOUR_SECOND
	// IF
	// IGNORE

	KindMap["IN"] = IN
	KindMap["INDEX"] = INDEX
	// INFILE
	// INNER
	// INOUT
	// INSENSITIVE
	// INSERT
	// INT
	// INT1
	// INT2
	// INT3
	// INT4
	// INT8
	// INTEGER
	KindMap["INTERVAL"] = INTERVAL
	// INTO
	// IO_AFTER_GTIDS
	// IO_BEFORE_GTIDS
	// IS
	// ITERATE
	KindMap["JOIN"] = JOIN
	// JSON_TABLE
	// KEY
	// KEYS
	// KILL
	// LAG
	// LAST_VALUE
	// LATERAL
	// LEAD
	// LEADING
	// LEAVE
	// LEFT
	// LIKE

	KindMap["LIMIT"] = LIMIT
	// LINEAR
	// LINES
	// LOAD
	// LOCALTIME
	// LOCALTIMESTAMP
	// LOCK
	// LONG
	// LONGBLOB
	// LONGTEXT
	// LOOP
	// LOW_PRIORITY
	// MASTER_BIND
	// MASTER_SSL_VERIFY_SERVER_CERT
	// MATCH
	// MAXVALUE
	// MEDIUMBLOB
	// MEDIUMINT
	// MEDIUMTEXT
	// MIDDLEINT
	// MINUTE_MICROSECOND
	// MINUTE_SECOND
	// MOD
	// MODIFIES
	// NATURAL
	// NOT
	// NO_WRITE_TO_BINLOG
	// NTH_VALUE
	// NTILE
	// NULL
	// NUMERIC
	// OF
	KindMap["ON"] = ON
	// OPTIMIZE
	// OPTIMIZER_COSTS
	// OPTION
	// OPTIONALLY
	KindMap["OR"] = OR
	KindMap["ORDER"] = ORDER
	// OUT
	// OUTER
	// OUTFILE
	// OVER
	KindMap["PARTITION"] = PARTITION
	// PERCENT_RANK
	// PRECISION
	// PRIMARY
	// PROCEDURE
	// PURGE
	// RANGE
	// RANK
	// READ
	// READS
	// READ_WRITE
	// REAL
	// RECURSIVE
	// REFERENCES
	// REGEXP
	// RELEASE
	// RENAME
	// REPEAT
	// REPLACE
	// REQUIRE
	// RESIGNAL
	// RESTRICT
	// RETURN
	// REVOKE
	// RIGHT
	// RLIKE
	// ROW
	// ROWS
	// ROW_NUMBER
	// SCHEMA
	// SCHEMAS
	// SECOND_MICROSECOND

	KindMap["SELECT"] = SELECT

	// SENSITIVE
	// SEPARATOR
	KindMap["SET"] = SET
	// SHOW
	// SIGNAL
	// SMALLINT
	// SPATIAL
	// SPECIFIC
	// SQL
	// SQLEXCEPTION
	// SQLSTATE
	// SQLWARNING
	// SQL_BIG_RESULT
	// SQL_CALC_FOUND_ROWS
	// SQL_SMALL_RESULT
	// SSL
	// STARTING
	// STORED
	// STRAIGHT_JOIN
	// SYSTEM
	// TABLE
	// TERMINATED
	// THEN
	// TINYBLOB
	// TINYINT
	// TINYTEXT
	// TO
	// TRAILING
	// TRIGGER
	KindMap["TRUE"] = TRUE
	// UNDO
	// UNION
	KindMap["UNION"] = UNION
	// UNIQUE
	// UNLOCK
	// UNSIGNED
	// UPDATE
	// USAGE
	// USE
	// USING
	// UTC_DATE
	// UTC_TIME
	// UTC_TIMESTAMP
	// VALUES
	// VARBINARY
	// VARCHAR
	// VARCHARACTER
	// VARYING
	// VIRTUAL

	// KindMap["WHEN"] = WHEN
	KindMap["WHERE"] = WHERE
	// KindMap["WHILE"] = WHILE
	// KindMap["WINDOW"] = WINDOW
	KindMap["WITH"] = WITH
	// KindMap["WRITE"] = WRITE
	// KindMap["XOR"] = XOR
	// KindMap["YEAR_MONTH"] = YEAR_MONTH
	// KindMap["ZEROFILL"] = ZEROFILL

}
