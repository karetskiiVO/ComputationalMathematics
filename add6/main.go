package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	xmin, xmax := -10.0, 10.0
	f := math.Sin
	deg := 17

	interp := MakeRemez(xmin, xmax, deg, f, 1e-8, 1000)

	p := plot.New()

	p.Title.Text = "plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	PlotHermite(p, xmin, xmax, 10000, interp, color.RGBA{255, 0, 0, 255})
	PlotFunc(p, xmin, xmax, 10000, f, color.RGBA{0, 255, 0, 255})

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "plot.png"); err != nil {
		panic(err)
	}
}

func PlotHermite(p *plot.Plot, xmin, xmax float64, steps int, herm *Remez, color color.Color) {
	ptsSpline := make(plotter.XYs, steps)
	xMin, xMax := xmin, xmax
	dx := (xMax - xMin) / float64(steps-1)
	for i := range steps {
		ptsSpline[i].X = xMin + float64(i)*dx
		y := herm.Evaluate(ptsSpline[i].X)
		ptsSpline[i].Y = y
	}

	hermLine, err := plotter.NewLine(ptsSpline)
	if err != nil {
		panic(err)
	}
	hermLine.Color = color
	p.Add(hermLine)
}

func PlotFunc(p *plot.Plot, xmin, xmax float64, steps int, f func(float64) float64, color color.Color) {
	ptsSpline := make(plotter.XYs, steps)
	xMin, xMax := xmin, xmax
	dx := (xMax - xMin) / float64(steps-1)
	for i := range steps {
		ptsSpline[i].X = xMin + float64(i)*dx
		ptsSpline[i].Y = f(ptsSpline[i].X)
	}

	hermLine, err := plotter.NewLine(ptsSpline)
	if err != nil {
		panic(err)
	}
	hermLine.Color = color
	p.Add(hermLine)
}
