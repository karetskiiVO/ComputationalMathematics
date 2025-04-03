package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	f := math.Sin
	xmin, xmax := 0.0, math.Pi
	ymin, ymax := 0.0, 1.0
	exact := 2.0

	p := plot.New()

	p.Title.Text = "Зависимость ошибки интегрирования от количества итераций"
	p.X.Label.Text = "Количество итераций (log scale)"
	p.Y.Label.Text = "Абсолютная ошибка (log scale)"
	p.X.Scale = plot.LogScale{}
	p.Y.Scale = plot.LogScale{}
	p.X.Tick.Marker = plot.LogTicks{}
	p.Y.Tick.Marker = plot.LogTicks{}

	iterations := []int{10, 100, 1000, 10000, 100000, 1000000}
	mmkPoints := make(plotter.XYs, len(iterations))
	rectPoints := make(plotter.XYs, len(iterations))

	for i, iter := range iterations {
		mmk := MMKIntegrate(f, xmin, xmax, ymin, ymax, iter)
		rect := Integrate(f, xmin, xmax, iter)

		mmkPoints[i].X = float64(iter)
		mmkPoints[i].Y = math.Abs(mmk - exact)
		rectPoints[i].X = float64(iter)
		rectPoints[i].Y = math.Abs(rect - exact)
	}

	err := plotutil.AddLinePoints(p,
		"Монте-Карло", mmkPoints,
		"Прямоугольники", rectPoints,
	)
	if err != nil {
		panic(err)
	}

	if err := p.Save(8*vg.Inch, 6*vg.Inch, "./add7/results/plot.png"); err != nil {
		panic(err)
	}
}
