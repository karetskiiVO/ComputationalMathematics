package main

import (
	"image/color"
	"log"

	"github.com/sbwhitecap/tqdm"
	. "github.com/sbwhitecap/tqdm/iterators"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	results := make([]ExperimentResult, 0)

	tqdm.With(Interval(1, 101), "simulation", func(v interface{}) (brk bool) {
		n := 10 * v.(int)

		results = append(results, Experiment(n))
		return
	})

	p := plot.New()
	p.Title.Text = "Сравнение алгоритмоа"
	p.X.Label.Text = "N"
	p.Y.Label.Text = "time, s"

	// p.Y.Scale = plot.LogScale{}
	// p.Y.Tick.Marker = plot.LogTicks{}

	GramSchmidtPoints := make(plotter.XYs, len(results))
	ModifiedGramSchmidtPoints := make(plotter.XYs, len(results))
	HouseholderPoints := make(plotter.XYs, len(results))

	for i, result := range results {
		n := float64(result.n)
		GramSchmidtPoints[i] = plotter.XY{X: n, Y: result.GramSchmidtDuration.Seconds()}
		ModifiedGramSchmidtPoints[i] = plotter.XY{X: n, Y: result.ModifiedGramSchmidtDuration.Seconds()}
		HouseholderPoints[i] = plotter.XY{X: n, Y: result.Householder.Seconds()}
	}

	GramSchmidtLine, err := plotter.NewLine(GramSchmidtPoints)
	GramSchmidtLine.Color = color.RGBA{255, 0, 0, 255}
	if err != nil {
		log.Fatal(err)
	}

	ModifiedGramSchmidtLine, err := plotter.NewLine(ModifiedGramSchmidtPoints)
	ModifiedGramSchmidtLine.Color = color.RGBA{0, 255, 0, 255}
	if err != nil {
		log.Fatal(err)
	}

	HouseholderLine, err := plotter.NewLine(HouseholderPoints)
	HouseholderLine.Color = color.RGBA{0, 0, 255, 255}
	if err != nil {
		log.Fatal(err)
	}

	p.Add(GramSchmidtLine, ModifiedGramSchmidtLine, HouseholderLine)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "./add2/results/plot.png"); err != nil {
		log.Fatal(err)
	}
}
