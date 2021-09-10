package simple_factory

import "testing"

func TestSimpleFactory(t *testing.T) {
	p, err := NewParser("json")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p.Parse())
}
