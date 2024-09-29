package integers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	assert.Equal(t, sum, 4)
}
