package adapter

import "testing"

func TestClassAdapter(t *testing.T) {
	classAdapter := &ClassAdapter{}
	classAdapter.request()
}

func TestObjAdapter(t *testing.T) {
	adaptee := &Adaptee{}
	objAdapter := NewObjAdapter(adaptee)
	objAdapter.request()
}
