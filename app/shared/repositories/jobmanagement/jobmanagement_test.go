package jobmanagement_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(123,123, "Are Equal")
}

func TestAddNewEmployerJob(t *testing.T) {
	assert := assert.New(t)

	p, err := AddNewEmployerJob("abcde", 1, 1, "Programmer", "Seattle", 1, 1, "Entry Level Position")

	assert.Equal(p,1, "Are Equal")
}