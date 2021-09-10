package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	f, err := NewParser("yaml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(f.CreateParser().Parse())
}
