// Code generated by goyacc -o y.go parser.y. DO NOT EDIT.

//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2

import "fmt"

//line parser.y:8
type yySymType struct {
	yys int
	val any
}

const String = 57346
const Number = 57347
const Print = 57348
const Label = 57349
const Type = 57350
const Get = 57351
const Put = 57352
const Post = 57353
const Delete = 57354
const Head = 57355
const Header = 57356
const Body = 57357
const Json = 57358
const JsonStr = 57359
const Comment = 57360
const Bool = 57361
const Write = 57362
const Append = 57363
const StringFunc = 57364
const CodeBlock = 57365
const If = 57366
const For = 57367
const Read = 57368
const Map = 57369
const Len = 57370
const ForEach = 57371
const Else = 57372

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"String",
	"Number",
	"Print",
	"Label",
	"Type",
	"Get",
	"Put",
	"Post",
	"Delete",
	"Head",
	"Header",
	"Body",
	"Json",
	"JsonStr",
	"Comment",
	"Bool",
	"Write",
	"Append",
	"StringFunc",
	"CodeBlock",
	"If",
	"For",
	"Read",
	"Map",
	"Len",
	"ForEach",
	"Else",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'('",
	"')'",
	"','",
	"'.'",
	"'['",
	"']'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 233

var yyAct = [...]uint8{
	40, 76, 144, 155, 45, 154, 170, 100, 35, 35,
	33, 44, 7, 7, 96, 97, 98, 99, 83, 78,
	149, 11, 57, 58, 59, 60, 96, 97, 98, 99,
	34, 119, 147, 74, 173, 81, 80, 172, 35, 94,
	67, 70, 73, 93, 169, 163, 33, 35, 98, 99,
	89, 84, 79, 162, 75, 33, 82, 160, 171, 35,
	90, 159, 109, 33, 111, 105, 106, 107, 108, 50,
	51, 158, 87, 35, 20, 23, 24, 21, 22, 157,
	118, 33, 56, 124, 52, 35, 123, 53, 33, 117,
	115, 26, 54, 55, 114, 32, 35, 130, 131, 132,
	133, 116, 88, 33, 104, 33, 35, 43, 103, 49,
	33, 146, 168, 137, 41, 140, 143, 96, 97, 98,
	99, 161, 135, 148, 153, 138, 129, 128, 127, 126,
	125, 152, 122, 102, 101, 63, 156, 18, 6, 19,
	20, 23, 24, 21, 22, 62, 92, 165, 91, 5,
	61, 25, 27, 86, 39, 28, 29, 26, 50, 51,
	30, 42, 38, 20, 23, 24, 21, 22, 37, 48,
	77, 56, 113, 52, 50, 51, 53, 42, 112, 85,
	26, 54, 55, 68, 139, 69, 136, 56, 65, 52,
	150, 164, 53, 151, 135, 145, 26, 54, 55, 71,
	66, 141, 95, 77, 142, 72, 55, 121, 53, 120,
	167, 166, 110, 46, 13, 13, 2, 36, 31, 64,
	134, 47, 17, 16, 15, 14, 12, 10, 9, 8,
	4, 3, 1,
}

var yyPact = [...]int16{
	131, 131, -1000, -1000, -1000, -1000, 64, -1, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 138, -1000, -1000, 126, 118,
	170, 170, 170, 170, 170, 114, 109, 99, 181, 178,
	198, -1000, 154, 12, 154, 11, -1000, 156, 65, 195,
	85, -1000, -29, -1000, -30, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 98, 97, 72, 68, 85, 85, 85,
	85, 170, 208, 170, 155, -29, -1000, -30, 149, 71,
	67, 78, 66, 57, 85, -1000, -1000, -5, -1000, 204,
	-1000, 85, -1000, -1000, 202, -1000, 95, 49, -1000, 46,
	93, 92, 91, 90, 85, 89, 170, 170, 170, 170,
	108, 179, 88, 177, 197, 108, 108, 180, 180, -6,
	86, -18, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 186,
	-36, -38, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	14, 14, -1000, -1000, -1000, 170, 42, 34, -1000, 24,
	20, 84, 16, 8, -1000, 187, -1000, 207, -1000, 206,
	75, 7, -31, 21, -1000, -1000, 85, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 0, -3, -1000, -1000,
	-1000, -1000, -1000, -1000,
}

var yyPgo = [...]uint8{
	0, 232, 216, 231, 230, 229, 228, 227, 21, 226,
	213, 225, 224, 223, 222, 114, 11, 109, 107, 4,
	0, 1, 221, 169, 220, 7, 2, 219, 217,
}

var yyR1 = [...]int8{
	0, 1, 1, 2, 2, 2, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 6, 15, 15, 15,
	5, 5, 5, 5, 5, 5, 5, 5, 9, 11,
	10, 7, 21, 21, 21, 21, 17, 17, 17, 3,
	3, 3, 3, 3, 3, 18, 18, 18, 18, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 24, 25,
	25, 26, 26, 26, 8, 8, 8, 8, 8, 16,
	16, 16, 16, 19, 19, 27, 27, 27, 28, 12,
	12, 13, 13, 13, 13, 22, 23, 23, 14, 14,
}

var yyR2 = [...]int8{
	0, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	4, 4, 3, 4, 4, 4, 4, 4, 6, 6,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 0,
	2, 2, 2, 0, 3, 3, 3, 4, 4, 3,
	3, 5, 5, 4, 4, 1, 1, 1, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 4, 3, 3,
}

var yyChk = [...]int16{
	-1000, -1, -2, -3, -4, 18, 7, -16, -5, -6,
	-7, -8, -9, -10, -11, -12, -13, -14, 6, 8,
	9, 12, 13, 10, 11, 20, 26, 21, 24, 25,
	29, -2, 31, 39, 31, 39, -28, 30, 36, 36,
	-20, -15, 7, -18, -16, -19, -10, -22, -23, -17,
	4, 5, 19, 22, 27, 28, 17, -20, -20, -20,
	-20, 36, 36, 36, -27, 7, 19, -16, 5, 7,
	-16, -23, 7, -16, -20, -8, -21, 16, 7, 40,
	-21, -20, -8, 7, 40, 23, -15, 7, 37, -16,
	-8, -17, -18, -19, -20, 7, 32, 33, 34, 35,
	-25, 36, 36, 36, 36, -25, -25, -25, -25, -20,
	4, -20, 23, 23, 23, 23, 23, 23, 23, 36,
	5, 5, 37, 37, 37, 37, 37, 37, 37, 37,
	-20, -20, -20, -20, -24, 14, 7, -16, 37, 7,
	-16, 4, 7, -16, -26, 15, -26, 38, 37, 38,
	4, 7, -16, -19, 41, 41, -20, 37, 37, 37,
	37, 37, 37, 37, 4, -21, 4, 4, 37, 37,
	37, 37, 37, 37,
}

var yyDef = [...]int8{
	0, -2, 1, 3, 4, 5, 16, 0, 6, 7,
	8, 9, 10, 11, 12, 13, 14, 15, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 2, 0, 0, 0, 0, 80, 0, 0, 0,
	59, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	17, 18, 19, 0, 0, 0, 0, 59, 59, 59,
	59, 0, 0, 0, 0, 75, 76, 77, 0, 0,
	0, 0, 0, 0, 39, 40, 41, 0, 69, 0,
	42, 43, 44, 70, 0, 78, 49, 50, 22, 52,
	0, 57, 51, 53, 0, 0, 0, 0, 0, 0,
	64, 0, 0, 0, 0, 65, 66, 63, 63, 0,
	0, 0, 79, 81, 82, 83, 84, 88, 89, 0,
	0, 0, 20, 21, 23, 24, 25, 26, 27, 31,
	45, 46, 47, 48, 60, 0, 0, 0, 85, 0,
	0, 0, 0, 0, 67, 0, 68, 0, 30, 0,
	0, 0, 0, 0, 71, 72, 58, 73, 74, 86,
	87, 36, 37, 38, 61, 62, 0, 0, 32, 33,
	34, 35, 28, 29,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	36, 37, 34, 32, 38, 33, 39, 35, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 31, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 40, 3, 41,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:66
		{
			PrintYySymDebug(yyDollar[1])
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:72
		{
			PrintYySym(yyDollar[3])
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:73
		{
			PrintGlobalYySym(yyDollar[3])
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:74
		{
			println("")
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:75
		{
			PrintYySym(yySymType{val: WalkField(yyDollar[3].val.(string))})
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:76
		{
			PrintYySym(yyDollar[3])
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:77
		{
			PrintYySym(yyDollar[3])
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:78
		{
			PrintYySym(yyDollar[3])
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:79
		{
			PrintYySym(yyDollar[3])
		}
	case 28:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:82
		{
			WriteFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 29:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:85
		{
			AppendFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:88
		{
			yyVAL = yySymType{val: ReadFile(yyDollar[3].val.(string))}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:91
		{
			PrintType(yyDollar[3])
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:94
		{
			yyVAL = yySymType{val: StrToJson(yyDollar[3].val)}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:95
		{
			yyVAL = yySymType{val: StrToJson(GlobalVarRead(yyDollar[3]))}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:96
		{
			yyVAL = yySymType{val: StrToJson(WalkField(yyDollar[3].val.(string)))}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:97
		{
			yyVAL = yySymType{val: StrToJson(yyDollar[3].val)}
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:100
		{
			yyVAL = yySymType{val: ToJsonStr(yyDollar[3].val)}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:101
		{
			yyVAL = yySymType{val: ToJsonStr(GlobalVarRead(yyDollar[3]))}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:102
		{
			yyVAL = yySymType{val: ToJsonStr(WalkField(yyDollar[3].val.(string)))}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:105
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:106
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:107
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:108
		{
			GlobalVarFieldWrite(yyDollar[1].val.(string), yyDollar[3].val)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:109
		{
			GlobalVarFieldWrite(yyDollar[1].val.(string), yyDollar[3].val)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:110
		{
			GlobalVarFieldWrite(yyDollar[1].val.(string), yyDollar[3].val)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:113
		{
			yyVAL = yySymType{val: ValAdd(yyDollar[1].val, yyDollar[3].val)}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:114
		{
			yyVAL = yySymType{val: ValMinus(yyDollar[1].val, yyDollar[3].val)}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:115
		{
			yyVAL = yySymType{val: ValMul(yyDollar[1].val, yyDollar[3].val)}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:116
		{
			yyVAL = yySymType{val: ValDiv(yyDollar[1].val, yyDollar[3].val)}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:119
		{
			yyVAL = yyDollar[1]
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:120
		{
			yyVAL = yySymType{val: GlobalVarRead(yyDollar[1])}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:121
		{
			yyVAL = yyDollar[1]
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:122
		{
			yyVAL = yySymType{val: WalkField(yyDollar[1].val.(string))}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:123
		{
			yyVAL = yyDollar[1]
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:124
		{
			yyVAL = yyDollar[1]
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:125
		{
			yyVAL = yyDollar[1]
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:126
		{
			yyVAL = yyDollar[1]
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:127
		{
			yyVAL = yyDollar[1]
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:130
		{
			yyVAL = yyDollar[2]
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:133
		{
			yyVAL = yySymType{val: nil}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:134
		{
			yyVAL = joinHeaders(yyDollar[1], yyDollar[2])
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:137
		{
			yyVAL = yySymType{val: yyDollar[2].val}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:138
		{
			yyVAL = yySymType{val: yyDollar[2].val}
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:139
		{
			yyVAL = yySymType{val: nil}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:142
		{
			yyVAL = HttpSend("GET", yyDollar[2].val.(string), yyDollar[3], yySymType{})
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:143
		{
			yyVAL = HttpSend("DELETE", yyDollar[2].val.(string), yyDollar[3], yySymType{})
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:144
		{
			yyVAL = HttpSend("HEAD", yyDollar[2].val.(string), yyDollar[3], yySymType{})
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:145
		{
			yyVAL = HttpSend("PUT", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:146
		{
			yyVAL = HttpSend("POST", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:150
		{
			yyVAL = yySymType{val: yyDollar[1].val.(string) + "." + yyDollar[3].val.(string)}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:153
		{
			yyVAL = yySymType{val: yyDollar[1].val.(string) + "." + yyDollar[3].val.(string)}
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:156
		{
			yyVAL = yySymType{val: fmt.Sprintf("%s.[%d]", yyDollar[1].val, yyDollar[4].val)}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:159
		{
			yyVAL = yySymType{val: fmt.Sprintf("%s.[%d]", yyDollar[1].val, yyDollar[4].val)}
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:164
		{
			yyVAL = yySymType{val: ToStr(GlobalVarRead(yyDollar[3]))}
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:165
		{
			yyVAL = yySymType{val: ToStr(WalkField(yyDollar[3].val.(string)))}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:168
		{
			yyVAL = yySymType{val: GlobalVarRead(yyDollar[1])}
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:169
		{
			yyVAL = yyDollar[1]
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:170
		{
			yyVAL = yySymType{val: WalkField(yyDollar[1].val.(string))}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:173
		{
			yyVAL = yyDollar[2]
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:176
		{
			yyVAL = yySymType{val: RunIfCond(yyDollar[2].val, yyDollar[3].val)}
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:177
		{
			yyVAL = yySymType{val: RunIfCond(!(yyDollar[1].val.(bool)), yyDollar[2].val)}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:180
		{
			RepeatBlock(yyDollar[2].val, yyDollar[3].val)
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:181
		{
			RepeatBlock(GlobalVarRead(yyDollar[2]), yyDollar[3].val)
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:182
		{
			RepeatBlock(WalkField(yyDollar[2].val.(string)), yyDollar[3].val)
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:183
		{
			RepeatBlock(yyDollar[2].val, yyDollar[3].val)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:186
		{
			yyVAL = yySymType{val: map[string]any{}}
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:189
		{
			yyVAL = yySymType{val: CalcLen(GlobalVarRead(yyDollar[3]))}
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:190
		{
			yyVAL = yySymType{val: CalcLen(WalkField(yyDollar[3].val.(string)))}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:193
		{
			DoForEach(GlobalVarRead(yyDollar[2]), yyDollar[3].val)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:194
		{
			DoForEach(WalkField(yyDollar[2].val.(string)), yyDollar[3].val)
		}
	}
	goto yystack /* stack new state and value */
}
