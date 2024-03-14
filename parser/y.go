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
	31, 59,
	32, 59,
	33, 59,
	34, 59,
	-2, 19,
	-1, 8,
	31, 62,
	32, 62,
	33, 62,
	34, 62,
	-2, 12,
	-1, 12,
	31, 60,
	32, 60,
	33, 60,
	34, 60,
	-2, 10,
	-1, 15,
	31, 64,
	32, 64,
	33, 64,
	34, 64,
	-2, 14,
	-1, 20,
	31, 58,
	32, 58,
	33, 58,
	34, 58,
	-2, 20,
	-1, 21,
	31, 57,
	32, 57,
	33, 57,
	34, 57,
	-2, 21,
	-1, 22,
	31, 61,
	32, 61,
	33, 61,
	34, 61,
	-2, 22,
	-1, 80,
	31, 67,
	32, 67,
	33, 67,
	34, 67,
	-2, 48,
	-1, 93,
	31, 67,
	32, 67,
	33, 67,
	34, 67,
	-2, 51,
}

const yyPrivate = 57344

const yyLast = 265

var yyAct = [...]uint8{
	79, 86, 8, 8, 111, 6, 40, 181, 156, 180,
	37, 13, 95, 110, 84, 12, 12, 192, 89, 49,
	191, 176, 47, 49, 175, 173, 47, 49, 172, 171,
	47, 49, 25, 170, 49, 47, 66, 69, 72, 149,
	148, 49, 47, 169, 96, 47, 48, 168, 135, 91,
	90, 134, 100, 80, 49, 93, 97, 102, 78, 167,
	94, 104, 101, 49, 165, 103, 47, 117, 119, 124,
	126, 113, 114, 115, 116, 137, 164, 140, 143, 77,
	163, 92, 198, 197, 132, 106, 107, 108, 109, 20,
	21, 23, 7, 24, 26, 27, 28, 29, 30, 49,
	46, 131, 44, 5, 22, 31, 33, 41, 47, 34,
	35, 32, 42, 43, 36, 155, 47, 196, 155, 155,
	155, 155, 159, 160, 161, 162, 129, 20, 21, 127,
	98, 195, 26, 27, 28, 29, 30, 194, 193, 144,
	44, 49, 22, 190, 47, 41, 179, 174, 166, 32,
	42, 43, 52, 53, 54, 55, 154, 39, 183, 99,
	82, 81, 153, 83, 152, 26, 27, 28, 29, 30,
	151, 150, 88, 44, 147, 85, 138, 76, 41, 54,
	55, 133, 32, 42, 43, 82, 81, 75, 83, 74,
	73, 130, 63, 70, 62, 61, 51, 50, 44, 67,
	85, 68, 128, 41, 64, 112, 182, 32, 42, 43,
	120, 121, 146, 125, 120, 121, 65, 118, 88, 112,
	157, 177, 43, 44, 178, 122, 139, 44, 141, 122,
	136, 142, 105, 71, 145, 189, 188, 187, 186, 185,
	184, 158, 123, 60, 59, 58, 57, 56, 87, 15,
	15, 2, 38, 45, 19, 18, 17, 16, 14, 11,
	10, 9, 4, 3, 1,
}

var yyPact = [...]int16{
	85, 85, -1000, -1000, -1000, -1000, -1000, 70, 16, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 162, 161, 121, 243, 242, 241, 240,
	239, 160, 159, 157, 197, 194, 226, -1000, -1000, -1000,
	-1000, 155, 154, 152, 142, -1000, 156, 11, 156, 5,
	123, 225, 181, 181, 181, 181, 191, 191, 191, 191,
	191, 210, 238, 206, 106, 179, 103, 168, 78, 61,
	158, 28, 25, 223, 140, 219, 224, 121, -1000, -1000,
	-1000, -1000, -1000, 7, -1000, -1000, -4, -1000, 104, -1000,
	229, -1000, 121, -1000, -1000, -1000, 207, 138, 4, -1000,
	3, 135, 134, 128, 126, 120, 146, 146, -1000, -1000,
	205, -1000, 237, 205, 205, 205, 205, 43, 39, 27,
	-1000, -1000, -1000, 112, 22, 10, 6, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -3, -7, -1000, -8,
	-11, 111, -12, -15, 217, -31, -33, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 202, -1000, -1000,
	-1000, -1000, -1000, 236, 235, 234, -1000, 233, 232, 231,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 107, -16, -19,
	-1000, -1000, -1000, -1000, 102, 101, 95, 81, 47, 46,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
}

var yyPgo = [...]int16{
	0, 264, 251, 263, 262, 5, 261, 260, 259, 14,
	11, 1, 258, 248, 257, 256, 255, 254, 6, 10,
	0, 32, 252, 157, 4, 13, 8,
}

var yyR1 = [...]int8{
	0, 1, 1, 2, 2, 2, 2, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 7,
	5, 5, 5, 6, 6, 6, 6, 6, 6, 6,
	6, 12, 12, 12, 14, 14, 14, 13, 8, 20,
	20, 20, 18, 18, 18, 3, 3, 3, 3, 3,
	3, 3, 3, 9, 9, 9, 9, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 24, 25,
	25, 25, 26, 26, 26, 10, 10, 10, 10, 10,
	11, 11, 11, 11, 19, 19, 15, 15, 15, 16,
	16, 16, 16, 22, 23, 23, 17, 17, 17,
}

var yyR2 = [...]int8{
	0, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 4, 4, 3, 4, 4, 4, 4,
	4, 6, 6, 6, 6, 6, 6, 4, 4, 4,
	4, 4, 4, 4, 4, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	2, 0, 2, 2, 0, 4, 4, 4, 4, 4,
	3, 3, 5, 5, 4, 4, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 4, 0, 3, 3,
}

var yyChk = [...]int16{
	-1000, -1, -2, -3, -4, 18, -5, 7, -11, -6,
	-7, -8, -9, -10, -12, -13, -14, -15, -16, -17,
	4, 5, 19, 6, 8, -21, 9, 10, 11, 12,
	13, 20, 26, 21, 24, 25, 29, -19, -22, -23,
	-18, 22, 27, 28, 17, -2, 30, 38, 30, 38,
	35, 35, 31, 32, 33, 34, 4, 4, 4, 4,
	4, 35, 35, 35, 7, 19, -11, 5, 7, -11,
	-23, 7, -11, 35, 35, 35, 35, -21, -10, -20,
	-18, 5, 4, 7, -9, 19, -11, -13, 16, 7,
	39, -20, -21, -18, -10, 7, 39, -5, 7, 36,
	-11, -10, -18, -9, -19, 7, -21, -21, -21, -21,
	-25, -24, 14, -25, -25, -25, -25, -5, 7, -18,
	4, 5, 19, 4, -5, 7, -18, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 7, -11, 36, 7,
	-11, 4, 7, -11, 35, 5, 5, 36, 36, 36,
	36, 36, 36, 36, 36, -24, -26, 15, 4, -26,
	-26, -26, -26, 37, 37, 37, 36, 37, 37, 37,
	36, 36, 36, 36, 36, 36, 36, 4, 7, -11,
	40, 40, 4, -20, 4, 4, 4, 4, 4, 4,
	36, 36, 36, 36, 36, 36, 36, 36, 36,
}

var yyDef = [...]int8{
	96, -2, 1, 3, 4, 5, 6, -2, -2, 7,
	8, 9, -2, 11, 13, -2, 15, 16, 17, 18,
	-2, -2, -2, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 63, 65, 66,
	67, 0, 0, 0, 0, 2, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 71, 71, 71, 71,
	71, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 45, 46, 47,
	-2, 57, 58, 59, 60, 61, 62, 64, 0, 80,
	0, 49, 50, -2, 52, 81, 0, 0, 59, 25,
	62, 0, 67, 60, 63, 0, 53, 54, 55, 56,
	74, 69, 0, 74, 74, 74, 74, 0, 0, 0,
	20, 21, 22, 0, 0, 0, 0, 86, 87, 88,
	89, 90, 91, 92, 97, 98, 0, 0, 93, 0,
	0, 0, 0, 0, 0, 0, 0, 23, 24, 26,
	27, 28, 29, 30, 38, 70, 75, 0, 68, 76,
	77, 78, 79, 0, 0, 0, 37, 0, 0, 0,
	84, 85, 94, 95, 42, 43, 44, 0, 0, 0,
	82, 83, 72, 73, 0, 0, 0, 0, 0, 0,
	39, 40, 41, 31, 32, 33, 34, 35, 36,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	35, 36, 33, 31, 37, 32, 38, 34, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 30, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 39, 3, 40,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29,
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

	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:68
		{
			PrintYySymDebug(yyDollar[1])
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:74
		{
			PrintYySym(yyDollar[3])
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:75
		{
			PrintGlobalYySym(yyDollar[3])
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:76
		{
			println("")
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:77
		{
			PrintYySym(yySymType{val: WalkField(yyDollar[3].val.(string))})
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:78
		{
			PrintYySym(yyDollar[3])
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:79
		{
			PrintYySym(yyDollar[3])
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:80
		{
			PrintYySym(yyDollar[3])
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:81
		{
			PrintYySym(yyDollar[3])
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:84
		{
			WriteFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:85
		{
			WriteFile(GlobalVarRead(yyDollar[3]), yyDollar[5].val.(string))
		}
	case 33:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:86
		{
			WriteFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 34:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:89
		{
			AppendFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 35:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:90
		{
			AppendFile(GlobalVarRead(yyDollar[3]), yyDollar[5].val.(string))
		}
	case 36:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:91
		{
			AppendFile(yyDollar[3].val, yyDollar[5].val.(string))
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:94
		{
			yyVAL = yySymType{val: ReadFile(yyDollar[3].val.(string))}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:97
		{
			PrintType(yyDollar[3])
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:100
		{
			yyVAL = yySymType{val: StrToJson(yyDollar[3].val)}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:101
		{
			yyVAL = yySymType{val: StrToJson(GlobalVarRead(yyDollar[3]))}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:102
		{
			yyVAL = yySymType{val: StrToJson(WalkField(yyDollar[3].val.(string)))}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:105
		{
			yyVAL = yySymType{val: ToJsonStr(yyDollar[3].val)}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:106
		{
			yyVAL = yySymType{val: ToJsonStr(GlobalVarRead(yyDollar[3]))}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:107
		{
			yyVAL = yySymType{val: ToJsonStr(WalkField(yyDollar[3].val.(string)))}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:110
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:111
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:112
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:113
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:114
		{
			GlobalVarFieldWrite(yyDollar[1].val.(string), yyDollar[3].val)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:115
		{
			GlobalVarFieldWrite(yyDollar[1].val.(string), yyDollar[3].val)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:116
		{
			GlobalVarFieldWrite(yyDollar[1].val.(string), yyDollar[3].val)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:117
		{
			GlobalVarFieldWrite(yyDollar[1].val.(string), yyDollar[3].val)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:120
		{
			yyVAL = yySymType{val: ValAdd(yyDollar[1].val, yyDollar[3].val)}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:121
		{
			yyVAL = yySymType{val: ValMinus(yyDollar[1].val, yyDollar[3].val)}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:122
		{
			yyVAL = yySymType{val: ValMul(yyDollar[1].val, yyDollar[3].val)}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:123
		{
			yyVAL = yySymType{val: ValDiv(yyDollar[1].val, yyDollar[3].val)}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:126
		{
			yyVAL = yyDollar[1]
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:127
		{
			yyVAL = yyDollar[1]
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:128
		{
			yyVAL = yySymType{val: GlobalVarRead(yyDollar[1])}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:129
		{
			yyVAL = yyDollar[1]
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:130
		{
			yyVAL = yyDollar[1]
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:131
		{
			yyVAL = yySymType{val: WalkField(yyDollar[1].val.(string))}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:132
		{
			yyVAL = yyDollar[1]
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:133
		{
			yyVAL = yyDollar[1]
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:134
		{
			yyVAL = yyDollar[1]
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:135
		{
			yyVAL = yyDollar[1]
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:136
		{
			yyVAL = yyDollar[1]
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:139
		{
			yyVAL = yyDollar[2]
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:142
		{
			yyVAL = yyDollar[1]
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:143
		{
			yyVAL = joinHeaders(yyDollar[1], yyDollar[2])
		}
	case 71:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:144
		{
			yyVAL = yySymType{val: nil}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:147
		{
			yyVAL = yySymType{val: yyDollar[2].val}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:148
		{
			yyVAL = yySymType{val: yyDollar[2].val}
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:149
		{
			yyVAL = yySymType{val: nil}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:152
		{
			yyVAL = HttpSend("GET", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:153
		{
			yyVAL = HttpSend("PUT", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:154
		{
			yyVAL = HttpSend("POST", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:155
		{
			yyVAL = HttpSend("DELETE", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:156
		{
			yyVAL = HttpSend("HEAD", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:160
		{
			yyVAL = yySymType{val: yyDollar[1].val.(string) + "." + yyDollar[3].val.(string)}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:163
		{
			yyVAL = yySymType{val: yyDollar[1].val.(string) + "." + yyDollar[3].val.(string)}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:166
		{
			yyVAL = yySymType{val: fmt.Sprintf("%s.[%d]", yyDollar[1].val, yyDollar[4].val)}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:169
		{
			yyVAL = yySymType{val: fmt.Sprintf("%s.[%d]", yyDollar[1].val, yyDollar[4].val)}
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:174
		{
			yyVAL = yySymType{val: ToStr(GlobalVarRead(yyDollar[3]))}
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:175
		{
			yyVAL = yySymType{val: ToStr(WalkField(yyDollar[3].val.(string)))}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:178
		{
			RunIfCond(GlobalVarRead(yyDollar[2]), yyDollar[3].val)
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:179
		{
			RunIfCond(yyDollar[2].val, yyDollar[3].val)
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:180
		{
			RunIfCond(WalkField(yyDollar[2].val.(string)), yyDollar[3].val)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:183
		{
			RepeatBlock(yyDollar[2].val, yyDollar[3].val)
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:184
		{
			RepeatBlock(GlobalVarRead(yyDollar[2]), yyDollar[3].val)
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:185
		{
			RepeatBlock(WalkField(yyDollar[2].val.(string)), yyDollar[3].val)
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:186
		{
			RepeatBlock(yyDollar[2].val, yyDollar[3].val)
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:189
		{
			yyVAL = yySymType{val: map[string]any{}}
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:192
		{
			yyVAL = yySymType{val: CalcLen(GlobalVarRead(yyDollar[3]))}
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:193
		{
			yyVAL = yySymType{val: CalcLen(WalkField(yyDollar[3].val.(string)))}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:196
		{
			DoForEach(GlobalVarRead(yyDollar[2]), yyDollar[3].val)
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:197
		{
			DoForEach(WalkField(yyDollar[2].val.(string)), yyDollar[3].val)
		}
	}
	goto yystack /* stack new state and value */
}
