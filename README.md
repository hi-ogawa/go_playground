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

Debugging:

```
$ brew install delve
$ cd p10310
$ dlv test github.com/hi-ogawa/playground/p10310
Type 'help' for list of commands.
(dlv) break mainRW
Breakpoint 1 set at 0x7e9ba for _/Users/hiogawa/repositories/mine/go/playground/p10310.mainRW ./main.go:21
(dlv) continue
...> _/Users/hiogawa/repositories/mine/go/playground/p10310.mainRW() ./main.go:21 (hits goroutine(5):1 total:1)
    16:		mainRW(os.Stdin, os.Stdout)
    17:	}
    18:
    19:	func mainRW(r io.Reader, w io.Writer) {
    20:		for {
=>  21:			var n int
    22:			var gop, dog P
    23:			var pts []P
    24:			_, err := fmt.Fscanf(r, "%d %f %f %f %f\n",
    25:				&n, &(gop.x), &(gop.y), &(dog.x), &(dog.y),
    26:			)
(dlv) print r
io.Reader(*struct os.File) *{
	*os.file: *struct os.file {
		fd: 5,
		name: "sample.input",
		dirinfo: *struct os.dirInfo nil,
	},
}
(dlv) exit
```
