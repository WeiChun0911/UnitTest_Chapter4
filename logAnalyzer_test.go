package chap4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_TooShortFileNameShouldLogError(t *testing.T) {
// 	fw := &FakeWebService{}
// 	la := NewLogAnalyzer(fw)

// 	tooShortFileName := "abc.txt"
// 	la.Analyze(tooShortFileName)

// 	expected := "File name too short: " + "abc.txt"
// 	actual := fw.lastError

// 	assert.Equal(t, expected, actual, "they should be equal")
// }

func Test_TooShortFileNameShouldSendEmail(t *testing.T) {
	fw := &FakeWebService{}
	fe := &FakeEmailService{}
	la := NewLogAnalyzer(fw, fe)

	tooShortFileName := "abc.txt"
	fw.ShouldThrowError = true;
	la.Analyze(tooShortFileName)
	
	expected_to := "support@going.cloud"
	expected_body := "Fake exception!"
	expected_subject := "Can't Log"
	
	actual_to := fe.To
	actual_body := fe.Body
	actual_subject := fe.Subject

	assert.Equal(t, expected_to, actual_to, "they should be equal")
	assert.Equal(t, expected_body, actual_body, "they should be equal")
	assert.Equal(t, expected_subject, actual_subject, "they should be equal")
}
