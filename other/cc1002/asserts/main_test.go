package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {

	// this case wouldn't need any table driven tests. A bit straight forward

	assert := assert.New(t)
	min := 0
	max := 10
	randomNumber := randomNumber(min, max)

	if randomNumber < min || randomNumber > max {
		t.Errorf("number %d should be contained in range [%d, %d]", randomNumber, min, max)
	}

	if !assert.GreaterOrEqual(randomNumber, min) || !assert.LessOrEqual(randomNumber, max) {
		err := errors.New("number not in between [min, max]")
		assert.Errorf(err, "number %d should be contained in range [%d, %d]", randomNumber, min, max)
	}
}
