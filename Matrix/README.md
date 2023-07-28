# Matrix Operations in Go

This repository contains a simple Matrix package in Go, providing basic operations for manipulating matrices. The `Matrix` struct is designed to represent a two-dimensional matrix of `float64` values.

## Matrix Operations

### 1. NewMatrix

```go
func NewMatrix(rows, cols int) *Matrix
```

The `NewMatrix` function returns a new instance of `Matrix` with the specified number of rows and columns.

### 2. Insert

```go
func (m *Matrix) Insert(data float64, i, j int) error
```

The `Insert` method sets the value of a specific element at row `i` and column `j` in the matrix. It returns an error if the given index is out of bounds.

### 3. Add

```go
func (m *Matrix) Add(otherMatrix *Matrix) error
```

The `Add` method adds another matrix `otherMatrix` to the current matrix element-wise. It returns an error if the two matrices have different dimensions.

### 4. Subtract

```go
func (m *Matrix) Subtract(otherMatrix *Matrix) error
```

The `Subtract` method subtracts another matrix `otherMatrix` from the current matrix element-wise. It returns an error if the two matrices have different dimensions.

### 5. Multiply

```go
func (m *Matrix) Multiply(otherMatrix *Matrix) error
```

The `Multiply` method multiplies the current matrix by another matrix `otherMatrix`. It returns an error if the number of columns in the current matrix is not equal to the number of rows in `otherMatrix`.

### 6. ScalarMultiply

```go
func (m *Matrix) ScalarMultiply(number float64)
```

The `ScalarMultiply` method multiplies all elements of the matrix by a given scalar `number`.

### 7. Transposition

```go
func (m *Matrix) Transposition() *Matrix
```

The `Transposition` method returns a new matrix that is the transpose of the current matrix.

### 8. Determinant

```go
func (m *Matrix) Determinant() (float64, error)
```

The `Determinant` method calculates and returns the determinant of the square matrix. It returns an error if the matrix is not square.

### 9. Trace

```go
func (m *Matrix) Trace() (float64, error)
```

The `Trace` method calculates and returns the trace of the square matrix (sum of diagonal elements). It returns an error if the matrix is not square.

### 10. Inverse

```go
func (m *Matrix) Inverse() (*Matrix, error)
```

The `Inverse` method calculates and returns the inverse of the square matrix. It returns an error if the matrix is not square or if it is singular (determinant is zero).

## Running Tests

To run the tests for the Matrix package, use the following command:

```bash
go test -v
```

The tests cover various matrix operations, including insertions, additions, subtractions, multiplications, determinants, transpositions, and inverses. Additional edge cases are also tested.

Feel free to explore the code and run the tests to see how the Matrix operations work!