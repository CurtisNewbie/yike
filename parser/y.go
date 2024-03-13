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
	-1, 7,
	29, 58,
	30, 58,
	31, 58,
	32, 58,
	-2, 18,
	-1, 8,
	29, 61,
	30, 61,
	31, 61,
	32, 61,
	-2, 12,
	-1, 12,
	29, 59,
	30, 59,
	31, 59,
	32, 59,
	-2, 10,
	-1, 15,
	29, 63,
	30, 63,
	31, 63,
	32, 63,
	-2, 14,
	-1, 19,
	29, 57,
	30, 57,
	31, 57,
	32, 57,
	-2, 19,
	-1, 20,
	29, 56,
	30, 56,
	31, 56,
	32, 56,
	-2, 20,
	-1, 21,
	29, 60,
	30, 60,
	31, 60,
	32, 60,
	-2, 21,
}

const yyPrivate = 57344

const yyLast = 239

var yyAct = [...]uint8{
	68, 75, 8, 8, 101, 69, 6, 35, 139, 162,
	161, 19, 20, 13, 88, 100, 25, 26, 27, 28,
	29, 43, 41, 85, 78, 79, 21, 152, 151, 37,
	24, 122, 150, 31, 38, 60, 63, 73, 12, 12,
	176, 89, 43, 81, 43, 175, 90, 41, 83, 148,
	92, 87, 94, 86, 67, 80, 84, 173, 91, 43,
	147, 109, 107, 116, 114, 146, 124, 103, 104, 105,
	106, 66, 172, 82, 41, 182, 181, 96, 97, 98,
	99, 154, 93, 43, 71, 70, 153, 72, 41, 25,
	26, 27, 28, 29, 180, 179, 77, 78, 132, 74,
	43, 131, 37, 41, 178, 138, 31, 38, 138, 138,
	138, 138, 142, 143, 144, 145, 19, 20, 22, 7,
	23, 25, 26, 27, 28, 29, 177, 121, 157, 160,
	5, 21, 30, 32, 37, 174, 33, 34, 31, 38,
	41, 164, 119, 117, 42, 40, 127, 46, 47, 48,
	49, 129, 43, 41, 126, 43, 41, 171, 149, 137,
	136, 135, 134, 133, 130, 125, 48, 49, 71, 70,
	65, 72, 64, 57, 56, 55, 45, 44, 120, 118,
	102, 110, 111, 74, 115, 58, 37, 163, 102, 140,
	31, 38, 110, 111, 78, 108, 112, 59, 158, 77,
	155, 159, 61, 156, 62, 78, 123, 112, 95, 128,
	170, 169, 168, 167, 166, 165, 141, 113, 54, 53,
	52, 51, 50, 76, 15, 15, 2, 36, 39, 18,
	17, 16, 14, 11, 10, 9, 4, 3, 1,
}

var yyPact = [...]int16{
	112, 112, -1000, -1000, -1000, -1000, -1000, 117, 116, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 144, 143, 118, 218, 217, 216, 215, 214,
	142, 141, 140, 178, 197, -1000, -1000, 139, 137, -1000,
	80, 18, 80, 16, 7, 201, 164, 164, 164, 164,
	166, 166, 166, 166, 166, 188, 213, 177, 120, 156,
	119, 155, 104, 8, 199, 131, 118, -1000, -1000, -1000,
	-1000, -1000, -14, -1000, -1000, -15, -1000, 121, 113, -1000,
	204, -1000, 118, -1000, -1000, -1000, 146, 130, 67, -1000,
	64, 129, 128, 127, 126, 125, 135, 135, -1000, -1000,
	174, -1000, 212, 174, 174, 174, 174, 30, 25, 14,
	-1000, -1000, -1000, 124, -3, -7, -8, -1000, -1000, -1000,
	-1000, -1000, -1000, 52, 47, -1000, 196, 194, -28, -29,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	183, -1000, -1000, -1000, -1000, -1000, 211, 210, 209, -1000,
	208, 207, 206, -1000, -1000, 123, 38, 23, 101, 11,
	6, -1000, -1000, -1000, -1000, 92, 70, 61, 60, 42,
	41, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000,
}

var yyPgo = [...]uint8{
	0, 238, 226, 237, 236, 6, 235, 234, 233, 37,
	13, 1, 232, 223, 231, 230, 229, 5, 7, 0,
	30, 227, 4, 15, 8,
}

var yyR1 = [...]int8{
	0, 1, 1, 2, 2, 2, 2, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 7, 5,
	5, 5, 6, 6, 6, 6, 6, 6, 6, 6,
	12, 12, 12, 14, 14, 14, 13, 8, 19, 19,
	19, 17, 17, 17, 3, 3, 3, 3, 3, 3,
	3, 3, 9, 9, 9, 9, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 22, 23, 23, 23, 24,
	24, 24, 10, 10, 10, 10, 10, 11, 11, 11,
	11, 18, 18, 15, 15, 15, 16, 16, 16, 21,
}

var yyR2 = [...]int8{
	0, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 4, 4, 3, 4, 4, 4, 4, 4,
	6, 6, 6, 6, 6, 6, 4, 4, 4, 4,
	4, 4, 4, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 1, 2, 0, 2,
	2, 0, 4, 4, 4, 4, 4, 3, 3, 5,
	5, 4, 4, 3, 3, 3, 3, 3, 3, 3,
}

var yyChk = [...]int16{
	-1000, -1, -2, -3, -4, 18, -5, 7, -11, -6,
	-7, -8, -9, -10, -12, -13, -14, -15, -16, 4,
	5, 19, 6, 8, -20, 9, 10, 11, 12, 13,
	20, 26, 21, 24, 25, -18, -21, 22, 27, -2,
	28, 36, 28, 36, 33, 33, 29, 30, 31, 32,
	4, 4, 4, 4, 4, 33, 33, 33, 7, 19,
	-11, 5, 7, -11, 33, 33, -20, -10, -19, -17,
	5, 4, 7, -9, 19, -11, -13, 16, 17, 7,
	37, -19, -20, -17, -10, 7, 37, -5, 7, 34,
	-11, -10, -17, -9, -18, 7, -20, -20, -20, -20,
	-23, -22, 14, -23, -23, -23, -23, -5, 7, -17,
	4, 5, 19, 4, -5, 7, -17, 23, 23, 23,
	23, 23, 23, 7, -11, 34, 33, 33, 5, 5,
	34, 34, 34, 34, 34, 34, 34, 34, -22, -24,
	15, 4, -24, -24, -24, -24, 35, 35, 35, 34,
	35, 35, 35, 34, 34, 4, 7, -11, 4, 7,
	-11, 38, 38, 4, -19, 4, 4, 4, 4, 4,
	4, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34,
}

var yyDef = [...]int8{
	0, -2, 1, 3, 4, 5, 6, -2, -2, 7,
	8, 9, -2, 11, 13, -2, 15, 16, 17, -2,
	-2, -2, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 64, 0, 0, 2,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	68, 68, 68, 68, 68, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 44, 45, 46, 47,
	56, 57, 58, 59, 60, 61, 63, 0, 0, 77,
	0, 48, 49, 50, 51, 78, 0, 0, 58, 24,
	61, 0, 0, 59, 62, 0, 52, 53, 54, 55,
	71, 66, 0, 71, 71, 71, 71, 0, 0, 0,
	19, 20, 21, 0, 0, 0, 0, 83, 84, 85,
	86, 87, 88, 0, 0, 89, 0, 0, 0, 0,
	22, 23, 25, 26, 27, 28, 29, 37, 67, 72,
	0, 65, 73, 74, 75, 76, 0, 0, 0, 36,
	0, 0, 0, 81, 82, 0, 0, 0, 0, 0,
	0, 79, 80, 69, 70, 0, 0, 0, 0, 0,
	0, 38, 39, 40, 41, 42, 43, 30, 31, 32,
	33, 34, 35,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	33, 34, 31, 29, 35, 30, 36, 32, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 28, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 37, 3, 38,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27,
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

	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:65
		{
			PrintYySymDebug(yyDollar[1])
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:71
		{
			PrintYySym(yyDollar[3])
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:72
		{
			PrintGlobalYySym(yyDollar[3])
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:73
		{
			println("")
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:74
		{
			PrintYySym(yySymType{val: WalkField(yyDollar[3].val.(string))})
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:75
		{
			PrintYySym(yyDollar[3])
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:76
		{
			PrintYySym(yyDollar[3])
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:77
		{
			PrintYySym(yyDollar[3])
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:78
		{
			PrintYySym(yyDollar[3])
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:81
		{
			WriteFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:82
		{
			WriteFile(GlobalVarRead(yyDollar[3]), yyDollar[5].val.(string))
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:83
		{
			WriteFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 33:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:86
		{
			AppendFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 34:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:87
		{
			AppendFile(GlobalVarRead(yyDollar[3]), yyDollar[5].val.(string))
		}
	case 35:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:88
		{
			AppendFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:91
		{
			yyVAL = yySymType{val: ReadFile(yyDollar[3].val.(string))}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:94
		{
			PrintType(yyDollar[3])
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:97
		{
			yyVAL = yySymType{val: StrToMap(yyDollar[3].val)}
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:98
		{
			yyVAL = yySymType{val: StrToMap(GlobalVarRead(yyDollar[3]))}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:99
		{
			yyVAL = yySymType{val: StrToMap(WalkField(yyDollar[3].val.(string)))}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:102
		{
			yyVAL = yySymType{val: ToJsonStr(yyDollar[3].val)}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:103
		{
			yyVAL = yySymType{val: ToJsonStr(GlobalVarRead(yyDollar[3]))}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:104
		{
			yyVAL = yySymType{val: ToJsonStr(WalkField(yyDollar[3].val.(string)))}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:107
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:108
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:109
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:110
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:111
		{
			GlobalVarFieldWrite(yyDollar[1].val.(string), yyDollar[3].val)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:112
		{
			GlobalVarFieldWrite(yyDollar[1].val.(string), yyDollar[3].val)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:113
		{
			GlobalVarFieldWrite(yyDollar[1].val.(string), yyDollar[3].val)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:114
		{
			GlobalVarFieldWrite(yyDollar[1].val.(string), yyDollar[3].val)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:117
		{
			yyVAL = yySymType{val: ValAdd(yyDollar[1].val, yyDollar[3].val)}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:118
		{
			yyVAL = yySymType{val: ValMinus(yyDollar[1].val, yyDollar[3].val)}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:119
		{
			yyVAL = yySymType{val: ValMul(yyDollar[1].val, yyDollar[3].val)}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:120
		{
			yyVAL = yySymType{val: ValDiv(yyDollar[1].val, yyDollar[3].val)}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:123
		{
			yyVAL = yyDollar[1]
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:124
		{
			yyVAL = yyDollar[1]
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:125
		{
			yyVAL = yySymType{val: GlobalVarRead(yyDollar[1])}
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:126
		{
			yyVAL = yyDollar[1]
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:127
		{
			yyVAL = yyDollar[1]
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:128
		{
			yyVAL = yySymType{val: WalkField(yyDollar[1].val.(string))}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:129
		{
			yyVAL = yyDollar[1]
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:130
		{
			yyVAL = yyDollar[1]
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:131
		{
			yyVAL = yyDollar[1]
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:134
		{
			yyVAL = yyDollar[2]
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:137
		{
			yyVAL = yyDollar[1]
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:138
		{
			yyVAL = joinHeaders(yyDollar[1], yyDollar[2])
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:139
		{
			yyVAL = yySymType{val: nil}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:142
		{
			yyVAL = yySymType{val: yyDollar[2].val}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:143
		{
			yyVAL = yySymType{val: yyDollar[2].val}
		}
	case 71:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:144
		{
			yyVAL = yySymType{val: nil}
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:147
		{
			yyVAL = HttpSend("GET", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:148
		{
			yyVAL = HttpSend("PUT", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:149
		{
			yyVAL = HttpSend("POST", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:150
		{
			yyVAL = HttpSend("DELETE", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:151
		{
			yyVAL = HttpSend("HEAD", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:155
		{
			yyVAL = yySymType{val: yyDollar[1].val.(string) + "." + yyDollar[3].val.(string)}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:158
		{
			yyVAL = yySymType{val: yyDollar[1].val.(string) + "." + yyDollar[3].val.(string)}
		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:161
		{
			yyVAL = yySymType{val: fmt.Sprintf("%s.[%d]", yyDollar[1].val, yyDollar[4].val)}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:164
		{
			yyVAL = yySymType{val: fmt.Sprintf("%s.[%d]", yyDollar[1].val, yyDollar[4].val)}
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:169
		{
			yyVAL = yySymType{val: ToStr(GlobalVarRead(yyDollar[3]))}
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:170
		{
			yyVAL = yySymType{val: ToStr(WalkField(yyDollar[3].val.(string)))}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:173
		{
			RunIfCond(GlobalVarRead(yyDollar[2]), yyDollar[3].val)
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:174
		{
			RunIfCond(yyDollar[2].val, yyDollar[3].val)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:175
		{
			RunIfCond(WalkField(yyDollar[2].val.(string)), yyDollar[3].val)
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:178
		{
			RepeatBlock(yyDollar[2].val, yyDollar[3].val)
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:179
		{
			RepeatBlock(GlobalVarRead(yyDollar[2]), yyDollar[3].val)
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:180
		{
			RepeatBlock(WalkField(yyDollar[2].val.(string)), yyDollar[3].val)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:183
		{
			yyVAL = yySymType{val: map[string]any{}}
		}
	}
	goto yystack /* stack new state and value */
}
