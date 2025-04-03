package main

import (
	"fmt"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	res := make([]T2, len(s))
	for i, si := range s {
		res[i] = f(si)
	}

	return res
}

func main() {
	p := plot.New()
	p.Title.Text = "Точность"
	p.X.Label.Text = "N"
	p.Y.Label.Text = "time, s"

	p.X.Scale = plot.LogScale{}
	p.X.Tick.Marker = plot.LogTicks{}
	p.Y.Scale = plot.LogScale{}
	p.Y.Tick.Marker = plot.LogTicks{}

	scale := 0.001

	N := 12
	points := make(plotter.XYs, N)
	for i := range N {
		points[i].X = scale
		points[i].Y = Experiment(scale)
		scale *= 2
	}

	fmt.Println(points)

	line, _ := plotter.NewLine(points)
	p.Add(line)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "./lab4/results/plot.png"); err != nil {
		log.Fatal(err)
	}
}
