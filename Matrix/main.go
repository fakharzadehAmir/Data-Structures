package matrix

import "fmt"

type Matrix struct {
	Items   [][]float64
	Rows    int
	Columns int
}

func NewMatrix(rows, cols int) *Matrix {
	items := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		items[i] = make([]float64, cols)
	}
	return &Matrix{
		Items:   items,
		Rows:    rows,
		Columns: cols,
	}
}

func (m *Matrix) Insert(data float64, i, j int) error {
	if i >= 0 && i < m.Rows && j >= 0 && j < m.Columns {
		m.Items[i][j] = data
		return nil
	}
	return fmt.Errorf("wrong given index for row or column")
}

func (m *Matrix) Add(otherMatrix *Matrix) error {
	if otherMatrix.Rows == m.Rows && otherMatrix.Columns == m.Columns {
		for i := 0; i < m.Rows; i++ {
			for j := 0; j < m.Columns; j++ {
				m.Items[i][j] += otherMatrix.Items[i][j]
			}
		}
		return nil
	}
	return fmt.Errorf("number of rows or columns are different")
}

func (m *Matrix) Subtract(otherMatrix *Matrix) error {
	if otherMatrix.Rows == m.Rows && otherMatrix.Columns == m.Columns {
		for i := 0; i < m.Rows; i++ {
			for j := 0; j < m.Columns; j++ {
				m.Items[i][j] -= otherMatrix.Items[i][j]
			}
		}
		return nil
	}
	return fmt.Errorf("number of rows or columns are different")
}

func (m *Matrix) Multiply(otherMatrix *Matrix) error {
	if otherMatrix.Rows == m.Columns {
		resultMatrix := NewMatrix(m.Rows, otherMatrix.Columns)
		for i := 0; i < m.Rows; i++ {
			for j := 0; j < otherMatrix.Columns; j++ {
				sum := 0.0
				for k := 0; k < m.Columns; k++ {
					sum += m.Items[i][k] * otherMatrix.Items[k][j]
				}
				resultMatrix.Insert(sum, i, j)
			}
		}
		m.Items = resultMatrix.Items
		m.Columns = resultMatrix.Columns
		return nil
	}
	return fmt.Errorf("invalid matrix dimensions for multiplication")
}

func (m *Matrix) ScalarMultiply(number float64) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			m.Items[i][j] *= number
		}
	}
}

func (m *Matrix) Transposition() *Matrix {
	transposeMatrix := NewMatrix(m.Columns, m.Rows)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			transposeMatrix.Items[j][i] = m.Items[i][j]
		}
	}
	return transposeMatrix
}

func (m *Matrix) Determinant() (float64, error) {
	if m.Rows == m.Columns {
		if m.Rows == 1 {
			return m.Items[0][0], nil
		} else if m.Rows == 2 {
			return m.Items[0][0]*m.Items[1][1] - m.Items[0][1]*m.Items[1][0], nil
		} else {
			var determinant float64
			for j := 0; j < m.Columns; j++ {
				determinant += m.Items[0][j] * cofactor(m, 0, j)
			}
			return determinant, nil
		}
	}
	return 0, fmt.Errorf("the matrix is not square, determinant cannot be calculated")
}

func cofactor(m *Matrix, row, col int) float64 {
	subMatrix := NewMatrix(m.Rows-1, m.Columns-1)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			if row != i && col != j {
				newRow := i
				newCol := j
				if i > row {
					newRow--
				}
				if j > col {
					newCol--
				}
				subMatrix.Items[newRow][newCol] = m.Items[i][j]
			}
		}
	}
	subDet, _ := subMatrix.Determinant()
	if (row+col)%2 != 0 {
		subDet *= -1
	}
	return subDet
}

func (m *Matrix) Trace() (float64, error) {
	if m.Rows == m.Columns {
		var trace float64
		for i := 0; i < m.Rows; i++ {
			trace += m.Items[i][i]
		}
		return trace, nil
	}
	return 0, fmt.Errorf("the matrix is not square, trace cannot be calculated")
}

func (m *Matrix) Inverse() (*Matrix, error) {
	if m.Rows != m.Columns {
		return nil, fmt.Errorf("the matrix is not square, inverse does not exist")
	}

	det, err := m.Determinant()
	if err != nil {
		return nil, err
	}

	if det == 0 {
		return nil, fmt.Errorf("the matrix is singular, inverse does not exist")
	}

	adj := NewMatrix(m.Rows, m.Columns)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			adj.Items[i][j] = cofactor(m, i, j)
		}
	}

	inverse := NewMatrix(m.Rows, m.Columns)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			inverse.Items[i][j] = adj.Items[j][i] / det
		}
	}

	return inverse, nil
}
