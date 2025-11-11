package matrix

import (
	"errors"
	"sync"
)

type matrix struct {
	Rows int
	Cols int
	Data []float64
}

func New(rows, cols int) *matrix {
	return &matrix{
		Rows: rows,
		Cols: cols,
		Data: make([]float64, rows*cols),
	}
}

func (m *matrix) At(i, j int) float64 {
	return m.Data[i*m.Cols+j]
}

func (m *matrix) Set(i, j int, v float64) {
	m.Data[i*m.Cols+j] = v
}

func sameShape(a, b *matrix) bool {
	return a.Rows == b.Rows && a.Cols == b.Cols
}

func AddSeq(a, b *matrix) (*matrix, error) {
	if !sameShape(a, b) {
		return nil, errors.New("shape mismatch")
	}

	c := New(a.Rows, a.Cols)

	ad, bd, cd := a.Data, b.Data, c.Data
	_ = ad[len(ad)-1]
	_ = bd[len(bd)-1]
	_ = cd[len(cd)-1]

	for i := range cd {
		cd[i] = ad[i] + bd[i]
	}

	return c, nil
}

func AddPar(a, b *matrix, workers int) (*matrix, error) {
	if !sameShape(a, b) {
		return nil, errors.New("shape mismatch")
	}

	if workers > a.Rows {
		workers = a.Rows
	}

	c := New(a.Rows, a.Cols)
	rowsPerWorker := a.Rows / workers
	cols := a.Cols

	wg := new(sync.WaitGroup)
	for w := 0; w < workers; w++ {
		i0 := w * rowsPerWorker
		i1 := i0 + rowsPerWorker

		if w == workers-1 {
			i1 = a.Rows // можно раскидать по другим воркерам, чтобы forall abs(w1-w1) <= 1
		}

		wg.Go(func() {
			start := i0 * cols
			end := i1 * cols

			ad := a.Data[start:end]
			bd := b.Data[start:end]
			cd := c.Data[start:end]

			for k := 0; k < len(ad); k++ {
				cd[k] = ad[k] + bd[k]
			}
		})
	}

	wg.Wait()
	return c, nil
}
