package main

import (
	"image/color"
	"log"
	"time"

	"github.com/sbwhitecap/tqdm"
	. "github.com/sbwhitecap/tqdm/iterators"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	dim := 1000

	steepestResults := make([]ExperimentResult, 0)
	minimalResults := make([]ExperimentResult, 0)
	A, b := CreateTask(dim)

	x := mat.NewVecDense(dim, nil)

	start := time.Now()
	x.SolveVec(A, b)
	libt := time.Since(start)

	tqdm.With(Interval(1, 501), "simulation", func(v any) (brk bool) {
		n := 10 * v.(int)

		minimalResults = append(minimalResults, Experiment(A, b, n, MinimalSolver))
		steepestResults = append(steepestResults, Experiment(A, b, n, SteepestSolver))

		return
	})

	timeMinimalPoints := make(plotter.XYs, len(minimalResults))
	residualMinimalPoints := make(plotter.XYs, len(minimalResults))
	for i, result := range minimalResults {
		timeMinimalPoints[i] = plotter.XY{
			X: float64(result.n),
			Y: result.t.Seconds(),
		}
		residualMinimalPoints[i] = plotter.XY{
			X: float64(result.n),
			Y: result.residual,
		}
	}

	timeSteepestPoints := make(plotter.XYs, len(steepestResults))
	residualSteepestPoints := make(plotter.XYs, len(steepestResults))
	for i, result := range steepestResults {
		timeSteepestPoints[i] = plotter.XY{
			X: float64(result.n),
			Y: result.t.Seconds(),
		}
		residualSteepestPoints[i] = plotter.XY{
			X: float64(result.n),
			Y: result.residual,
		}
	}

	timeMinimalLine, err := plotter.NewLine(timeMinimalPoints)
	if err != nil {
		log.Fatal(err)
	}
	timeMinimalLine.Color = color.RGBA{0, 255, 0, 255}
	residualMinimalLine, err := plotter.NewLine(residualMinimalPoints)
	if err != nil {
		log.Fatal(err)
	}
	residualMinimalLine.Color = color.RGBA{0, 255, 0, 255}
	timeSteepestLine, err := plotter.NewLine(timeSteepestPoints)

	if err != nil {
		log.Fatal(err)
	}
	timeSteepestLine.Color = color.RGBA{0, 0, 255, 255}
	residualSteepestLine, err := plotter.NewLine(residualSteepestPoints)
	if err != nil {
		log.Fatal(err)
	}
	residualSteepestLine.Color = color.RGBA{0, 0, 255, 255}

	libraryLine, err := plotter.NewLine(plotter.XYs{
		plotter.XY{X: 100, Y: libt.Seconds()},
		plotter.XY{X: 5000, Y: libt.Seconds()},
	})
	if err != nil {
		log.Fatal(err)
	}
	libraryLine.Color = color.RGBA{255, 0, 0, 255}

	p := plot.New()
	p.Title.Text = "Сравнение алгоритмов(время)"
	p.X.Label.Text = "N"
	p.Y.Label.Text = "time, s"

	p.Add(timeMinimalLine, timeSteepestLine, libraryLine)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "./lab2/results/time_plot.png"); err != nil {
		log.Fatal(err)
	}

	p = plot.New()
	p.Title.Text = "Сравнение алгоритмов(ошибка)"
	p.X.Label.Text = "N"
	p.Y.Label.Text = "err"

	p.Y.Scale = plot.LogScale{}
	p.Y.Tick.Marker = plot.LogTicks{Prec: -1}
	p.Y.AutoRescale = true

	p.Add(residualMinimalLine, residualSteepestLine)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "./lab2/results/residual_plot.png"); err != nil {
		log.Fatal(err)
	}

}
