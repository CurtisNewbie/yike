package parser

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/spf13/cast"
)

const (
	Github = "https://github.com/curtisnewbie/yikes"
)

var (
	lineBreakPat = regexp.MustCompile(`([^\\](\\n)|^\\n$)`)
)

type KwTableInf struct {
	Kw       string
	Type     int
	Callback func(lval *yySymType)
}

var (
	vmrt = newVm()

	// order matters here
	kwTable = []KwTableInf{
		{Kw: "jsonstr", Type: JsonStr},
		{Kw: "print", Type: Print},
		{Kw: "type", Type: Type},
		{Kw: "json", Type: Json},
		{Kw: "GET", Type: Get},
		{Kw: "PUT", Type: Put},
		{Kw: "POST", Type: Post},
		{Kw: "HEAD", Type: Head},
		{Kw: "DELETE", Type: Delete},
		{Kw: "true", Type: Bool, Callback: func(lval *yySymType) { lval.val = true }},
		{Kw: "false", Type: Bool, Callback: func(lval *yySymType) { lval.val = false }},
		{Kw: "write", Type: Write},
		{Kw: "append", Type: Append},
		{Kw: "string", Type: StringFunc},
		{Kw: "if", Type: If},
		{Kw: "else", Type: Else},
		{Kw: "for", Type: For},
		{Kw: "read", Type: Read},
		{Kw: "map", Type: Map},
		{Kw: "len", Type: Len},
		{Kw: "foreach", Type: ForEach},
	}
)

type vm struct {
	script    string
	offset    int
	globalvar map[string]any
}

func (v *vm) Lex(lval *yySymType) int {
	for {
		if c, ok := v.next(); ok {
			switch {
			case c == '\'' || c == '"':
				return v.parseString(lval, c)
			case c == '\n' || c == ' ' || c == '\t':
				v.move(1)
				continue
			case c == '{':
				return v.parseCodeBlock(lval)
			case unicode.IsLetter(c) || c == '_':
				if d, ok := v.parseKeywords(lval); ok {
					return d
				}
				return v.parseLabel(lval)
			case unicode.IsDigit(c):
				return v.parseNumber(lval)
			case (c == '/' && v.lookAheadIs(1, '/')) || c == '#':
				return v.parseComment(lval)
			case c == '-':
				if v.lookAheadIs(1, 'h') || v.lookAheadIs(1, 'H') { // -h / -H
					v.move(2)
					return Header
				}
				if v.lookAheadIs(1, 'd') { // -d
					v.move(2)
					return Body
				}
				Debugf("default for '-' %v", string(c))
				return v.skipOne()
			default:
				Debugf("default %v", string(c))
				return v.skipOne()
			}
		} else {
			return 0
		}
	}
}

func (v *vm) Error(s string) {
	row := 1
	start := 0
	for i := 0; i < v.offset; i++ {
		if v.script[i] == '\n' {
			row += 1
			if i+1 < v.offset {
				start = i + 1
			}
		}
	}
	es := fmt.Sprintf("'%v' - at row %d, offset %v\n  -> '%v'", s, row, v.offset, v.script[start:v.offset])
	panic(es)
}

func (v *vm) next() (rune, bool) {
	// Debug("next")
	if v.offset >= len(v.script) {
		return 0, false
	}
	c := v.script[v.offset]
	return rune(c), true
}

func (v *vm) lookAheadAt(n int) (rune, bool) {
	Debugf("lookAheadAt %v", n)
	if v.offset+n >= len(v.script) {
		return 0, false
	}
	c := v.script[v.offset+n]
	return rune(c), true
}

func (v *vm) lookAheadIs(n int, c rune) bool {
	if ch, ok := v.lookAheadAt(1); ok && ch == c {
		return true
	}
	return false
}

func (v *vm) move(gap int) {
	v.offset = v.offset + gap
	Debugf("move %v to %v", gap, v.offset)
}

func (v *vm) skipOne() int {
	c := v.script[v.offset]
	v.move(1)
	return int(c)
}

func (v *vm) parseNumber(lval *yySymType) int {
	Debugf("parseNumber, offset: %v", v.offset)
	i := 1
	isFloat := false
	for {
		if c, ok := v.lookAheadAt(i); ok {
			if c != '.' && !unicode.IsDigit(c) {
				break
			} else {
				if c == '.' {
					isFloat = true
				}
				i += 1
			}
		} else {
			break
		}
	}
	if isFloat {
		lval.val = cast.ToFloat64(v.script[v.offset : v.offset+i])
	} else {
		lval.val = cast.ToInt64(v.script[v.offset : v.offset+i])
	}
	Debugf("parseNumber -> label.val: %v", lval.val)
	v.move(i)
	Debugf("parseNumber -> offset: %v", v.offset)
	return Number
}

func (v *vm) parseLabel(lval *yySymType) int {
	Debugf("parseLabel, offset: %v", v.offset)
	i := 1
	for {
		if c, ok := v.lookAheadAt(i); ok {
			if !unicode.IsLetter(c) && c != '_' {
				break
			} else {
				i += 1
			}
		} else {
			break
		}
	}
	lval.val = v.script[v.offset : v.offset+i]
	Debugf("label.val: %v", lval.val)
	v.move(i)
	Debugf("offset: %v", v.script[v.offset:])
	return Label
}

func (v *vm) parseString(lval *yySymType, quote rune) int {
	Debugf("parsestring, starting at: %v", v.offset)
	i := 1 // [0] is the quote
	for {
		if c, ok := v.lookAheadAt(i); ok {
			if c == quote {
				ss := v.script[v.offset+1 : v.offset+i]
				lval.val = lineBreakPat.ReplaceAllStringFunc(ss, func(s string) string {
					if len(s) == 2 {
						return "\n"
					}
					return s[:len(s)-2] + "\n"
				})
				// Debugf("string lval.val: '%v', %v, %v", v.script[v.offset+1:v.offset+i], v.offset, v.offset+i)
				v.move(i + 1)
				return String
			}
			i += 1
		} else {
			break
		}
	}
	panic(fmt.Sprintf("illegal syntax for string, %#v", v))
}

func (v *vm) parseKeywords(lval *yySymType) (int, bool) {
	for _, kwtp := range kwTable {
		if i, ok := v.parseKeyword(lval, kwtp.Kw, kwtp); ok {
			return i, true
		}
	}
	return 0, false
}

func (v *vm) parseKeyword(lval *yySymType, keyword string, inf KwTableInf) (int, bool) {
	kwl := len(keyword)
	if v.offset+kwl > len(v.script) {
		return 0, false
	}

	pre := v.script[v.offset : v.offset+kwl]
	if pre == keyword && (v.offset+kwl == len(v.script) || !unicode.IsLetter(rune(v.script[v.offset+kwl]))) {
		Debugf("matched: %#v", inf)
		if inf.Callback != nil {
			inf.Callback(lval)
		}
		v.move(kwl)
		return inf.Type, true
	}
	return 0, false
}

func newVm() *vm {
	return &vm{
		globalvar: make(map[string]any),
	}
}

func Run(s string, interactiveMode bool) {
	yyErrorVerbose = true

	start := time.Now()
	defer func() { Debugf("VM ran for %v\n", time.Since(start)) }()
	interpret(s, interactiveMode)
}

func interpret(s string, interactiveMode bool) {
	if interactiveMode {
		s = strings.TrimSpace(s)
		if s == "" {
			return
		}
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("Fatal Error: %v\n\n", e)
			}
		}()
	}
	vmrt.script = s
	vmrt.offset = 0
	yyParse(vmrt)
	Debugf("vars: %#v\n", vmrt.globalvar)
}

type BlockPos struct {
	Start int // pos of {
	End   int // pos of }
}

func (v *vm) parseCodeBlock(yst *yySymType) int {
	i := 1 // [0] is the quote
	for {
		c, ok := v.lookAheadAt(i)
		if !ok {
			break
		}
		if c == '}' {
			yst.val = v.script[v.offset+1 : v.offset+i]
			yst.val = BlockPos{
				Start: v.offset,
				End:   v.offset + i,
			}
			Debugf("parsed codeblock: '%v'", yst.val)
			v.move(i + 1)
			return CodeBlock
		}
		i++
	}
	panic(fmt.Sprintf("illegal syntax for codeblock, %#v", v))
}

func (v *vm) RunBlock(block BlockPos) {
	original := v.script

	// replace the original script, run the code block
	vmrt.script = v.script[block.Start+1 : block.End]
	vmrt.offset = 0
	Debugf("RunBlock: script: %v, original offset remaining: %v", vmrt.script, original[block.End+1:])
	yyParse(vmrt)

	v.script = original
	v.offset = block.End + 1
}

func (v *vm) parseComment(lval *yySymType) int {
	i := 1
	for {
		if c, ok := v.lookAheadAt(i); ok {
			if c == '\n' {
				break
			}
			i++
		} else {
			break
		}
	}
	Debugf("Comment: '%v'", v.script[v.offset:v.offset+i])
	v.move(i)
	return Comment
}
