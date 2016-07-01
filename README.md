[![Build Status](https://travis-ci.org/hi-ogawa/go_playground.svg?branch=master)](https://travis-ci.org/hi-ogawa/go_playground)

### Go Playground

Development:

```
$ mkdir -p $GOPATH/src/github.com/hi-ogawa
$ ln -s $PWD $GOPATH/src/github.com/hi-ogawa/playground
```

Testing:

```
$ go test github.com/hi-ogawa/playground/<directory name>
```

Testing in Docker:

```
$ docker-compose run --rm test
```
