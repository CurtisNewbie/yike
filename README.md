# yikes

yikes - simple programming language used to learn lexer and parser

Supports:
- number (int64 / float64)
- bool
- value assignment
- `+ - * /` arithmatic operations
- print(...) function
- type(...) functino
- write(string, file) function
- append(string, file) function
- json(...) function
- jsonstr(...) function
- string concatenation `+`
- HTTP requests, e.g.,
    - `GET 'http://localhost:80' -h 'my-header:123' -d '123'`
    - `GET 'http://localhost:80' -h 'my-header:123' -d json('{}')`
- comments, e.g.,
    - `//...`
    - `print('yo') // this prints yo`
- walk fields recursively using `.FIELD` syntax.

*Sadly, this language interprets code on the fly, it won't go very far without generating an AST. But it's still a good way to learn :D*

```
>>> yikes v0.0.1 https://github.com/curtisnewbie/yikes
>>> Entered interactive mode:
>>>
>>>
>>> res = GET 'http://localhost:80' -d json('{}')
Sending 'GET http://localhost:80', parser.yySymType{yys:56, val:interface {}(nil)}, parser.yySymType{yys:68, val:map[string]interface {}{}}
>>>
>>> print(res)
map[body:[] code:200]
>>> res
res: map[string]interface {}{"body":[]uint8{}, "code":200} <map[string]interface {}>
>>>
>>>
>>>
>>> print(res.code)
200
>>>
>>>
>>> a = 1
>>> print(a)
1
>>>
>>> b = 'i am ' + 'very good'
>>> print(b)
i am very good
>>>
>>>
>>> exit
bye
```