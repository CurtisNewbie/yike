%{
package parser

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

%right '='
%left '+' '-'
%left '*' '/'

%start expression

%%

expression:
    assignment
    | statement
    | Comment

statement:
    print_st
    | label_st
    | type_st
    | arith_st
    | network_st
    | field_st
    | write_st
    | append_st

label_st:
    Label { PrintYySymDebug($1) }

Value:
    String | Number | Bool

print_st:
    Print '(' Value ')' { PrintYySym($3) }
    | Print '(' Label ')' { PrintGlobalYySym($3) }
    | Print '(' ')' { println("") }
    | Print '(' field_st ')' { PrintYySym($3) }
    | Print '(' network_st ')' { PrintYySym($3) }
    | Print '(' jsonstr_st ')' { PrintYySym($3) }
    | Print '(' arith_st ')' { PrintYySym($3) }

write_st:
    Write '(' Value ',' String ')' { WriteFile($3.val, $5.val.(string)) }
    | Write '(' Label ',' String ')' { WriteFile(GlobalVarRead($3), $5.val.(string)) }
    | Write '(' jsonstr_st ',' String ')' { WriteFile($3.val, $5.val.(string)) }

append_st:
    Append '(' Value ',' String ')' { AppendFile($3.val, $5.val.(string)) }
    | Append '(' Label ',' String ')' { AppendFile(GlobalVarRead($3), $5.val.(string)) }
    | Append '(' jsonstr_st ',' String ')' { AppendFile($3.val, $5.val.(string)) }

type_st:
    Type '(' Label ')' { PrintType($3) }

json_st:
    Json '(' String ')' { $$ = yySymType{ val: StrToMap($3.val) } }
    | Json '(' Label ')' { $$ = yySymType{ val: StrToMap(GlobalVarRead($3)) } }
    | Json '(' field_st ')' { $$ = yySymType{ val: StrToMap($3.val) } }

jsonstr_st:
    JsonStr '(' String ')' { $$ = yySymType{ val: ToJsonStr($3.val) } }
    | JsonStr '(' Label ')' { $$ = yySymType{ val: ToJsonStr(GlobalVarRead($3)) } }
    | JsonStr '(' field_st ')' { $$ = yySymType{ val: ToJsonStr($3.val) } }

assignment:
    Label '=' eval_expr { GlobalVarWrite($1, $3.val) }
    | Label '=' { SyntaxError() }
    | Label '=' network_st { GlobalVarWrite($1, $3.val) }
    | Label '=' field_st { GlobalVarWrite($1, $3.val) }
    | Label '=' json_st { GlobalVarWrite($1, $3.val) }
    | Label '=' jsonstr_st { GlobalVarWrite($1, $3.val) }

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

header_sg:
    Header String { $$ = $2 }

header_st:
    header_sg { $$ = $1 }
    | /* empty */
    | header_sg header_sg { $$ = joinHeaders($1, $2) }

body_st:
    /* empty */
    | Body String { $$ = yySymType{ val: $2.val } }
    | Body json_st { $$ = yySymType{ val: $2.val } }

network_st:
    Get String header_st body_st { $$ = HttpSend("GET", $2.val.(string), $3, $4) }
    | Put String header_st body_st { $$ = HttpSend("PUT", $2.val.(string), $3, $4) }
    | Post String header_st body_st { $$ = HttpSend("POST", $2.val.(string), $3, $4) }
    | Delete String header_st body_st { $$ = HttpSend("DELETE", $2.val.(string), $3, $4) }
    | Head String header_st body_st { $$ = HttpSend("HEAD", $2.val.(string), $3, $4) }

/*
    TODO: make field_st lazy, only evals on terminal statement like print or something else
    so that we can not only use it to read but also to write
*/
field_st:
    Label '.' Label {
        v := GlobalVarRead($1)
        $$ = yySymType{ val: WalkField(v, $3.val) }
    }
    | field_st '.' Label {
        $$ = yySymType{ val: WalkField($1.val, $3.val) }
    }