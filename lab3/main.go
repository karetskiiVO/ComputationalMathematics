package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func PlotSpline(xs, ys []float64, spline Spline, title string) {
	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := make(plotter.XYs, len(xs))
	for i := range xs {
		pts[i].X = xs[i]
		pts[i].Y = ys[i]
	}
	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	line.Color = color.RGBA{0, 0, 0, 0}
	p.Add(line)

	steps := 1000

	ptsSpline := make(plotter.XYs, steps)
	xMin, xMax := xs[0], xs[len(xs)-1]
	dx := (xMax - xMin) / float64(steps-1)
	for i := range steps {
		ptsSpline[i].X = xMin + float64(i)*dx
		ptsSpline[i].Y = spline.Predict(ptsSpline[i].X)
	}

	lineSpline, err := plotter.NewLine(ptsSpline)
	if err != nil {
		panic(err)
	}
	lineSpline.Color = plotutil.Color(1)
	p.Add(lineSpline)

	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}
	scatter.GlyphStyle.Color = plotutil.Color(2)
	p.Add(scatter)

	if err := p.Save(6*vg.Inch, 6*vg.Inch, title+".png"); err != nil {
		panic(err)
	}
}

func main() {
	var spline Spline = &CubicFreeSpline{}

	xUniform := []float64{0, 1, 2, 3, 4, 5}
	yUniform := []float64{0, 1, 4, 9, 16, 25}
	spline.Fit(xUniform, yUniform)
	PlotSpline(xUniform, yUniform, spline, "Uniform_Grid")

	xNonUniform := []float64{0, 1, 2.5, 3.5, 5}
	yNonUniform := []float64{0, 1, 6.25, 12.25, 25}
	spline.Fit(xNonUniform, yNonUniform)
	PlotSpline(xNonUniform, yNonUniform, spline, "Non_Uniform_Grid")
}
