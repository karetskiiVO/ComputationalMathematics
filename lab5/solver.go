package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// ButcherTable представляет таблицу Бутчера для метода Рунге-Кутты
type ButcherTable struct {
	A      [][]float64 // Матрица коэффициентов
	B      []float64   // Вектор весов
	C      []float64   // Вектор узлов
	Stages int         // Количество стадий
}

// NewButcherTable создает таблицу Бутчера из условия
func NewButcherTable() ButcherTable {
	return ButcherTable{
		A: [][]float64{
			{0, 0.5},
			{0.5, 0},
		},
		B:      []float64{0.5, 0.5},
		C:      []float64{0, 0.5},
		Stages: 2,
	}
}

// ODESystem представляет систему ОДУ
type ODESystem struct {
	Function      func(t float64, u []float64) []float64   // Функция правой части системы
	Jacobian      func(t float64, u []float64) [][]float64 // Якобиан системы
	ExactSolution func(t float64) []float64                // Точное решение (для тестов)
	Dimension     int                                      // Размерность системы
}

// RK4Solver реализует метод Рунге-Кутты
type RK4Solver struct {
	Table  ButcherTable
	System ODESystem
}

// NewRK4Solver создает новый решатель
func NewRK4Solver(table ButcherTable, system ODESystem) *RK4Solver {
	return &RK4Solver{
		Table:  table,
		System: system,
	}
}

// Solve решает систему ОДУ на интервале [0, T] с начальным условием u0
func (s *RK4Solver) Solve(T float64, u0 []float64, steps int) ([]float64, [][]float64) {
	h := T / float64(steps)
	time := make([]float64, steps+1)
	solution := make([][]float64, steps+1)

	// Инициализация
	for i := 0; i <= steps; i++ {
		time[i] = float64(i) * h
	}

	solution[0] = make([]float64, len(u0))
	copy(solution[0], u0)

	// Основной цикл
	for i := 1; i <= steps; i++ {
		t := time[i-1]
		uPrev := solution[i-1]

		// Вычисление стадийных значений
		k := make([][]float64, s.Table.Stages)
		for j := 0; j < s.Table.Stages; j++ {
			k[j] = make([]float64, s.System.Dimension)
		}

		// Первая стадия (явная)
		k[0] = s.System.Function(t, uPrev)

		// Вторая стадия (неявная - решаем методом Ньютона)
		// u2 = uPrev + h*a21*k1 + h*a22*k2
		// k2 = f(t + c2*h, u2)
		// Решаем нелинейное уравнение для k2
		k[1] = s.newton(t+s.Table.C[1]*h, uPrev, k[0], h, 1)

		// Собираем решение
		solution[i] = make([]float64, s.System.Dimension)
		for j := 0; j < s.System.Dimension; j++ {
			solution[i][j] = uPrev[j] + h*(s.Table.B[0]*k[0][j]+s.Table.B[1]*k[1][j])
		}
	}

	return time, solution
}

// newton реализует метод Ньютона для решения нелинейного уравнения на стадии
func (s *RK4Solver) newton(t float64, uPrev, kPrev []float64, h float64, stage int) []float64 {
	dim := s.System.Dimension
	eps := 1e-12
	maxIter := 100

	k := make([]float64, dim)
	copy(k, kPrev) // Начальное приближение

	for iter := 0; iter < maxIter; iter++ {
		// Вычисляем u = uPrev + h*sum(a[stage][j]*k[j])
		u := make([]float64, dim)
		for i := 0; i < dim; i++ {
			u[i] = uPrev[i] + h*s.Table.A[stage][0]*kPrev[i] + h*s.Table.A[stage][1]*k[i]
		}

		// Вычисляем F(k) = k - f(t, u)
		f := s.System.Function(t, u)
		F := make([]float64, dim)
		for i := 0; i < dim; i++ {
			F[i] = k[i] - f[i]
		}

		// Проверка нормы F
		norm := 0.0
		for i := 0; i < dim; i++ {
			norm += F[i] * F[i]
		}
		norm = math.Sqrt(norm)
		if norm < eps {
			break
		}

		// Вычисляем Якобиан J = I - h*a[stage][stage]*df/du
		dfdu := s.System.Jacobian(t, u)
		J := mat.NewDense(dim, dim, nil)
		for i := 0; i < dim; i++ {
			for j := 0; j < dim; j++ {
				val := -h * s.Table.A[stage][stage] * dfdu[i][j]
				if i == j {
					val += 1.0
				}
				J.Set(i, j, val)
			}
		}

		// Решаем систему J*delta = -F
		Fvec := mat.NewVecDense(dim, F)
		delta := mat.NewVecDense(dim, nil)
		err := delta.SolveVec(J, Fvec)
		if err != nil {
			panic("Newton method failed to converge")
		}

		// Обновляем k
		for i := 0; i < dim; i++ {
			k[i] -= delta.AtVec(i)
		}
	}

	return k
}
