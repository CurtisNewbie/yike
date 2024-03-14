%{
package parser

import "fmt"

%}

%union{
    val any
}

%token String
%token Number
%token Print
%token Label
%token Type
%token Get Put Post Delete Head
%token Header Body
%token Json
%token JsonStr
%token Comment
%token Bool
%token Write
%token Append
%token StringFunc
%token CodeBlock
%token If
%token For
%token Read
%token Map
%token Len
%token ForEach

%right '='
%left '+' '-'
%left '*' '/'
%right Body Header

%start expressions

%%

expressions:
    expression
    | expressions expression

expression:
    assignment
    | statement
    | Comment
    | Value

statement:
    print_st
    | label_st
    | type_st
    | arith_st
    | network_st
    | field_st
    | write_st
    | read_st
    | append_st
    | if_st
    | for_st
    | foreach_st

label_st:
    Label { PrintYySymDebug($1) }

Value:
    String | Number | Bool

print_st:
    Print '(' Value ')' { PrintYySym($3) }
    | Print '(' Label ')' { PrintGlobalYySym($3) }
    | Print '(' ')' { println("") }
    | Print '(' field_st ')' { PrintYySym(yySymType{ val: WalkField($3.val.(string)) }) }
    | Print '(' network_st ')' { PrintYySym($3) }
    | Print '(' jsonstr_st ')' { PrintYySym($3) }
    | Print '(' arith_st ')' { PrintYySym($3) }
    | Print '(' string_st ')' { PrintYySym($3) }

write_st:
    Write '(' Value ',' String ')' { WriteFile($3.val, $5.val.(string)) }
    | Write '(' Label ',' String ')' { WriteFile(GlobalVarRead($3), $5.val.(string)) }
    | Write '(' jsonstr_st ',' String ')' { WriteFile($3.val, $5.val.(string)) }

append_st:
    Append '(' Value ',' String ')' { AppendFile($3.val, $5.val.(string)) }
    | Append '(' Label ',' String ')' { AppendFile(GlobalVarRead($3), $5.val.(string)) }
    | Append '(' jsonstr_st ',' String ')' { AppendFile($3.val, $5.val.(string)) }

read_st:
    Read '(' String ')' { $$ = yySymType{ val: ReadFile($3.val.(string)) } }

type_st:
    Type '(' Label ')' { PrintType($3) }

json_st:
    Json '(' String ')' { $$ = yySymType{ val: StrToJson($3.val) } }
    | Json '(' Label ')' { $$ = yySymType{ val: StrToJson(GlobalVarRead($3)) } }
    | Json '(' field_st ')' { $$ = yySymType{ val: StrToJson(WalkField($3.val.(string))) } }

jsonstr_st:
    JsonStr '(' String ')' { $$ = yySymType{ val: ToJsonStr($3.val) } }
    | JsonStr '(' Label ')' { $$ = yySymType{ val: ToJsonStr(GlobalVarRead($3)) } }
    | JsonStr '(' field_st ')' { $$ = yySymType{ val: ToJsonStr(WalkField($3.val.(string))) } }

assignment:
    Label '=' eval_expr { GlobalVarWrite($1, $3.val) }
    | Label '=' network_st { GlobalVarWrite($1, $3.val) }
    | Label '=' json_st { GlobalVarWrite($1, $3.val) }
    | Label '=' jsonstr_st { GlobalVarWrite($1, $3.val) }
    | field_st '=' json_st { GlobalVarFieldWrite($1.val.(string), $3.val) }
    | field_st '=' eval_expr { GlobalVarFieldWrite($1.val.(string), $3.val) }
    | field_st '=' jsonstr_st { GlobalVarFieldWrite($1.val.(string), $3.val) }
    | field_st '=' network_st { GlobalVarFieldWrite($1.val.(string), $3.val) }

arith_st:
    eval_expr '+' eval_expr { $$ = yySymType{ val: ValAdd($1.val, $3.val) } }
    | eval_expr '-' eval_expr { $$ = yySymType{ val: ValMinus($1.val, $3.val) } }
    | eval_expr '*' eval_expr { $$ = yySymType{ val: ValMul($1.val, $3.val) } }
    | eval_expr '/' eval_expr { $$ = yySymType{ val: ValDiv($1.val, $3.val) } }

eval_expr:
    Number { $$ = $1 }
    | String { $$ = $1 }
    | Label { $$ = yySymType{ val: GlobalVarRead($1) } }
    | arith_st { $$ = $1 }
    | Bool { $$ = $1 }
    | field_st { $$ = yySymType{ val: WalkField($1.val.(string)) } }
    | string_st { $$ = $1 }
    | read_st { $$ = $1 }
    | map_st { $$ = $1 }
    | len_st { $$ = $1 }

header_sg:
    Header String { $$ = $2 }

header_st:
    header_sg { $$ = $1 }
    | header_st header_sg { $$ = joinHeaders($1, $2) }
    | /* empty */ { $$ = yySymType{ val : nil } }

body_st:
    Body String { $$ = yySymType{ val: $2.val } }
    | Body json_st { $$ = yySymType{ val: $2.val } }
    | /* empty */ { $$ = yySymType{ val: nil } }

network_st:
    Get String header_st body_st { $$ = HttpSend("GET", $2.val.(string), $3, $4) }
    | Put String header_st body_st { $$ = HttpSend("PUT", $2.val.(string), $3, $4) }
    | Post String header_st body_st { $$ = HttpSend("POST", $2.val.(string), $3, $4) }
    | Delete String header_st body_st { $$ = HttpSend("DELETE", $2.val.(string), $3, $4) }
    | Head String header_st body_st { $$ = HttpSend("HEAD", $2.val.(string), $3, $4) }

/* this is lazy, it doesn't walk the fields recursively until meeting a terminal statement */
field_st:
    Label '.' Label {
        $$ = yySymType{ val: $1.val.(string) + "." + $3.val.(string) }
    }
    | field_st '.' Label {
        $$ = yySymType{ val: $1.val.(string) + "." + $3.val.(string) }
    }
    | Label '.' '[' Number ']' {
        $$ = yySymType{ val: fmt.Sprintf("%s.[%d]", $1.val, $4.val) }
    }
    | field_st '.' '[' Number ']' {
        $$ = yySymType{ val: fmt.Sprintf("%s.[%d]", $1.val, $4.val) }
    }

string_st:
    StringFunc '(' Label ')' { $$ = yySymType{ val: ToStr(GlobalVarRead($3)) } }
    | StringFunc '(' field_st ')' { $$ = yySymType{ val: ToStr(WalkField($3.val.(string))) } }

if_st:
    If Label CodeBlock { RunIfCond(GlobalVarRead($2), $3.val) }
    | If Bool CodeBlock { RunIfCond($2.val, $3.val) }
    | If field_st CodeBlock { RunIfCond(WalkField($2.val.(string)), $3.val) }

for_st:
    For Number CodeBlock { RepeatBlock($2.val, $3.val) }
    | For Label CodeBlock { RepeatBlock(GlobalVarRead($2), $3.val) }
    | For field_st CodeBlock { RepeatBlock(WalkField($2.val.(string)), $3.val) }
    | For len_st CodeBlock { RepeatBlock($2.val, $3.val) }

map_st:
    Map '(' ')' { $$ = yySymType{ val: map[string]any{} } }

len_st:
    Len '(' Label ')' { $$ = yySymType{ val: CalcLen(GlobalVarRead($3)) } }
    | Len '(' field_st ')' { $$ = yySymType{ val: CalcLen(WalkField($3.val.(string))) } }

foreach_st:
    | ForEach Label CodeBlock { DoForEach(GlobalVarRead($2), $3.val) }
    | ForEach field_st CodeBlock { DoForEach(WalkField($2.val.(string)), $3.val) }