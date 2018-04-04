package strategy

import "testing"

func TestAddition_Apply(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 7},
		{2, 10, 12},
		{14, 3, 17},
		{1, 9, 10},
		{-1, 1, 0},
		{999, 999, 1998},
	}
	op := Operation{Addition{}}
	for _, v := range tests {
		res := op.Operator.Apply(v.a, v.b)
		if res != v.c {
			t.Fatalf("Addition %d + %d = %d expect %d ", v.a, v.b, res, v.c)
		}
	}

}
func TestMultiplication_Apply(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 12},
		{2, 10, 20},
		{14, 3, 42},
		{1, 9, 9},
		{-1, 1, -1},
		{10, 10, 100},
	}
	op := Operation{Multiplication{}}
	for _, v := range tests {
		res := op.Operator.Apply(v.a, v.b)
		if res != v.c {
			t.Fatalf("Multiplication %d * %d= %d expect %d ", v.a, v.b, res, v.c)
		}
	}
}
