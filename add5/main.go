package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	nodes := []float64{-1, 0, 1}

	values := [][]float64{
		{2, -8, 56},
		{1, 0, 0},
		{2, 8, 56},
	}

	interp := NewHermite(nodes, values)

	p := plot.New()

	p.Title.Text = "plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	PlotHermite(p, -1, 1, 10000, interp, color.RGBA{255, 0, 0, 255})
	PlotFunc(p, -1, 1, 10000, func(x float64) float64 { return math.Pow(x, 8) + 1 }, color.RGBA{0, 255, 0, 255})

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "plot.png"); err != nil {
		panic(err)
	}
}

func PlotHermite(p *plot.Plot, xmin, xmax float64, steps int, herm *Hermite, color color.Color) {
	ptsSpline := make(plotter.XYs, steps)
	xMin, xMax := xmin, xmax
	dx := (xMax - xMin) / float64(steps-1)
	for i := range steps {
		ptsSpline[i].X = xMin + float64(i)*dx
		y := herm.Evaluate(ptsSpline[i].X)
		if math.Abs(y) > 2 {
			y = 0
		}
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
