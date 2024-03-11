# yikes

yikes - simple programming language used to learn lexer and parser

Supports:
- value assignment
- `+ - * /` arithmatic operations
- print(...) function
- string concatenation `+`
- HTTP requests, e.g.,
    - `GET 'http://localhost:80' -h 'my-header:123' -d '123'`
    - `GET 'http://localhost:80' -h 'my-header:123' -d json('{}')`
- comments, e.g.,
    - `//...`
    - `print('yo') // this prints yo`
- walk fields recursively using `.FIELD` syntax.
