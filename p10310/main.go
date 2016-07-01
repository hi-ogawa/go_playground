package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type P struct {
	x float64
	y float64
}

func main() {
	mainRW(os.Stdin, os.Stdout)
}

func mainRW(r io.Reader, w io.Writer) {
	for {
		var n int
		var gop, dog P
		var pts []P
		_, err := fmt.Fscanf(r, "%d %f %f %f %f\n",
			&n, &(gop.x), &(gop.y), &(dog.x), &(dog.y),
		)
		if err != nil {
			break
		}
		for i := 0; i < n; i++ {
			var pt P
			fmt.Fscanf(r, "%f %f\n", &(pt.x), &(pt.y))
			pts = append(pts, pt)
		}
		fmt.Fscanf(r, "\n")
		p, ok := solve(gop, dog, pts)
		fmt.Fprint(w, formatOutput(p, ok))
	}
}

func solve(gop P, dog P, pts []P) (P, bool) {
	var p P
	ok := false
	for _, pt := range pts {
		if distance(gop, pt)*2 <= distance(dog, pt) {
			ok = true
			p = pt
			break
		}
	}
	return p, ok
}

func distance(p, q P) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func formatOutput(p P, ok bool) string {
	if ok {
		return fmt.Sprintf(
			"The gopher can escape through the hole at (%.3f,%.3f).\n",
			p.x, p.y,
		)
	} else {
		return "The gopher cannot escape.\n"
	}
}
