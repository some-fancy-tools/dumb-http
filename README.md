# dumb-http
## Simple HTTP Server

## Usage

```
dumb-http [-path path-to-serve] [port]
```
### Example
```
$ dumb-http -path ./docs
Serving at http://0.0.0.0:8000/ from ./docs
127.0.0.1 - - [05/Jul/2018 18:08:16] "GET / HTTP/1.1" 200 38 0.0000 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:61.0) Gecko/20100101 Firefox/61.0"
```

>Inspired by Python module http.server

> Reference: https://gist.github.com/cespare/3985516
