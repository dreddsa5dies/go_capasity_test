package main

import (
	"runtime"

	glot "github.com/Arafatk/glot"
)

func main() {
	t := 1024 * 5
	var (
		a []int
		b []string
		c []byte
		e []float32
		f []bool
		g []map[int]int
	)

	res := make(map[string][]int, 6)

	res["int"] = make([]int, 0, t)
	res["string"] = make([]int, 0, t)
	res["byte"] = make([]int, 0, t)
	res["float32"] = make([]int, 0, t)
	res["bool"] = make([]int, 0, t)
	res["map[int]int"] = make([]int, 0, t)

	for i := 1; i < t; i++ {
		if i%3 == 0 {
			runtime.GC()
		}

		a = append(a, 0)
		b = append(b, "b")
		c = append(c, byte(10))
		e = append(e, 0.1)
		f = append(f, true)
		g = append(g, make(map[int]int))

		res["int"] = append(res["int"], cap(a))
		res["string"] = append(res["string"], cap(b))
		res["byte"] = append(res["byte"], cap(c))
		res["float32"] = append(res["float32"], cap(e))
		res["bool"] = append(res["bool"], cap(f))
		res["map[int]int"] = append(res["map[int]int"], cap(g))
	}

	plot, _ := glot.NewPlot(1, false, false)
	defer plot.Close()
	plot.SetTitle("capacity")
	style := "lines"

	for k, v := range res {
		plot.AddPointGroup(k, style, v)
	}

	plot.SavePlot("capacity_more_1024.png")
}
