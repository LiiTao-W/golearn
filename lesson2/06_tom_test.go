package main

import (
	//"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	if output != expectOutput {
		t.Errorf("Expected %s do not match actual %s", expectOutput, output)
	}
}

func TestJudgeTrue(t *testing.T) {
	output := judge(60)
	if output != true {
		t.Errorf("Expected %v do not match actual %v", true, output)
	}
}

func TestJudgeFalse(t *testing.T) {
	output := judge(40)
	if output != false {
		t.Errorf("Expected %v do not match actual %v", false, output)
	}
}

//func TestHelloTomTrue(t *testing.T) {
//	output := HelloTom()
//	expectOutput := "Jerry"
//	assert.Equal(t, expectOutput, output)
//}
