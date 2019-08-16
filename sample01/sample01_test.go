package sample01

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld("hoge")
	exptedted := "hello world, hoge"
	if actual != exptedted {
		t.Errorf("actual %v\nwant %v", actual, exptedted)
	}
}
