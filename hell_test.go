package hellomod

import "testing"

func TestHello(t *testing.T){
	want := "v3:Hello World!"
	if Hello() != want {
		t.Errorf("Hello() != %s",want)
	}
}
