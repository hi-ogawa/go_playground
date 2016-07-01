package main

import (
  "testing"
  . "github.com/smartystreets/goconvey/convey"
  "os"
  "bytes"
  "io/ioutil"
)

func TestSpec(t *testing.T) {
  Convey("p10310", t, func() {
    Convey("solve", func() {
      Convey("example 0", func() {
        _, ok := solve(
          P{ 1.0, 1.0 },
          P{ 2.0, 2.0 },
          []P{
            P{ 1.5, 1.5 },
          },
        )
        So(ok, ShouldEqual, false)
      })

      Convey("example 1", func() {
        p, ok := solve(
          P{ 2.0, 2.0 },
          P{ 1.0, 1.0 },
          []P{
            P{ 1.5, 1.5 },
            P{ 2.5, 2.5 },
          },
        )
        So(p, ShouldResemble, P{ 2.5, 2.5 })
        So(ok, ShouldEqual, true)
      })
    })

    Convey("mainRW", func() {
      reader, _ := os.Open("sample.input")
      answerWriter, _ := os.Open("sample.output")
      answer, _ := ioutil.ReadAll(answerWriter)
      var writer bytes.Buffer
      mainRW(reader, &writer)
      So(writer.String(), ShouldEqual, string(answer))
    })
  })
}
