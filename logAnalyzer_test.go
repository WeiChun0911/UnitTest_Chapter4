package chap4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TooShortFileNameShouldLogError(t *testing.T) {
	fw := &FakeWebService{}
	la := NewLogAnalyzer(fw)

	tooShortFileName := "abc.txt"
	la.Analyze(tooShortFileName)
	
	expected := "File name too short: " + "abc.txt"
	actual := fw.lastError

	assert.Equal(t, expected, actual, "they should be equal")
}