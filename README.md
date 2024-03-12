# yikes

yikes - simple programming language used to learn lexer and parser.

The lexer is hand-written, it's fairly simple and straight-forward. The parser is generated using goyacc.

Supports:
- number (int64 / float64)
- bool
- value assignment, even for fields of an object
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
- read field values recursively using `.field` syntax.

*Sadly, this language interprets code on the fly, it won't go very far without generating an AST. But it's still a good way to learn :D*

# Demo

Some random code copied from terminal:

```

 Welcome to yikes v0.0.1
 Github: https://github.com/curtisnewbie/yikes

>>> Entering interactive mode:
>>>
>>> res = GET 'http://localhost:80' -d json('{}')
Sending 'GET http://localhost:80', header: <nil>, body: map[string]interface {}{}
>>> print(res)
map[body:[] code:200]
>>> print(res.code)
200
>>> print('http status code is ' + res.code)
http status code is 200
>>>
>>> print(res.body)
[]
>>> res.body = 'empty body'
>>> print(res.body)
empty body
>>>
>>> res.body = json('{ "name" : "shady" }')
>>> print(res.body)
map[name:shady]
>>> print(res.body.name)
shady
>>>
>>> a = 1 + 2 * 3
>>> print(a) // this should be 7
7
>>>
>>> b = 1 * 2 + 3
>>> print(b) // this should be 5
5
>>> print()

>>>
>>> j = json('{ "abc": {"jpg": "yo"} }')
>>> print(j.abc.jpg)
yo
>>> print(jsonstr(j))
{"abc":{"jpg":"yo"}}
>>> s = jsonstr(j)
>>> print(s)
{"abc":{"jpg":"yo"}}
>>> type(s)
s: <string>
>>>
>>> // some syntax error
>>> json(
Fatal Error: 'syntax error: unexpected Json' - 4
    json(
       ^
```