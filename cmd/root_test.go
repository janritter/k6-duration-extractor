package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveSliceElement(t *testing.T) {
	test := []string{"a", "b", "c"}
	result := removeSliceElement(test, 1)

	assert.Equal(t, []string{"a", "c"}, result)
}

func TestGetNameFromFilename(t *testing.T) {
	test := "test1.csv"
	result := getNameFromFilename(test)

	assert.Equal(t, "test1", result)
}
