package chap4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TooShortFileNameShouldLogError(t *testing.T) {
	la := NewLogAnalyzer()

	tooShortFileName := "abc.txt"
	la.Analyze(tooShortFileName)
	
	expected := "File name too short: " + "abc.txt"
	actual := la.service.lastError

	assert.Equal(t, expected, actual, "they should be equal")
}