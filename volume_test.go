package gounit

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTsp(t *testing.T) {	
	assert.True(t, Tsp(3).Equal(Tsp(3)))
	assert.True(t, Tsp(1).NotEqual(Tsp(3)))
}

func TestTspAndTbsp(t *testing.T) {	
	assert.True(t, Tbsp(1).Equal(Tsp(3)))
	assert.True(t, Tbsp(1).NotEqual(Tsp(4)))
}

func TestOzAndTbsp(t *testing.T) {	
	assert.True(t, Oz(1).Equal(Tbsp(2)))
	assert.True(t, Oz(1).NotEqual(Tbsp(3)))
}

func TestVolumePlus(t *testing.T) {	
	assert.True(t, Tsp(3).Plus(Tbsp(1)).Equal(Oz(1)))
}

func TestVolumeMinus(t *testing.T) {	
	volume, ok := Oz(1).Minus(Tbsp(1))
	assert.True(t, ok && volume.Equal(Tsp(3)))
}

func TestVolumeInvalidMinus(t *testing.T) {	
	_, ok := Tsp(1).Minus(Tsp(2))
	assert.False(t, ok)
}

func TestFormatVolumeWithBaseUnit(t *testing.T) {	
	assert.Equal(t, "12 TSP", Oz(2).String())
}