package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/cast"
)

var (
	debug = false
)

func ValAdd(a any, b any) any {
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta.Kind() == reflect.Float64 || tb.Kind() == reflect.Float64 {
		return cast.ToFloat64(a) + cast.ToFloat64(b)
	} else if ta.Kind() == reflect.String || tb.Kind() == reflect.String {
		return cast.ToString(a) + cast.ToString(b)
	} else {
		return a.(int64) + b.(int64)
	}
}

func ValMinus(a any, b any) any {
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta.Kind() == reflect.Float64 || tb.Kind() == reflect.Float64 {
		return cast.ToFloat64(a) - cast.ToFloat64(b)
	} else {
		return a.(int64) - b.(int64)
	}
}

func ValMul(a any, b any) any {
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta.Kind() == reflect.Float64 || tb.Kind() == reflect.Float64 {
		return cast.ToFloat64(a) * cast.ToFloat64(b)
	} else {
		return a.(int64) * b.(int64)
	}
}

func ValDiv(a any, b any) any {
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta.Kind() == reflect.Float64 || tb.Kind() == reflect.Float64 {
		return cast.ToFloat64(a) / cast.ToFloat64(b)
	} else {
		return a.(int64) / b.(int64)
	}
}

func GlobalVar() map[string]any {
	return vmrt.globalvar
}

func GlobalVarRead(y yySymType) any {
	n := y.val.(string)
	return GlobalVar()[n]
}

func GlobalVarFieldWrite(expr string, val any) {
	if expr == "" {
		return
	}
	tk := strings.Split(expr, ".")
	if len(tk) < 1 {
		return
	}
	label := tk[0]

	v := GlobalVar()[label]
	if v == nil {
		v = map[string]any{}
	}

	m, ok := v.(map[string]any)
	if !ok {
		panic("Invalid datatype")
	}

	for i := 1; i < len(tk)-1; i++ {
		if f, ok := m[tk[i]]; ok {
			if _, ok := f.(map[string]any); !ok {
				panic("Invalid datatype of fields")
			}
		} else {
			m[tk[i]] = map[string]any{}
		}
		m = m[tk[i]].(map[string]any)
	}
	m[tk[len(tk)-1]] = val
	GlobalVar()[label] = v
}

func GlobalVarWrite(y yySymType, val any) {
	Debugf("GlobalVarWrite: %v, %v\n", y.val, val)
	n := y.val.(string)
	GlobalVar()[n] = val
}

func GlobalVarCopy(to yySymType, from yySymType) {
	v := GlobalVarRead(from)
	GlobalVarWrite(to, v)
}

func PrintYySymDebug(y yySymType) {
	n := y.val.(string)
	glob := GlobalVar()
	v := glob[n]
	fmt.Printf("%v: %#v <%v>\n", n, v, reflect.TypeOf(v))
}

func PrintYySym(y yySymType) {
	fmt.Printf("%v\n", y.val)
}

func PrintGlobalYySym(y yySymType) {
	n := y.val.(string)
	glob := GlobalVar()
	v := glob[n]
	fmt.Printf("%v\n", v)
}

func PrintType(y yySymType) {
	n := y.val.(string)
	val := vmrt.globalvar[n]
	typ := reflect.TypeOf(val)
	fmt.Printf("%s: <%v>\n", n, typ)
}

func SyntaxError() {
	println("syntax error")
}

func EnableDebug() {
	debug = true
}

func Debug(args ...any) {
	if debug {
		j := make([]any, len(args)+1)
		j[0] = "[DEBUG]"
		for i := range args {
			j[i+1] = args[i]
		}
		fmt.Println(j...)
	}
}

func Debugf(s string, args ...any) {
	if debug {
		fmt.Printf("[DEBUG] "+s+"\n", args...)
	}
}

func HttpSend(method string, url string, header yySymType, body yySymType) yySymType {
	fmt.Printf("Sending '%v %v', header: %#v, body: %#v\n", method, url, header.val, body.val)

	var typ string
	var br io.Reader = nil
	if body.val != nil {
		if s, ok := body.val.(string); ok {
			br = bytes.NewReader([]byte(s))
			typ = "text"
		} else {
			byt, err := json.Marshal(body.val)
			if err != nil {
				panic(fmt.Sprintf("failed to marshal json, %#v, %v", body, err))
			}
			br = bytes.NewReader(byt)
			typ = "json"
		}
	}

	r, err := http.NewRequest(method, url, br)
	if err != nil {
		fmt.Printf("ERROR: failed to send http request, %v", err)
		return yySymType{}
	}

	if br != nil {
		if typ == "json" {
			r.Header.Set("Content-Type", "application/json")
		} else if typ == "text" {
			r.Header.Set("Content-Type", "text/plain")
		}
	}
	if harr, ok := header.val.([]string); ok {
		for _, v := range harr {
			tk := strings.SplitN(v, ":", 1)
			if len(tk) > 1 {
				r.Header.Add(tk[0], tk[1])
			}
		}
	} else if hs, ok := header.val.(string); ok {
		tk := strings.SplitN(hs, ":", 1)
		if len(tk) > 1 {
			r.Header.Add(tk[0], tk[1])
		}
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Printf("ERROR: failed to send http request, %v\n", err)
		return yySymType{}
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ERROR: failed to read http response, %v\n", err)
		return yySymType{}
	}

	Debugf("Returned response: %d, body: %s", res.StatusCode, buf)
	return yySymType{val: map[string]any{"code": res.StatusCode, "body": buf}}
}

func joinHeaders(h1 yySymType, h2 yySymType) yySymType {
	v1 := h1.val
	v2 := h2.val

	arr := []string{}
	if varr, ok := v1.([]string); ok {
		arr = append(arr, varr...)
	} else {
		arr = append(arr, cast.ToString(v1))
	}

	if varr, ok := v2.([]string); ok {
		arr = append(arr, varr...)
	} else {
		arr = append(arr, cast.ToString(v2))
	}
	return yySymType{
		val: arr,
	}
}

func WalkField(expr string) any {
	if expr == "" {
		return nil
	}
	tk := strings.Split(expr, ".")
	if len(tk) < 1 {
		return nil
	}
	v := GlobalVar()[tk[0]]
	for i := 1; i < len(tk)-1; i++ {
		if m, ok := v.(map[string]any); ok {
			v = m[tk[i]]
		} else {
			return nil
		}
	}
	return v.(map[string]any)[tk[len(tk)-1]]
}

func StrToMap(v any) any {
	if s, ok := v.(string); ok {
		var m map[string]any
		if err := json.Unmarshal([]byte(s), &m); err != nil {
			panic(fmt.Sprintf("Invalid json, %v, %v", s, err))
		}
		return m
	}
	return nil
}

func WriteFile(v any, file string) {
	f, err := ReadWriteFile(file)
	if err != nil {
		panic(fmt.Errorf("failed to open file %v, %v", file, err))
	}
	f.Truncate(0)
	defer f.Close()
	if s, ok := v.(string); ok {
		_, err := f.Write([]byte(s))
		if err != nil {
			panic(fmt.Errorf("failed to write file %v, %v", file, err))
		}
		fmt.Printf("Wrote to file %v\n", file)
	} else {
		panic("can only write string to file")
	}
}

func AppendFile(v any, file string) {
	f, err := AppendableFile(file)
	if err != nil {
		panic(fmt.Errorf("failed to open file %v, %v", file, err))
	}
	defer f.Close()
	if s, ok := v.(string); ok {
		_, err := f.Write([]byte(s))
		if err != nil {
			panic(fmt.Errorf("failed to write file %v, %v", file, err))
		}
		fmt.Printf("Appended to file %v\n", file)
	} else {
		panic("can only append string to file")
	}
}

// Open file with 0666 permission.
func OpenFile(name string, flag int) (*os.File, error) {
	return os.OpenFile(name, flag, 0666)
}

// Create appendable file with 0666 permission.
func AppendableFile(name string) (*os.File, error) {
	return OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY)
}

// Create readable & writable file with 0666 permission.
func ReadWriteFile(name string) (*os.File, error) {
	return OpenFile(name, os.O_CREATE|os.O_RDWR)
}

func ToJsonStr(v any) any {
	if s, ok := v.(string); ok {
		return s
	}

	buf, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Errorf("failed to marshal json, %v", err))
	}
	return string(buf)
}
