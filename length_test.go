package quantity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMile(t *testing.T) {	
	assert.True(t, Mile(3).Equal(Mile(3)))
	assert.True(t, Mile(3).NotEqual(Mile(2)))
}
