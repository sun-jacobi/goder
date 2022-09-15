package judger

import "testing"

func TestGetFileName(t *testing.T) {
	a := "helloworld.cpp"
	b := "你好世界.cs"
	if filename := getFileName(a); filename != "helloworld" {
		t.Fatalf(a + filename)
	}
	if filename := getFileName(b); filename != "你好世界" {
		t.Fatalf(b + filename)
	}
}
