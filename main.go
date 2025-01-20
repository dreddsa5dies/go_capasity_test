package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/Arafatk/glot"
)

var (
	iteration int
)

func init() {
	flag.IntVar(&iteration, "i", 3, "How many iterations (iteration * 512), but not > 10")
}

func main() {
	flag.Parse()

	if iteration > 10 {
		println("Please " + os.Args[0] + " -h")
		os.Exit(0)
	}

	for j := 1; j <= iteration; j++ {
		t := 512 * j
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

		plot, err := glot.NewPlot(3, false, false)
		if err != nil {
			log.Print(err)
			continue
		}
		plot.SetTitle("capacity")
		style := "lines"

		for k, v := range res {
			plot.AddPointGroup(k, style, v)
		}

		plot.SavePlot(fmt.Sprintf("capacity_%d.png", t))
		plot.Close()
	}
}
