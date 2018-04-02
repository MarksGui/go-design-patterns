package builder

import "fmt"

type Builder interface {
	CreateVehicle()
	AddWheel()
	AddEngine()
	AddDoors()
	GetVehicle() string
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Build() string {
	d.builder.CreateVehicle()
	d.builder.AddWheel()
	d.builder.AddEngine()
	d.builder.AddDoors()
	return d.builder.GetVehicle()
}

type TruckBuilder struct {
	vehicle *Vehicle
}

func (t *TruckBuilder) CreateVehicle() {
	t.vehicle = &Vehicle{}
}

func (t *TruckBuilder) AddWheel() {
	t.vehicle.Wheel = "四个卡车轮子"
}

func (t *TruckBuilder) AddEngine() {
	t.vehicle.Engine = "一个卡车发动机"
}

func (t *TruckBuilder) AddDoors() {
	t.vehicle.Doors = "两个卡车门"
}
func (t *TruckBuilder) GetVehicle() string {
	return fmt.Sprintf("卡车，车轮: %s，发动机：%s，车门：%s",
		t.vehicle.Wheel, t.vehicle.Engine, t.vehicle.Doors)
}

type CarBuilder struct {
	vehicle *Vehicle
}

func (c *CarBuilder) CreateVehicle() {
	c.vehicle = &Vehicle{}
}

func (c *CarBuilder) AddWheel() {
	c.vehicle.Wheel = "四个轿车轮子"
}

func (c *CarBuilder) AddEngine() {
	c.vehicle.Engine = "一个轿车发动机"
}

func (c *CarBuilder) AddDoors() {
	c.vehicle.Doors = "四个轿车车门"
}

func (c *CarBuilder) GetVehicle() string {
	return fmt.Sprintf("轿车，车轮: %s，发动机：%s，车门：%s",
		c.vehicle.Wheel, c.vehicle.Engine, c.vehicle.Doors)
}

type Vehicle struct {
	Wheel  string
	Engine string
	Doors  string
}
