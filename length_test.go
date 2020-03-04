package gounit

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMile(t *testing.T) {	
	assert.True(t, Mile(3).Equal(Mile(3)))
	assert.True(t, Mile(3).NotEqual(Mile(2)))
}

func TestMileAndYard(t *testing.T) {	
	assert.True(t, Mile(1).Equal(Yard(1760)))
	assert.True(t, Mile(1).NotEqual(Yard(1761)))
	
	assert.True(t, Yard(3).Equal(Yard(3)))
	assert.True(t, Yard(3).NotEqual(Yard(4)))
}

func TestLength(t *testing.T) {	
	assert.True(t, Yard(1).Equal(Feet(3)))
	assert.True(t, Feet(1).Equal(Inch(12)))

	assert.True(t, Mile(1).Equal(Inch(1760 * 3 * 12)))
}

func TestLengthPlus(t *testing.T) {	
	assert.True(t, Inch(13).Plus(Inch(11)).Equal(Feet(2)))
}

func TestLengthMinus(t *testing.T) {	
	length, ok := Yard(2).Minus(Feet(3))
	assert.True(t, ok && length.Equal(Yard(1)))
}

func TestLengthInvalidMinus(t *testing.T) {	
	_, ok := Yard(1).Minus(Yard(2))
	assert.False(t, ok)
}

func TestFormatLengthWithBaseUnit(t *testing.T) {	
	assert.Equal(t, "24 INCH", Feet(2).String())
}