package main

import (
	"image/color"
	"log"
	"math"

	"github.com/sbwhitecap/tqdm"
	. "github.com/sbwhitecap/tqdm/iterators"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	results := make([][]ExperimentResult, 0)

	series := 100

	tqdm.With(Interval(1, 2000), "simulation", func(v interface{}) (brk bool) {
		n := 10 * v.(int)

		buf := make([]ExperimentResult, series)
		for j := 0; j < series; j++ {
			buf[j] = Experiment(n, GenerateRandomSample)
		}

		results = append(results, buf)
		return
	})

	p := plot.New()
	p.Title.Text = "Относительная ошибка"
	p.X.Label.Text = "N"
	p.Y.Label.Text = "val"

	p.Y.Scale = plot.LogScale{}
	p.Y.Tick.Marker = plot.LogTicks{}

	simpleLinePoints := make(plotter.XYs, len(results))
	treeLinePoints := make(plotter.XYs, len(results))
	kahanLinePoints := make(plotter.XYs, len(results))

	for i, result := range results {
		diff := func(val, absolute float64) float64 {
			res := math.Abs(val-absolute) / absolute

			if math.IsNaN(res) {
				return 0
			}
			if math.IsInf(res, 0) {
				return 0
			}

			return res
		}

		x := float64(result[0].n)

		var simple, tree, kahan float64
		for _, val := range result {
			simple += diff(val.simple, val.absolute)
			tree += diff(val.tree, val.absolute)
			kahan += diff(val.kahan, val.absolute)
		}

		simple /= float64(series)
		tree /= float64(series)
		kahan /= float64(series)

		simpleLinePoints[i] = plotter.XY{X: x, Y: simple}
		treeLinePoints[i] = plotter.XY{X: x, Y: tree}
		kahanLinePoints[i] = plotter.XY{X: x, Y: kahan}
	}

	simpleLine, err := plotter.NewLine(simpleLinePoints)
	simpleLine.Color = color.RGBA{1, 0, 0, 0}
	if err != nil {
		log.Fatal(err)
	}
	treeLine, err := plotter.NewLine(treeLinePoints)
	treeLine.Color = color.RGBA{0, 1, 0, 0}
	if err != nil {
		log.Fatal(err)
	}
	kahanLine, err := plotter.NewLine(kahanLinePoints)
	kahanLine.Color = color.RGBA{0, 0, 1, 0}
	if err != nil {
		log.Fatal(err)
	}

	p.Add(simpleLine, treeLine, kahanLine)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "./sem1/results/plot.png"); err != nil {
		log.Fatal(err)
	}
}
