[![Build Status](https://travis-ci.org/hi-ogawa/go_playground.svg?branch=master)](https://travis-ci.org/hi-ogawa/go_playground)

### Go Playground

Development:

```
$ mkdir -p $GOPATH/src/github.com/hi-ogawa
$ ln -s $PWD $GOPATH/src/github.com/hi-ogawa/playground
```

Run program:

```
$ go run p10310/main.go < p10310/sample.input
```

Testing:

```
$ go test -v github.com/hi-ogawa/playground/p10310
```

Testing in Docker:

```
$ docker-compose run --rm test
```
