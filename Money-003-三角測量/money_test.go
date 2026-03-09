package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.True(t, five.Times(2).Equals(NewDollar(10)))
	assert.True(t, five.Times(3).Equals(NewDollar(15)))
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).Equals(NewDollar(5)))
}
