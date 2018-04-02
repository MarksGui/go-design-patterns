package builder

import (
	"testing"
)

func TestBuilderTruck(t *testing.T) {
	builder := &TruckBuilder{}
	director := NewDirector(builder)
	vehicle := director.Build()
	exceptStr := "卡车，车轮: 四个卡车轮子，发动机：一个卡车发动机，车门：两个卡车门"
	if vehicle != exceptStr {
		t.Fatalf("TruckBuilder fail expect %s acture %s", exceptStr, vehicle)
	}
}

func TestBuilderCar(t *testing.T) {
	builder := &CarBuilder{}
	director := NewDirector(builder)
	vehicle := director.Build()

	exceptStr := "轿车，车轮: 四个轿车轮子，发动机：一个轿车发动机，车门：四个轿车车门"
	if vehicle != exceptStr {
		t.Fatalf("TruckBuilder fail expect %s acture %s", exceptStr, vehicle)
	}
}
