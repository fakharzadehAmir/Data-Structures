package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	m := NewMatrix(3, 3)
	err := m.Insert(10, 0, 0)
	assert.NoError(t, err)

	// Test edge values
	err = m.Insert(20, 4, 4)
	assert.Error(t, err)
}

func TestAdd(t *testing.T) {
	m1 := NewMatrix(2, 2)
	m1.Insert(1, 0, 0)
	m1.Insert(2, 0, 1)
	m1.Insert(3, 1, 0)
	m1.Insert(4, 1, 1)

	m2 := NewMatrix(2, 2)
	m2.Insert(5, 0, 0)
	m2.Insert(6, 0, 1)
	m2.Insert(7, 1, 0)
	m2.Insert(8, 1, 1)

	err := m1.Add(m2)
	assert.NoError(t, err)

	expected := NewMatrix(2, 2)
	expected.Insert(6, 0, 0)
	expected.Insert(8, 0, 1)
	expected.Insert(10, 1, 0)
	expected.Insert(12, 1, 1)

	assert.Equal(t, expected.Items, m1.Items)

	// Test edge values
	m3 := NewMatrix(3, 3)
	err = m1.Add(m3)
	assert.Error(t, err)
}

func TestSubtract(t *testing.T) {
	m1 := NewMatrix(2, 2)
	m1.Insert(1, 0, 0)
	m1.Insert(2, 0, 1)
	m1.Insert(3, 1, 0)
	m1.Insert(4, 1, 1)

	m2 := NewMatrix(2, 2)
	m2.Insert(5, 0, 0)
	m2.Insert(6, 0, 1)
	m2.Insert(7, 1, 0)
	m2.Insert(8, 1, 1)

	err := m1.Subtract(m2)
	assert.NoError(t, err)

	expected := NewMatrix(2, 2)
	expected.Insert(-4, 0, 0)
	expected.Insert(-4, 0, 1)
	expected.Insert(-4, 1, 0)
	expected.Insert(-4, 1, 1)

	// Test edge values
	m3 := NewMatrix(3, 3)
	err = m1.Subtract(m3)
	assert.Error(t, err)
}

func TestMultiply(t *testing.T) {
	m1 := NewMatrix(2, 3)
	m1.Insert(1, 0, 0)
	m1.Insert(2, 0, 1)
	m1.Insert(3, 0, 2)
	m1.Insert(4, 1, 0)
	m1.Insert(5, 1, 1)
	m1.Insert(6, 1, 2)

	m2 := NewMatrix(3, 2)
	m2.Insert(7, 0, 0)
	m2.Insert(8, 0, 1)
	m2.Insert(9, 1, 0)
	m2.Insert(10, 1, 1)
	m2.Insert(11, 2, 0)
	m2.Insert(12, 2, 1)

	err := m1.Multiply(m2)
	assert.NoError(t, err)

	expected := NewMatrix(2, 2)
	expected.Insert(58, 0, 0)
	expected.Insert(64, 0, 1)
	expected.Insert(139, 1, 0)
	expected.Insert(154, 1, 1)

	assert.Equal(t, expected.Items, m1.Items)

	// Test edge values
	m3 := NewMatrix(4, 4)
	err = m1.Multiply(m3)
	assert.Error(t, err)
}

func TestScalarMultiply(t *testing.T) {
	m := NewMatrix(2, 2)
	m.Insert(1, 0, 0)
	m.Insert(2, 0, 1)
	m.Insert(3, 1, 0)
	m.Insert(4, 1, 1)

	m.ScalarMultiply(2)

	expected := NewMatrix(2, 2)
	expected.Insert(2, 0, 0)
	expected.Insert(4, 0, 1)
	expected.Insert(6, 1, 0)
	expected.Insert(8, 1, 1)

	assert.Equal(t, expected.Items, m.Items)

}

func TestTransposition(t *testing.T) {
	m := NewMatrix(2, 3)
	m.Insert(1, 0, 0)
	m.Insert(2, 0, 1)
	m.Insert(3, 0, 2)
	m.Insert(4, 1, 0)
	m.Insert(5, 1, 1)
	m.Insert(6, 1, 2)

	transpose := m.Transposition()

	expected := NewMatrix(3, 2)
	expected.Insert(1, 0, 0)
	expected.Insert(4, 0, 1)
	expected.Insert(2, 1, 0)
	expected.Insert(5, 1, 1)
	expected.Insert(3, 2, 0)
	expected.Insert(6, 2, 1)

	assert.Equal(t, expected.Items, transpose.Items)
}

func TestDeterminant(t *testing.T) {
	// Test 1x1 matrix
	m1 := NewMatrix(1, 1)
	m1.Insert(5, 0, 0)
	det, err := m1.Determinant()
	assert.NoError(t, err)
	assert.Equal(t, 5.0, det)

	// Test 2x2 matrix
	m2 := NewMatrix(2, 2)
	m2.Insert(2, 0, 0)
	m2.Insert(4, 0, 1)
	m2.Insert(1, 1, 0)
	m2.Insert(3, 1, 1)
	det, err = m2.Determinant()
	assert.NoError(t, err)
	assert.Equal(t, 2.0, det)

	// Test 3x3 matrix
	m3 := NewMatrix(3, 3)
	m3.Insert(1, 0, 0)
	m3.Insert(2, 0, 1)
	m3.Insert(3, 0, 2)
	m3.Insert(4, 1, 0)
	m3.Insert(5, 1, 1)
	m3.Insert(6, 1, 2)
	m3.Insert(7, 2, 0)
	m3.Insert(8, 2, 1)
	m3.Insert(9, 2, 2)
	det, err = m3.Determinant()
	assert.NoError(t, err)
	assert.Equal(t, 0.0, det)

	// Test edge values
	m4 := NewMatrix(2, 3)
	_, err = m4.Determinant()
	assert.Error(t, err)
}

func TestInverse(t *testing.T) {
	// Test non-square matrix
	m1 := NewMatrix(2, 3)
	_, err := m1.Inverse()
	assert.Error(t, err)

	// Test singular matrix
	m2 := NewMatrix(2, 2)
	m2.Insert(2, 0, 0)
	m2.Insert(4, 0, 1)
	m2.Insert(1, 1, 0)
	m2.Insert(2, 1, 1)
	_, err = m2.Inverse()
	assert.Error(t, err)

	// Test valid 2x2 matrix
	m3 := NewMatrix(2, 2)
	m3.Insert(1, 0, 0)
	m3.Insert(2, 0, 1)
	m3.Insert(3, 1, 0)
	m3.Insert(4, 1, 1)

	expectedInverse := NewMatrix(2, 2)
	expectedInverse.Insert(-2, 0, 0)
	expectedInverse.Insert(1, 0, 1)
	expectedInverse.Insert(1.5, 1, 0)
	expectedInverse.Insert(-0.5, 1, 1)

	inverse, err := m3.Inverse()
	assert.NoError(t, err)
	assert.Equal(t, expectedInverse.Items, inverse.Items)

	// Test edge values
	m4 := NewMatrix(3, 3)
	_, err = m4.Inverse()
	assert.Error(t, err)
}
