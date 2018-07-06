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
127.0.0.1 - - [06/Jul/2018 17:05:36] "GET / HTTP/1.1" 200 0 16.76µs "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:61.0) Gecko/20100101 Firefox/61.0"
127.0.0.1 - - [06/Jul/2018 17:05:37] "GET / HTTP/1.1" 304 0 173ns "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:61.0) Gecko/20100101 Firefox/61.0"
```

>Inspired by Python module http.server

> Reference: https://gist.github.com/cespare/3985516
