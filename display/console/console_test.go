package console

import (
	"strings"
	"testing"
)

func TestConsoleWrite(t *testing.T) {
	w := strings.Builder{}
	dev, err := New(&w)
	if err != nil {
		t.Fatal(dev)
	}
	err = dev.Cls()
	if err != nil {
		t.Fatal(err)
	}
	err = dev.Println("test message!")
	if err != nil {
		t.Fatal(err)
	}

	content := w.String()

	if "test message!\n" != content {
		t.Fatalf("String %s is not expected", content)
	}
}
