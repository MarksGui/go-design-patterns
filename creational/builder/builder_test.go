package builder

import (
	"testing"
)

func TestBuilderTruck(t *testing.T) {
	builder := &TruckBuilder{}
	director := NewDirector(builder)
	vehicle := director.Construct()
	t.Logf("%+v", vehicle)
}

func TestBuilderCar(t *testing.T) {
	builder := &CarBuilder{}
	director := NewDirector(builder)
	vehicle := director.Construct()
	t.Logf("%+v", vehicle)
}
