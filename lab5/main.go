package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func TestProblem() ODESystem {
	// Тестовая система: u' = -u, v' = -2v
	return ODESystem{
		Function: func(t float64, u []float64) []float64 {
			return []float64{-u[0], -2 * u[1]}
		},
		Jacobian: func(t float64, u []float64) [][]float64 {
			return [][]float64{
				{-1, 0},
				{0, -2},
			}
		},
		ExactSolution: func(t float64) []float64 {
			return []float64{math.Exp(-t), math.Exp(-2 * t)}
		},
		Dimension: 2,
	}
}

func PlotSolutions(t []float64, numerical [][]float64, exact func(float64) []float64, filename string) error {
	p := plot.New()
	p.Title.Text = "Numerical vs Exact Solution"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	ptsNum1 := make(plotter.XYs, len(t))
	for i := range t {
		ptsNum1[i].X = t[i]
		ptsNum1[i].Y = numerical[i][0]
	}
	lineNum1, err := plotter.NewLine(ptsNum1)
	if err != nil {
		return err
	}
	lineNum1.Color = plotutil.Color(0)
	lineNum1.Dashes = []vg.Length{vg.Points(1), vg.Points(1)}
	p.Add(lineNum1)
	p.Legend.Add("Numerical u1", lineNum1)

	ptsExact1 := make(plotter.XYs, len(t))
	for i := range t {
		ptsExact1[i].X = t[i]
		ptsExact1[i].Y = exact(t[i])[0]
	}
	lineExact1, err := plotter.NewLine(ptsExact1)
	if err != nil {
		return err
	}
	lineExact1.Color = plotutil.Color(0)
	p.Add(lineExact1)
	p.Legend.Add("Exact u1", lineExact1)

	ptsNum2 := make(plotter.XYs, len(t))
	for i := range t {
		ptsNum2[i].X = t[i]
		ptsNum2[i].Y = numerical[i][1]
	}
	lineNum2, err := plotter.NewLine(ptsNum2)
	if err != nil {
		return err
	}
	lineNum2.Color = plotutil.Color(1)
	lineNum2.Dashes = []vg.Length{vg.Points(1), vg.Points(1)}
	p.Add(lineNum2)
	p.Legend.Add("Numerical u2", lineNum2)

	ptsExact2 := make(plotter.XYs, len(t))
	for i := range t {
		ptsExact2[i].X = t[i]
		ptsExact2[i].Y = exact(t[i])[1]
	}
	lineExact2, err := plotter.NewLine(ptsExact2)
	if err != nil {
		return err
	}
	lineExact2.Color = plotutil.Color(1)
	p.Add(lineExact2)
	p.Legend.Add("Exact u2", lineExact2)

	if err := p.Save(10*vg.Inch, 6*vg.Inch, filename); err != nil {
		return err
	}

	return nil
}

func main() {
	table := NewButcherTable()
	testSystem := TestProblem()
	solver := NewRK4Solver(table, testSystem)

	T := 2.0
	u0 := []float64{1.0, 1.0}
	steps := 20

	t, numerical := solver.Solve(T, u0, steps)

	// fmt.Println("Time\tNumerical u1\tNumerical u2\tExact u1\tExact u2")
	// for i := range t {
	// 	exact := testSystem.ExactSolution(t[i])
	// 	fmt.Printf("%.4f\t%.6f\t%.6f\t%.6f\t%.6f\n",
	// 		t[i], numerical[i][0], numerical[i][1], exact[0], exact[1])
	// }

	err := PlotSolutions(t, numerical, testSystem.ExactSolution, "solution.png")
	if err != nil {
		fmt.Printf("Error plotting: %v\n", err)
	} else {
		fmt.Println("Plot saved to solution.png")
	}

	// Исследование сходимости на последовательности сеток
	// fmt.Println("\nConvergence study:")
	// stepsList := []int{10, 20, 40, 80, 160}
	// errors := make([]float64, len(stepsList))

	// for idx, n := range stepsList {
	// 	_, num := solver.Solve(T, u0, n)
	// 	// Вычисляем ошибку в конечной точке
	// 	exactFinal := testSystem.ExactSolution(T)
	// 	err1 := math.Abs(num[n][0] - exactFinal[0])
	// 	err2 := math.Abs(num[n][1] - exactFinal[1])
	// 	errors[idx] = math.Sqrt(err1*err1 + err2*err2)
	// 	fmt.Printf("Steps: %d, Error: %.6f\n", n, errors[idx])
	// }
}
